package goaccepts_test

import (
	"testing"

	"github.com/Svjard/goaccepts"

	"github.com/stretchr/testify/assert"
)

var testMessageC = "Expect charsets to match parsed header charsets"

func TestCharsets(t *testing.T) {
	assert.Equal(t, goaccepts.Languages(nil), []string{"*"}, testMessageC)
	assert.Equal(t, goaccepts.Languages("*"), []string{"*"}, testMessageC)
	assert.Equal(t, goaccepts.Languages("*, UTF-8"), []string{"*", "UTF-8"}, testMessageC)
	assert.Equal(t, goaccepts.Languages("*, UTF-8;q=0"), []string{"*"}, testMessageC)
	assert.Equal(t, goaccepts.Languages("ISO-8859-1"), []string{"ISO-8859-1"}, testMessageC)
	assert.Equal(t, goaccepts.Languages("UTF-8;q=0"), []string{}, testMessageC)
	assert.Equal(t, goaccepts.Languages("UTF-8, ISO-8859-1"), []string{"UTF-8", "ISO-8859-1"}, testMessageC)
	assert.Equal(t, goaccepts.Languages("UTF-8;q=0.8, ISO-8859-1"), []string{"ISO-8859-1", "UTF-8"}, testMessageC)
}

func testCompareC(a, b []goaccepts.CharsetResult) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i].Charset != b[i].Charset {
			return false
		}

		if a[i].Weight != b[i].Weight {
			return false
		}
	}

	return true
}

func TestCharsetDetails(t *testing.T) {
	asterisk := goaccepts.CharsetResult{
		Charset: "*",
		Weight:  100,
	}
	utf8 := goaccepts.CharsetResult{
		Charset: "UTF-8",
		Weight:  1.0,
	}
	iso := goaccepts.CharsetResult{
		Charset: "ISO-8859-1",
		Weight:  1.0,
	}
	assert.Equal(t, testCompareC(goaccepts.CharsetsDetails(nil), []goaccepts.CharsetResult{asterisk}), true, testMessageC)
	assert.Equal(t, testCompareC(goaccepts.CharsetsDetails("*"), []goaccepts.CharsetResult{asterisk}), true, testMessageC)
	assert.Equal(t, testCompareC(goaccepts.CharsetsDetails("*, UTF-8"), []goaccepts.CharsetResult{asterisk, utf8}), true, testMessageC)
	assert.Equal(t, testCompareC(goaccepts.CharsetsDetails("*, UTF-8;q=0"), []goaccepts.CharsetResult{asterisk}), true, testMessageC)
	assert.Equal(t, testCompareC(goaccepts.CharsetsDetails("ISO-8859-1"), []goaccepts.CharsetResult{iso}), true, testMessageC)
	assert.Equal(t, testCompareC(goaccepts.CharsetsDetails("UTF-8;q=0"), []goaccepts.CharsetResult{}), true, testMessageC)
	assert.Equal(t, testCompareC(goaccepts.CharsetsDetails("UTF-8, ISO-8859-1"), []goaccepts.CharsetResult{utf8, iso}), true, testMessageC)
	utf88 := goaccepts.CharsetResult{
		Charset: "UTF-8",
		Weight:  0.8,
	}
	assert.Equal(t, testCompareC(goaccepts.CharsetsDetails("UTF-8;q=0.8, ISO-8859-1"), []goaccepts.CharsetResult{iso, utf88}), true, testMessageC)
}
