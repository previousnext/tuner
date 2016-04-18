package apache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const r = `<IfModule mpm_prefork_module>
	StartServers		32
	MinSpareServers		16
	MaxSpareServers		16
	MaxRequestWorkers	32
	MaxConnectionsPerChild	32
</IfModule>`

func TestApacheBuild(t *testing.T) {
	a := Apache{}
	a.Max(2048)
	a.Proc(64)
	b, _ := a.Build()
	assert.Equal(t, r, b, "Generated correct Apache configuration.")
}