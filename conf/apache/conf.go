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
	max  int
	proc int
}

func init() {
	conf.Register("apache", &Apache{})
}

// Max is the max available memory.
func (a *Apache) Max(m int) {
	a.max = m
}

// Proc is the max proc size.
func (a *Apache) Proc(pr int) {
	a.proc = pr
}

// Build builds the template.
func (a *Apache) Build() (string, error) {
	// This is the number of concurrent processes that can be at a given time.
	maxClients := a.max / a.proc

	// We setup the templating with a special "divide" function, that way we can do inline division.
	fm := template.FuncMap{"divide": func(a, b int) int {
		return a / b
	}}
	t := template.Must(template.New(tpl).Funcs(fm).Parse(tpl))

	// Write the contents to memory.
	buf := new(bytes.Buffer)
	err := t.Execute(buf, maxClients)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
