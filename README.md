# Rage4 DNS Provider for Terraform

## Overview
This project is a plugin for Terraform, that allows it to work with the Rage4 DNS service.  When you use terraform to create servers on a cloud, you can also have it create a matching DNS entry on Rage4.

## Building
```
go build -o terraform-provider-rage4
```
## Installing
* clone the repo to your local machine
* add the following to the file `~/.terraformrc`
```
providers {
    rage4 = "/path_directory/terraform-rage4/terraform-provider-rage4"
}
```
* you can test that the plugin is installed by having terraform create an execation plan for `sample.tf`
```
$ terraform plan /path_to_directory/sample.tf
```

## Using
To use the plugin, you can use the file `sample.tf` as a reference.  First you need to add your account email and API key to the terrform script (in the provider section).  Once that is done, you are have terraform manage the DNS records for the servers that you have terraform managing.

## Debugging

```
terraform plan
```
or 
```
export TF_LOG=1 
terraform apply 
```