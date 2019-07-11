package php

import (
	"bytes"
	"text/template"

	"github.com/previousnext/tuner/conf"
)

const tpl = `memory_limit = {{ . }}M`

// PHP struct.
type PHP struct {
	max  int
	proc int
}

func init() {
	conf.Register("php", &PHP{})
}

// Max is the max available memory.
func (p *PHP) Max(m int) {
	p.max = m
}

// Proc is the max proc size.
func (p *PHP) Proc(pr int) {
	p.proc = pr
}

// Build builds the template.
func (p *PHP) Build() (string, error) {
	t := template.Must(template.New(tpl).Parse(tpl))

	// Write the contents to memory.
	buf := new(bytes.Buffer)
	err := t.Execute(buf, p.proc)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
