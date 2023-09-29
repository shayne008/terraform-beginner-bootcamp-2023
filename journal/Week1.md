# Terraform Beginner Bootcamp 2023 - Week 1

## Root Module Structure

The root module structure is as follows:


```
PROJECT ROOT
│
├── main.tf - everything else
├── variables.tf - stores the structure of input variables
├── terraform.tfvars - the data of variables we want to load into our Terraform project
├── providers.tf - defined required providers and their organization
├── outputs.tf - stores our outputs
└── README.md - required for root modules
```

[Standard Module Structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure)

## Terraform and Input Variables
### Terraform Cloud Variables

In terraform you can set two kinds of variables:
- Enviroment Variables - you would set in bash terminal eg. AWS credentials
- Terraform Variables - you would set in tfvars file

You can set terraform cloud variables to be sensitive so they are not shown in the UI.

## Loading Terraform Input Variables

[Input Variables](https://developer.hashicorp.com/terraform/language/values/variables)
### var flag
You can use the `-var` flag to set an input variable or override a variable in the tfvars file eg. `terraform -var user_uuid="my_user_id"`

### var-file flag

You use to pass Input Variable values into Terraform plan and apply commands using a file that contains the values eg. below

```
# Pass .tfvar files to 'plan' command
terraform plan \
-var-file 'production.tfvars' \
-var-file 'secrets.tfvars'

# Pass .tfvar files to 'apply ' command
terraform apply \
-var-file 'production.tfvars' \
-var-file 'secrets.tfvars'
```
[Use Terraform Input Variables](https://build5nines.com/use-terraform-input-variables-to-parameterize-infrastructure-deployments/#:~:text=The%20%2Dvar%2Dfile%20flag%20is,file%20that%20contains%20the%20values.)

### terraform.tvfars

This is the the default file to load in terraform variables in blank

### auto.tfvars

`auto.tfvars` is a special filename that Terraform automatically looks for and loads when you run Terraform commands like `terraform apply` or `terraform plan`

### Order of Terraform Variables

- Command-Line Variables
- Variable Files
- Environment Variables
- Default Values in Configuration Files
- Providers and Modules


## Dealing With Configuration Drift 

## What happens if you lose the state file?

If you lose the statefile, you will more than likely have to tear down all the cloud infrastructure manually.

You can use terraform import but it will not work for all cloud resources. Please check the terraform providers documentation.
### Fix Missing Resources with Terraform Import


`terraform import aws_s3_bucket.bucket bucket-name`

[Terraform Input](https://developer.hashicorp.com/terraform/language/import)
[aws_s3_bucket import](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#import)


### Fix Manual Configuration

If someone deletes or modifies cloud resources manually through ClickOps.

If we run Terraform Plan with the attempt to put our infrastructure back into the expected state fixing Configuration Drift

## Fix using Terraform Refresh

```sh
terraform apply -refresh-only -auto-approve
```

## Terraform Modules

### Terraform Module Structure
It is recommend to place modules in a `modules` directory when locally
developing modules but you can name it whatever you like.

### Passing Input Variables

You can pass input variables to the module.
The module has to declare the terraform variables in its own variables.tf

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
}
```

### Modules Sources

Using the source you can import the module from various places eg:
- locally
- GitHub
- Terraform Registry

```tf
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
}
```

[Module Sources](https://developer.hashicorp.com/terraform/language/modules/sources)

## Considerations when using ChatGPT to write Terraform

LLMs such as ChatGPT may not be trained on the latest documentation or information about Terraform.

It may likely produce older examples that could be depreciated, often affecting the providers.

## Working with Files in Terraform

## Fileexists Function

This is a builtin terraform function to check the file existance.

```tf
    condition     = fileexists(var.error_html_filepath)
```
[fileexists Function](https://developer.hashicorp.com/terraform/language/functions/fileexists)

## Filemd5



[filemd5 Function](https://developer.hashicorp.com/terraform/language/functions/filemd5)

### Path Variable

In Terraform, a "path variable" typically refers to a variable that represents a file or directory path.
- path.module = the filesystem path of the module where the expression is placed
- path.root = the filesystem path of the root module of the configuration.


[Filesystem and Workspace Info](https://developer.hashicorp.com/terraform/language/expressions/references#filesystem-and-workspace-info)

resource "aws_s3_object" "index_html" {
  bucket = aws_s3_bucket.website_bucket.bucket
  key    = "index.html"
  source = "${path.root}/public/index.html"

  # The filemd5() function is available in Terraform 0.11.12 and later
  # For Terraform 0.11.11 and earlier, use the md5() function and the file() function:
  # etag = "${md5(file("path/to/file"))}"
  etag = filemd5(var.index_html_filepath)
}

## Terraform Locals
 
In Terraform, locals are a way to define reusable values or expressions within your configuration.

 This is useful when you want to reference cloud resources but not import them.

 ```tf
 locals {
  service_name = "forum"
  owner        = "Community Team"
}

 ```

 [Terraform Locals](https://developer.hashicorp.com/terraform/language/values/locals)
### Terraform Data Sources

Data sources allow Terraform to use information defined outside of Terraform, defined by another separate Terraform configuration, or modified by functions.

[Data Sources](https://developer.hashicorp.com/terraform/language/data-sources)

### Working with JSON

Used the jsonencode to create the json policy inline in the hcl.

```tf
> jsonencode({"hello"="world"})
{"hello":"world"}

```

[jsonencode Function](https://developer.hashicorp.com/terraform/language/functions/jsonencode)
