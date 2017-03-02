
# duply-gopher

**This is still work in progress.**

Go service configuring and running Duplicity backups to OpenStack Swift using the [Duply](http://duply.net/) wrapper for
additional tooling. The service has it's own built in scheduler running backups.

See [configuration](docs/configuration.md).

## What the service is currently doing:

- Fully configures Duply ready to start backing up
  - Creates a gpg2 key chain and import the supplied keys including owner trusts
  - Generates Duply configration files
    - Profile configuration file controlling backup parameters
    - Globbing file list

## Additional features (Not done yet):

- Report backup status through HTTP or AMQP.
- Manually trigger additional backups through HTTP or AMQP
- Permanently or temporarily change configuration

# Enviroment variables

Configuration can also be defined or overridden with environment variables.

| Name                      | Description |
|---------------------------|-------------|
| SWIFT_AUTHURL             | Swift authentication
| SWIFT_AUTHVERSION         | Swift authentication
| SWIFT_REGION_NAME         | Swift authentication
| SWIFT_USERNAME            | Swift authentication
| SWIFT_PASSWORD            | Swift authentication
| SWIFT_TENANTNAME          | Swift authentication
| SWIFT_USER_DOMAIN_NAME    | Swift authentication
| SWIFT_PROJECT_DOMAIN_NAME | Swift authentication

## Dependencies

- GnuPG 2.x
- Haveged

## Setting for development

Run `godep restore` in the root of the project to get all dependencies.
If new dependencies are added run `godep save ./...` in the project root.

### Libraries

- https://github.com/mitchellh/multistep
- https://github.com/robfig/cron
- https://github.com/golang/glog

### Tools

https://github.com/tools/godep
