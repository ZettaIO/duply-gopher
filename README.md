
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

## Dependencies

- GnuPG 2.x
- Haveged

## Setting for development

Run `godep restore` in the root of the project to get all dependencies.
If new dependencies are added run `godep save ./...` in the project root.

### Libraries

- https://github.com/mitchellh/multistep
- https://github.com/jasonlvhit/gocron
- https://github.com/golang/glog

### Tools

https://github.com/tools/godep
