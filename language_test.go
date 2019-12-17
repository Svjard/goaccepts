package goaccepts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage(t *testing.T) {
	assert.Equal(t, preferredLanguages(nil), []string{}, "The two words should be the same.")
	assert.Equal(t, preferredLanguages("*"), []string{}, "The two words should be the same.")
	assert.Equal(t, preferredLanguages("*, en"), []string{"en"}, "The two words should be the same.")
	assert.Equal(t, preferredLanguages("*, en;q=0"), []string{"en"}, "The two words should be the same.")
}
