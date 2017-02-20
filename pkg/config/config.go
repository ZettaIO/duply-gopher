package config

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/go-yaml/yaml"
)

// Config for the service
type Config struct {
	Duply    DuplyConfig    `yaml:"duply"`
	HTTP     HTTPConfig     `yaml:"http"`
	RabbitMq RabbitMqConfig `yaml:"rabbitmq"`
}

// DuplyConfig parameters
type DuplyConfig struct {
	Home          string    `yaml:"config_home"`
	Root          string    `yaml:"root"`
	ContainerName string    `yaml:"container_name"`
	Auth          SwiftAuth `yaml:"auth"`
	Keys          DuplyKeys `yaml:"keys"`
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

// NewConfig creats a new config
func NewConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	conf := NewDefaults()
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	err = conf.Verify()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// GPGRoot returns the root directory for gpg data
func (d *DuplyConfig) GPGRoot() string {
	return path.Join(d.Home, "gpg")
}

// Env returns enviroment variables
func (c *Config) Env() []string {
	return []string{
		fmt.Sprintf("GNUPGHOME=%s", path.Join(c.Duply.Home, "gpg")),
		// Swift authentication
		fmt.Sprintf("SWIFT_AUTHURL=%s", c.Duply.Auth.URL),
		fmt.Sprintf("SWIFT_USERNAME=%s", c.Duply.Auth.Username),
		fmt.Sprintf("SWIFT_PASSWORD=%s", c.Duply.Auth.Password),
		fmt.Sprintf("SWIFT_REGION_NAME=%s", c.Duply.Auth.Region),
		fmt.Sprintf("SWIFT_USER_DOMAIN_NAME=%s", c.Duply.Auth.UserDomainName),
		fmt.Sprintf("SWIFT_PROJECT_DOMAIN_NAME=%s", c.Duply.Auth.ProjectDomainName),
		fmt.Sprintf("SWIFT_TENANTNAME=%s", c.Duply.Auth.ProjectName),
		fmt.Sprintf("SWIFT_AUTHVERSION=%s", c.Duply.Auth.Version),
	}
}
