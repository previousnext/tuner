package apache

import (
	"bytes"
	"text/template"

	"github.com/previousnext/tuner/conf"
)

const tpl = `<IfModule mpm_prefork_module>
	StartServers		2
	MinSpareServers		2
	MaxSpareServers		2
	MaxRequestWorkers	{{ . }}
	MaxConnectionsPerChild	1024
</IfModule>`

// Apache struct
type Apache struct {
	totalMemory int
	avgProc     int
	maxProc     int
}

func init() {
	conf.Register("apache", &Apache{})
}

// TotalMemory is the max available totalMemory.
func (a *Apache) TotalMemory(m int) {
	a.totalMemory = m
}

// AvgProc is the average proc size.
func (a *Apache) AvgProc(pr int) {
	a.avgProc = pr
}

// MaxProc is the average proc size.
func (a *Apache) MaxProc(pr int) {
	a.maxProc = pr
}

// Build builds the template.
func (a *Apache) Build() (string, error) {
	// This is the number of concurrent processes that can be at a given time.
	maxClients := a.totalMemory / a.avgProc

	// We setup the templating with a special "divide" function, that way we can do inline division.
	fm := template.FuncMap{"divide": func(a, b int) int {
		return a / b
	}}
	t := template.Must(template.New(tpl).Funcs(fm).Parse(tpl))

	// Write the contents to totalMemory.
	buf := new(bytes.Buffer)
	err := t.Execute(buf, maxClients)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
