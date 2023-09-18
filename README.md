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