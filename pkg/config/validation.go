package config

import (
	"fmt"
	"os"
)

// Verify config value
func (c *Config) Verify() error {
	return c.Duply.Verify()
}

// Verify DuplyConfig values
// FIXME: There's probably a much simpler way to verify things..
func (d *DuplyConfig) Verify() error {
	err := errEmpty("duply.container_name", d.ContainerName)
	if err != nil {
		return err
	}
	err = errEmpty("duply.home", d.Home)
	if err != nil {
		return err
	}
	err = d.Auth.Verify()
	if err != nil {
		return err
	}
	err = d.Keys.Verify()
	if err != nil {
		return err
	}
	return nil
}

// Verify swift auth parameters
func (a *SwiftAuth) Verify() error {
	// Finally check all values
	err := envOrVal("duply.auth.swift_auth_url", "SWIFT_AUTHURL", &a.URL)
	if err != nil {
		return err
	}
	err = envOrVal("duply.auth.swift_auth_version", "SWIFT_AUTHVERSION", &a.Version)
	if err != nil {
		return err
	}
	err = envOrVal("duply.auth.swift_region_name", "SWIFT_REGION_NAME", &a.Region)
	if err != nil {
		return err
	}
	err = envOrVal("duply.auth.swift_username", "SWIFT_USERNAME", &a.Username)
	if err != nil {
		return err
	}
	err = envOrVal("duply.auth.swift_password", "SWIFT_PASSWORD", &a.Password)
	if err != nil {
		return err
	}
	err = envOrVal("duply.auth.swift_project_name", "SWIFT_TENANTNAME", &a.ProjectName)
	if err != nil {
		return err
	}
	err = envOrVal("duply.auth.swift_user_domain_name", "SWIFT_USER_DOMAIN_NAME", &a.UserDomainName)
	if err != nil {
		return err
	}
	err = envOrVal("duply.auth.swift_project_domain_name", "SWIFT_PROJECT_DOMAIN_NAME", &a.ProjectDomainName)
	if err != nil {
		return err
	}
	return nil
}

// Verify keys
func (d *DuplyKeys) Verify() error {
	err := errEmpty("duply.keys.master.id", d.Master.ID)
	if err != nil {
		return err
	}
	err = errEmpty("duply.keys.master.data", d.Master.Data)
	if err != nil {
		return err
	}
	return nil
}

// errEmpty checks if a string is empty and creates an appropriate error
func errEmpty(name, value string) error {
	if value == "" {
		return fmt.Errorf("Field '%v' cannot be empty", name)
	}
	return nil
}

func envOrVal(name, env string, value *string) error {
	e := os.Getenv("SWIFT_PROJECT_DOMAIN_NAME")
	if e != "" {
		*value = e
		return nil
	}
	return errEmpty(name, *value)
}

// // errEmpty checks if a string is empty and creates an appropriate error
// func errEmpty(name, env string, value *string) error {
// 	if value == "" {
// 		return fmt.Errorf("Field '%v' cannot be empty", name)
// 	}
// 	// os.Getenv("SWIFT_PROJECT_DOMAIN_NAME")
// 	return nil
// }
