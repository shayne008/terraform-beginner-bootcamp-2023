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