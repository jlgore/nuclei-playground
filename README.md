# Nuclei Code Protocol Playground

## Proving Concepts

### Requirements

* Pulumi and a recent version of Golang installed
* Nuclei installed

### Steps

* Clone this repo
* Sign the yaml template: `nuclei -t rds-public-config.yaml -sign`
* Set your `AWS_PROFILE` environment variable: `export AWS_PROFILE=profile_name`
* Execute the Nuclei template: `nuclei -t rds-public-config.yaml -code`

The Pulumi will create two RDS instances, one public and one private. Both exist in the default VPC (for now). The Nuclei code template will deploy the Pulumi and then scan for a publicly accessible databases.  
