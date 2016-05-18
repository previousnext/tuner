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
	p.Max(2048)
	p.Proc(64)
	b, _ := p.Build()
	assert.Equal(t, r, b, "Generated correct Passenger configuration.")
}
