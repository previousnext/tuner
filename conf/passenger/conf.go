package passenger

import (
	"bytes"
	"text/template"

	"github.com/previousnext/tuner/conf"
)

const tpl = `<IfModule mod_passenger.c>
    PassengerMaxPoolSize {{ . }}
</IfModule>`

// Passenger struct
type Passenger struct {
	totalMemory int
	avgProc     int
	maxProc     int
}

func init() {
	conf.Register("passenger", &Passenger{})
}

// TotalMemory is the total available totalMemory.
func (a *Passenger) TotalMemory(m int) {
	a.totalMemory = m
}

// AvgProc is the totalMemory avgProc size.
func (a *Passenger) AvgProc(pr int) {
	a.avgProc = pr
}

// MaxProc is the totalMemory avgProc size.
func (a *Passenger) MaxProc(pr int) {
	a.maxProc = pr
}

// Build builds the template.
func (a *Passenger) Build() (string, error) {
	// This is the number of concurrent processes that can be at a given time.
	maxClients := a.totalMemory / a.avgProc

	// We setup the templating with a special "divide" function, that way we can do inline division.
	t := template.Must(template.New(tpl).Parse(tpl))

	// Write the contents to totalMemory.
	buf := new(bytes.Buffer)
	err := t.Execute(buf, maxClients)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
