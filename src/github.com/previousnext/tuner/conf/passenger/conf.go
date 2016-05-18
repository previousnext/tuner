package passenger

import (
	"bytes"
	"text/template"

	"github.com/previousnext/tuner/conf"
)

const tpl = `<IfModule mod_passenger.c>
    PassengerMaxPoolSize {{ . }}
</IfModule>`

type Passenger struct {
	max  int
	proc int
}

func init() {
	conf.Register("passenger", &Passenger{})
}

func (a *Passenger) Max(m int) {
	a.max = m
}

func (a *Passenger) Proc(pr int) {
	a.proc = pr
}

func (a *Passenger) Build() (string, error) {
	// This is the number of concurrent processes that can be at a given time.
	maxClients := a.max / a.proc

	// We setup the templating with a special "divide" function, that way we can do inline division.
	t := template.Must(template.New(tpl).Parse(tpl))

	// Write the contents to memory.
	buf := new(bytes.Buffer)
	err := t.Execute(buf, maxClients)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
