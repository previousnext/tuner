package php

import (
	"bytes"
	"text/template"

	"github.com/previousnext/tuner/conf"
)

const tpl = `memory_limit = {{ . }}M`

// PHP struct.
type PHP struct {
	totalMemory int
	avgProc     int
	maxProc     int
}

func init() {
	conf.Register("php", &PHP{})
}

// TotalMemory is the total available totalMemory.
func (p *PHP) TotalMemory(m int) {
	p.totalMemory = m
}

// AvgProc is the max proc size.
func (p *PHP) AvgProc(pr int) {
	p.avgProc = pr
}

// MaxProc is the max proc size.
func (p *PHP) MaxProc(pr int) {
	p.maxProc = pr
}

// Build builds the template.
func (p *PHP) Build() (string, error) {
	t := template.Must(template.New(tpl).Parse(tpl))

	// Write the contents to memory_limit.
	buf := new(bytes.Buffer)
	err := t.Execute(buf, p.maxProc)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
