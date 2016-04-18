package php

import (
	"bytes"
	"text/template"

	"github.com/previousnext/tuner/conf"
)

const tpl = `memory_limit = {{ . }}M`

type PHP struct {
	max  int
	proc int
}

func init() {
	conf.Register("php", &PHP{})
}

func (p *PHP) Max(m int) {
	p.max = m
}

func (p *PHP) Proc(pr int) {
	p.proc = pr
}

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
