package php

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const r = `memory_limit = 64M`

func TestPHPBuild(t *testing.T) {
	a := PHP{}
	a.Max(2048)
	a.Proc(64)
	b, _ := a.Build()
	assert.Equal(t, r, b, "Generated correct PHP configuration.")
}