package config

import (
	"bytes"
	"html/template"
	"os"
	"os/exec"
	"path"

	"github.com/golang/glog"
)

const defBufferSize = 4096

// Template ..
type Template struct {
	tmplFile  string
	dstFile   string
	tmpl      *template.Template
	tmplBuf   *bytes.Buffer
	outCmdBuf *bytes.Buffer
}

// NewTemplate ..
func NewTemplate(tmplFile, dstFile string) (*Template, error) {
	tmpl := &Template{
		tmplFile:  tmplFile,
		dstFile:   dstFile,
		tmplBuf:   bytes.NewBuffer(make([]byte, 0, defBufferSize)),
		outCmdBuf: bytes.NewBuffer(make([]byte, 0, defBufferSize)),
	}
	funcMap := template.FuncMap{
		"ShortID": func(s string) string {
			return s[(len(s) - 8):]
		},
	}
	t, err := template.New(path.Base(tmplFile)).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		return nil, err
	}
	tmpl.tmpl = t
	return tmpl, nil
}

// Execute the template
func (t *Template) Execute(data interface{}) error {
	f, err := os.Create(t.dstFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.tmpl.Execute(t.tmplBuf, data)
	if err != nil {
		return err
	}

	cmd := exec.Command("scripts/clean-conf.sh")
	cmd.Stdin = t.tmplBuf
	cmd.Stdout = t.outCmdBuf
	if err = cmd.Run(); err != nil {
		glog.Warningf("unexpected error cleaning config file: %v", err)
		_, err = f.Write(t.tmplBuf.Bytes())
		if err != nil {
			return err
		}
		return nil
	}
	_, err = f.Write(t.outCmdBuf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
