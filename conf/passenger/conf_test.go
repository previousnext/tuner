package passenger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const r = `<IfModule mod_passenger.c>
    PassengerMaxPoolSize 32
</IfModule>`

func TestPassengerBuild(t *testing.T) {
	p := Passenger{}
	p.TotalMemory(2048)
	p.AvgProc(64)
	p.MaxProc(128)
	b, _ := p.Build()
	assert.Equal(t, r, b, "Generated correct Passenger configuration.")
}
