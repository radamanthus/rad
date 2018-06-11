# rad

rad is a companion CLI app for EngineNiRad

# Usage

rad init
- provision and configure resources for the first time
- does not run if there are already existing resources, unless the --force flag is enabled
implementation:
- generates a deploy key pair for a machine user and prints the public key that the user should add to Github
- parses rad.yml, generates configure.sh and runs it across all instances
- terraform apply

rad update
- parses rad.yml, generates configure.sh and runs it across all instances

implementation:
- parses rad.yml, generates configure.sh and runs it across all instances
- terraform apply

rad rebuild
- same as rad init --force

rad destroy
- terraform destroy
