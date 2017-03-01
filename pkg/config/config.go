package config

import (
	"fmt"
	"path"
)

// Config for the service
type Config struct {
	Duply    DuplyConfig    `yaml:"duply"`
	HTTP     HTTPConfig     `yaml:"http"`
	RabbitMq RabbitMqConfig `yaml:"rabbitmq"`
}

// DuplyConfig parameters
type DuplyConfig struct {
	Home                   string              `yaml:"config_home"`
	Root                   string              `yaml:"root"`
	ArchDir                string              `yaml:"arch_dir"`
	MaxAge                 string              `yaml:"max_age"`
	MaxFullBackupAge       string              `yaml:"max_full_backup_age"`
	MaxFullBackups         string              `yaml:"max_full_backups"`
	MaxFullBackupsWithIncr string              `yaml:"max_full_with_incrs"`
	VolSize                string              `yaml:"volsize"`
	ContainerName          string              `yaml:"container_name"`
	Auth                   SwiftAuth           `yaml:"auth"`
	Keys                   DuplyKeys           `yaml:"keys"`
	Globbing               map[string][]string `yaml:"globbing"`
}

// SwiftAuth config
type SwiftAuth struct {
	URL               string `yaml:"swift_auth_url"`
	Version           string `yaml:"swift_auth_version"`
	Region            string `yaml:"swift_region_name"`
	Username          string `yaml:"swift_username"`
	Password          string `yaml:"swift_password"`
	ProjectName       string `yaml:"swift_project_name"`
	UserDomainName    string `yaml:"swift_user_domain_name"`
	ProjectDomainName string `yaml:"swift_project_domain_name"`
}

// DuplyKeys contains gpg keys
type DuplyKeys struct {
	Master DuplyMasterKey `yaml:"master"`
	Host   DuplyHostKey   `yaml:"host"`
}

// DuplyMasterKey ...
type DuplyMasterKey struct {
	ID   string `yaml:"id"`
	Data string `yaml:"data"`
}

// DuplyHostKey ...
type DuplyHostKey struct {
	ID       string `yaml:"id"`
	Password string `yaml:"password"`
	Public   string `yaml:"public"`
	Private  string `yaml:"private"`
}

// HTTPConfig config
type HTTPConfig struct {
	Port         string `yaml:"port"`
	SharedSecret string `yaml:"shared_secret"`
}

// RabbitMqConfig ..
type RabbitMqConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// GPGRoot returns the root directory for gpg data
func (d *DuplyConfig) GPGRoot() string {
	return path.Join(d.Home, "gpg")
}

// ProfileRoot returns the root directory for gpg data
func (d *DuplyConfig) ProfileRoot() string {
	return path.Join(d.Home, "profile")
}

// Env returns enviroment variables
func (d *DuplyConfig) Env() []string {
	return []string{
		fmt.Sprintf("GNUPGHOME=%s", path.Join(d.Home, "gpg")),
		// Swift authentication
		fmt.Sprintf("SWIFT_AUTHURL=%s", d.Auth.URL),
		fmt.Sprintf("SWIFT_USERNAME=%s", d.Auth.Username),
		fmt.Sprintf("SWIFT_PASSWORD=%s", d.Auth.Password),
		fmt.Sprintf("SWIFT_REGION_NAME=%s", d.Auth.Region),
		fmt.Sprintf("SWIFT_USER_DOMAIN_NAME=%s", d.Auth.UserDomainName),
		fmt.Sprintf("SWIFT_PROJECT_DOMAIN_NAME=%s", d.Auth.ProjectDomainName),
		fmt.Sprintf("SWIFT_TENANTNAME=%s", d.Auth.ProjectName),
		fmt.Sprintf("SWIFT_AUTHVERSION=%s", d.Auth.Version),
	}
}
