package apache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const r = `<IfModule mpm_prefork_module>
	StartServers		2
	MinSpareServers		2
	MaxSpareServers		2
	MaxRequestWorkers	32
	MaxConnectionsPerChild	1024
</IfModule>`

func TestApacheBuild(t *testing.T) {
	a := Apache{}
	a.TotalMemory(2048)
	a.AvgProc(64)
	a.MaxProc(64)
	b, _ := a.Build()
	assert.Equal(t, r, b, "Generated correct Apache configuration.")
}
