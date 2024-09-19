package action

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func (a *action) cleanVMs(ctx context.Context, input *CleanupScope) error {
	computeService, err := compute.NewService(ctx, option.WithCredentialsJSON(input.CredentialsJSON))
	if err != nil {
		log.Fatalf("Failed to create compute service: %v", err)
	}

	instancesService := compute.NewInstancesService(computeService)
	zonesService := compute.NewZonesService(computeService)

	for _, zone := range input.Zones {
		if zone == "*" {
			if err := deleteVMsinAllZones(ctx, instancesService, zonesService, input); err != nil {
				return fmt.Errorf("failed to delete instances in all zones: %w", err)
			}
			return nil
		}
		if err := deleteVMsinZone(ctx, zone, instancesService, input); err != nil {
			return fmt.Errorf("failed to delete instances in zone %s: %w", zone, err)
		}
	}

	return nil
}

func deleteVMsinAllZones(ctx context.Context, instancesService *compute.InstancesService, zonesService *compute.ZonesService, input *CleanupScope) error {
	Log("Cleaning VMs in all zones")

	zonesList, err := zonesService.List(input.ProjectID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("failed to list zones: %w", err)
	}

	for _, zone := range zonesList.Items {
		if err := deleteVMsinZone(ctx, zone.Name, instancesService, input); err != nil {
			return fmt.Errorf("failed to delete instances in zone %s: %w", zone.Name, err)
		}
	}

	return nil
}

func deleteVMsinZone(ctx context.Context, zone string, instancesService *compute.InstancesService, input *CleanupScope) error {
	Log("Cleaning VMs in zone %s", zone)

	filter := fmt.Sprintf("labels.%s=%s", input.ResourceLabelKey, input.ResourceLabelValue)

	instancesList, err := instancesService.List(input.ProjectID, zone).Filter(filter).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("failed to list instances: %w", err)
	}

	for _, instance := range instancesList.Items {
		creationTime, err := time.Parse(time.RFC3339, instance.CreationTimestamp)
		if err != nil {
			return fmt.Errorf("failed to parse creation time: %w", err)
		}

		ageInHoursTime := time.Now().Add(-time.Duration(input.AgeInHours) * time.Hour)

		if creationTime.Before(ageInHoursTime) {
			Log("Deleting instance %s", instance.Name)

			if input.DryRun {
				Log("Dry run: skipping deletion of instance %s", instance.Name)
				continue
			}

			if _, err := instancesService.Delete(input.ProjectID, zone, instance.Name).Context(ctx).Do(); err != nil {
				return fmt.Errorf("failed to delete instance: %w", err)
			}
		}
	}

	return nil
}
