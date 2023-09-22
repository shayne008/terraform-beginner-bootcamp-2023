# Terraform Beginner Bootcamp 2023

## Installing Terraform CLI


### Considerations with the Terraform CLI changes
The Terraform CLI install instructions have changed due to gpg keyring changes. The original gitpod.yml did not work. So we referred to the latest Terraform Documentation and changed the scripting for install.

[Install Terraform CLI](https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)

## Considerations for Linux Distribution

This project is built with Ubuntu.
Check your Linux Distribution and change to fit your needs.

Example of Version

```
$ cat /etc/os-release 
PRETTY_NAME="Ubuntu 22.04.3 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.3 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
```

[How To Check Linux Version](https://linuxize.com/post/how-to-check-linux-version/)

### Refactoring into Bash Scripts

While fixing the Terraform CLI gpg depreciation issue we noticed it was more code with new steps to install. So a bash script was created to install the Terraform CLI.

The bash script is located here: [./bin/install_terraform_cli](/bin/install_terraform_cli)

- This will keep the Gitpod Task File [.gitpod.yml](.gitpod.yml) tidy.
- This allows us an easier path to debug and execute manually install Terraform CLI
- This will allow better portability for other projects that need to install Terraform CLI.

### Shebang Considerations

A shebang (pronouced Sha-bang) tells the bash script what program that will interpret the script. eg.`#!/bin/bash`

ChatGPT recommends the following format: `#!/usr/bin/env bash`

This shebang uses the /usr/bin/env command to locate the bash interpreter in the system's PATH environment variable. This approach is more flexible because it relies on the user's environment to find the Bash interpreter, which can be in different locations on different systems.

### Execution Considerations

When executing the bash script use `./` shorthand notation.

eg. `./bin/install_terraform_cli`

If using a script in .gitpod.yml, it needs to point the script to a program to interpet it.

eg. `source ./bin/install_terraform_cli`

[Shebang](https://en.wikipedia.org/wiki/Shebang_(Unix))


#### Permission Considerations

[Chmod](https://en.wikipedia.org/wiki/Chmod)


Needed to change the permissions to make the file executable at user level.

```sh
chmod u+x ./bin/install_terraform_cli
```
### Gitpod Lifecycle (Before, Init, Command)

We changed for init to before so we can run the bash script with no issues. Use caution when using init because it will not rerun when using an existing workspace.




[GitPod Tasks](https://www.gitpod.io/docs/configure/workspaces/tasks)

### Working with Env Vars

#### env command

You can list out all the Environment Variables (Env Vars) using `env` command

Can filter specfic env vars using grep eg. `env | grep AWS_`

#### Setting and Unsetting Env Vars

In the terminal, set by using `export HELLO='world'`
Can unset by using `unset HELLO`

Can set an env var temporarily when running a command

```sh
HELLO='world' ./bin/print_message
```

Within a bash script, you can set a env without wriiting export eg.

```sh
#!/usr/bin/env bash
HELLO='world'

echo $HELLO
```

#### Printing Env Vars

Can print an env var using echo eg. `echo $HELLO`

#### Scoping of Env Vars

When you open up new bash terminals in VSCode it will not be aware of env vars that are set in another window.

If you want Env Vars to persist across all future bash terminals that are open you need to set env vars in the bash profile. eg. `.bash_profile`

#### Persisting Env Vars in Gitpod

You can persist env vars into gitpod by storing them in Gitpod Secrets.

```
gp env HELLO='world'
```

All future workspaces launched will set the env vars for all bash terminals opened in those workspaces.

You can set env vars in the `.gitpod.yml` but this can only contain non senstive env vars.

### AWS CLI Installation

AWS CLI is installed for the project via bash script [`./bin/install_aws-cli`](./bin/install_aws_cli)

[Getting Started AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
[AWS Env Vars](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)


You can check the AWS credentials configuration by running the following command

```sh
aws sts get-caller-identity
```

If it is successful then you should see a json payload return that looks like the below:

```json
{
    "UserId": "AKIAIOSFODNN7EXAMPLE",
    "Account": "123456789012",
    "Arn": "arn:aws:iam::123456789012:user/terraform-beginner-bootcamp"
}
```
Will need to generate AWS CLI credentials for IAM User in order to use the AWS CLI.

These are the steps to create an IAM User in AWS:

[Create an IAM User](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html)


# Terraform Basics

### Terraform Registry

Terraform sources providers and modules from the Terraform registry which is located at [registry.terraform.io](https://registry.terraform.io/)

- **Providers** Terraform providers are software components that enable Terraform, an Infrastructure as Code tool, to interact with and manage resources in various cloud, infrastructure, and service environments by abstracting the underlying APIs and allowing users to define, configure, and manage resources using a declarative language.
- **Modules** Terraform modules are reusable, encapsulated units of Terraform configuration that allow you to organize and package infrastructure code in a modular and maintainable way. They are a fundamental concept in Terraform that promotes code reuse, modularity, and best practices in infrastructure as code (IaC) development.

[Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random/latest)

## Terraform Console

Terraform console is a command-line tool in Terraform that allows you to interactively evaluate expressions and query data about your infrastructure. Just type in `terraform`

### Terraform Init

At the start of a project you will use `terraform init` to download essential files for the terraform providers.

### Terraform Plan

`terraform plan` is a command in Terraform that analyzes your infrastructure code and generates an execution plan. Provides you with valuable information about what Terraform intends to do when you apply your configuration.

Can ouput this plan to passed a apply.

### Terraform Apply

`terraform apply` is a command in Terraform used to execute the changes specified in your Terraform configuration and apply them to your infrastructure. It is a critical step in the Terraform workflow and is used to create, update, or delete resources to achieve the desired state defined in your configuration.

Should prompt for yes or no. Can automatically approve with auto approve flag. `terraform apply --auto-approve`

**NOTE** Make sure you follow the bucket name rules, had to change the lower, upper and special characters around to make it work.
- Also got a few errors with s3 access, had to allow permissions
[S3 Permissions](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-policy.html)
- Got a invalidbucketname error:
> invalidbucketname: the specified bucket is not valid.
- Resolved it by checking for errors had quotes in the wrong place

### Terraform Destroy

`terraform destroy` is a command in Terraform that is used to tear down and destroy the infrastructure resources that were created and managed by Terraform.

### Terraform Lock Files

`.terraform.lock.hcl` contains the locked versionsing for the providers or modules that should be used with the project.

The Terraform Lock File ** should be committed** to the Version Control System (VSC) example GitHub

### Terraform State Files

`terraform.tfstate` contains information about the state of the infrastructure. 

This file **should not be committed** to your VSC.

The file contains senstive data, do not lose the file. Knows the state of infrastructure.

### Terraform Directory

`.terraform` is a directory on your file system that contains all the configuration files and associated resources for a specific Terraform project or deployment. This directory serves as the workspace for managing and defining your infrastructure as code (IaC). 

## Issues with Terraform Cloud login and Gitpod Workspace

When using `terraform init` will launch bash to view token. Had to hit the P (print)
and Ctrl Click to open the link to create the token.

### Migrate to Terraform Cloud

Use the below document to help migrate to terraform cloud:

[Migrate state to Terraform Cloud](https://developer.hashicorp.com/terraform/tutorials/cloud/cloud-migrate)


We automated the process with a bash script [bin/generate_credentials_tfrc](bin/generate_credentials_tfrc)


