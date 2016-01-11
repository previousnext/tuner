package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApacheProcs(t *testing.T) {
	procs, err := ApacheProcs(1000, 128)
	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, 7, procs, "Generated max procs for apache.")
	}
}
