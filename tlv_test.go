package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = "11AB398765UJ1A050200N23"

func TestMapearBasic(t *testing.T) {
	t.Run("Maper funciona correctamente", func(t *testing.T) {
		res, err := MapearTlv([]byte(s))
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})
	t.Run("Maper funciona erroneamente", func(t *testing.T) {
		res, err := MapearTlv([]byte(s + "asdf"))
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})
	t.Run("Maper sin parametros", func(t *testing.T) {
		res, err := MapearTlv(nil)
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "Sin parametros", err.Error())
	})
}
