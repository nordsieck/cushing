package baton

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaton_Twirl(t *testing.T) {
	b := &bytes.Buffer{}
	baton := Baton{W: b}

	err := baton.Twirl()
	assert.Nil(t, err)
	assert.Equal(t, []byte{'-', backspace}, b.Bytes())

	err = baton.Twirl()
	assert.Nil(t, err)
	assert.Equal(t, []byte{'-', backspace, '/', backspace}, b.Bytes())
}

func TestPhantom_Twirl(t *testing.T) {
	b := &bytes.Buffer{}
	phantom := Phantom{}

	err := phantom.Twirl()
	assert.Nil(t, err)
	assert.Equal(t, []byte(nil), b.Bytes())
}
