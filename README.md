# GCP Janitor

A GitHub Action to cleanup GCP resources.

It supports cleaning up the following services:

- VM Instances and Disks

It follows this strict order to avoid failures caused by inter-resource dependencies. Although intermittent failures may occur, they should be resolved in subsequent executions.

## Inputs

| Name               | Required | Description                                                                                                 |
| ------------------ | -------- | ----------------------------------------------------------------------------------------------------------- |
| CredentialsJSON    | Y        | The JSON key file for GCP service account credentials.                                                      |
| Zones              | Y        | A comma-separated list of zones to clean resources in. Use `*` for all zones within the selected regions.   |
| ProjectID          | Y        | The GCP project ID where the resources are located.                                                         |
| AgeInHours         | N        | The minimum age (in hours) of resources to be eligible for cleaning. Default to 0.                          |
| DryRun             | N        | If set to `true`, performs a dry run without deleting any resources. Defaults to `false`.                   |
| ResourceLabelKey   | Y        | The label key used to identify resources that should not be deleted.                                        |
| ResourceLabelValue | Y        | The label value associated with `ResourceLabelKey` to mark resources for exclusion from deletion.           |

## Example Usage

```yaml
jobs:
  cleanup:
    runs-on: ubuntu-latest
    name: Cleanup resources
    steps:
      - name: Cleanup
        uses: rancher-sandbox/gcp-janitor@v0.1.0
        with:
            credentialsJSON: ${{secrets.GCP_CREDENTIALS}}
            zones: europe-west2-c
            projectID: my-project
            ageInHours: 6
            resourceLabelKey: name
            resourceLabelValue: highlander
```
