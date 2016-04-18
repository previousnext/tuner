package clamd

import (
	"bytes"
	"text/template"

	"github.com/previousnext/tuner/conf"
)

const tpl = `<IfModule mpm_prefork_module>
	StartServers		{{ . }}
	MinSpareServers		{{ divide . 2 }}
	MaxSpareServers		{{ divide . 2 }}
	MaxRequestWorkers	{{ . }}
	MaxConnectionsPerChild  {{ . }}
</IfModule>`

type Apache struct {
	max  int
	proc int
}

func init() {
	conf.Register("apache", &Apache{})
}

func (a *Apache) Max(m int) {
	a.max = m
}

func (a *Apache) Proc(pr int) {
	a.proc = pr
}

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
