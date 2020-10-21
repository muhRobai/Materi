package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	resp, err := getValue("omama")
	assert.Equal(t, "you are not coder", err.Error())
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "", resp)

	resp, err = getValue("eat")
	assert.Equal(t, nil, err)
	assert.Equal(t, "eat", resp)

	resp, err = getValue("code")
	assert.Equal(t, nil, err)
	assert.Equal(t, "code", resp)

	resp, err = getValue("sleep")
	assert.Equal(t, nil, err)
	assert.Equal(t, "sleep", resp)
}
