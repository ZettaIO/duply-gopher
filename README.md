
# duply-gopher

Work in progress!

Go service configuring and running Duply and duplicity backups.
The service has it's own built in scheduler running backups.
Can communicate through AMQP and HTTP.

Currently needs the following packages:
- GnuPG
- Haveged

## Setting for development

Run `godep restore` in the root of the project to get all dependencies.
If new dependencies are added run `godep save ./...` in the project root.

### Libraries

- https://github.com/jasonlvhit/gocron
- https://github.com/golang/glog

### Tools

https://github.com/tools/godep
