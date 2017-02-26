
# Configuration

The service is currently configured using a yaml file or environment variables.

Currently you need to supply both the master and host key in the configuration
but we may also let the host generate it's own key. The advantage of supplying
both keys is that you don't need persistent storage for the configuration
and will not accidentally generate a new host key on a re-deploy.

See also [notes about gpg](gpg.md).

# Configuration sections

See also [config.json example](../examples/config.yaml).

## duply

- `config_home`: Configuration location
- `container_name`: Swift container name
- `root`: The root path you are backing up
- `arch_dir`: Where Duply stores its local archive
- `volsize`: Size in MB
- `max_age`: Max age of backup
- `max_full_backups`: Max number of full backups duply will retain
- `max_full_backup_age`: how old a full backup must be before a new full backup will be created
- `max_full_with_incrs`: ..
- `auth`: See auth section below
- `globbing`: See globbing section below
- `keys`: See keys section

### auth

- swift_auth_url
- swift_auth_version
- swift_region_name
- swift_username
- swift_password
- swift_project_name
- swift_user_domain_name
- swift_project_domain_name

Can also be omitted and defined thought environment variables (Here mapped from openstack rc values):
```
SWIFT_AUTHURL=${OS_AUTH_URL}
SWIFT_AUTHVERSION=${OS_IDENTITY_API_VERSION}
SWIFT_USERNAME=${OS_USERNAME}
SWIFT_PASSWORD=${OS_PASSWORD}
SWIFT_REGION_NAME=${OS_REGION_NAME}
SWIFT_USER_DOMAIN_NAME=${OS_PROJECT_DOMAIN_NAME}
SWIFT_PROJECT_DOMAIN_NAME=${OS_USER_DOMAIN_NAME}
SWIFT_TENANTNAME=${OS_PROJECT_NAME}
```

### globbing

Currently we only support either include or exclude. If both are defined the service will not start.

- `exclude`: List of paths to exclude relative to `root`
- `include`: List of paths to include relative to `root`. All other paths will be excluded.

### keys

- `master`
  - `id`: Id of the master key
  - `data`: The public part of master key
- `host`
  - `id`: id of the host key
  - `password`: Password for the host key
  - `public`: Public key data
  - `private`: Private key data

## amqp

Configuration parameters for an apmq service (dummy values for now)

## http

Configuration parameters for an http service (dummy values for now)
