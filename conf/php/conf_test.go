package php

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const r = `memory_limit = 64M`

func TestPHPBuild(t *testing.T) {
	a := PHP{}
	a.TotalMemory(2048)
	a.AvgProc(64)
	a.MaxProc(64)
	b, _ := a.Build()
	assert.Equal(t, r, b, "Generated correct PHP configuration.")
}
