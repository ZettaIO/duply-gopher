package config

import "path"

const (
	profileTemplate  = "templates/profile.tmpl"
	globbingTemplate = "templates/globbing.tmpl"
)

// WriteProfile writes the duply profile
func (d *DuplyConfig) WriteProfile() error {
	dest := path.Join(d.ProfileRoot(), "conf")
	t, err := NewTemplate(profileTemplate, dest)
	if err != nil {
		return err
	}
	err = t.Execute(d)
	if err != nil {
		return err
	}
	return nil
}

// WriteGlobbingList writes the globbing list file
func (d *DuplyConfig) WriteGlobbingList() error {
	dest := path.Join(d.ProfileRoot(), "exclude")
	t, err := NewTemplate(globbingTemplate, dest)
	if err != nil {
		return err
	}

	err = t.Execute(d)
	if err != nil {
		return err
	}
	return nil
}
