package config

import "fmt"

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

// Verify swift auth
// FIXME: There's probably a much simpler way to verify things..
func (a *SwiftAuth) Verify() error {
	err := errEmpty("duply.auth.swift_auth_url", a.URL)
	if err != nil {
		return err
	}
	err = errEmpty("duply.auth.swift_auth_version", a.Version)
	if err != nil {
		return err
	}
	err = errEmpty("duply.auth.swift_region_name", a.Region)
	if err != nil {
		return err
	}
	err = errEmpty("duply.auth.swift_username", a.Username)
	if err != nil {
		return err
	}
	err = errEmpty("duply.auth.swift_password", a.Password)
	if err != nil {
		return err
	}
	err = errEmpty("duply.auth.swift_project_name", a.ProjectName)
	if err != nil {
		return err
	}
	err = errEmpty("duply.auth.swift_user_domain_name", a.UserDomainName)
	if err != nil {
		return err
	}
	err = errEmpty("duply.auth.swift_project_domain_name", a.ProjectDomainName)
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

func errEmpty(name, value string) error {
	if value == "" {
		return fmt.Errorf("Field '%v' cannot be empty", name)
	}
	return nil
}
