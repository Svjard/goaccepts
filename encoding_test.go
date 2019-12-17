package goaccepts_test

import (
	"testing"

	"github.com/Svjard/goaccepts"

	"github.com/stretchr/testify/assert"
)

var testMessageE = "Expect encodings to match parsed header encodings"

func TestEncodings(t *testing.T) {
	assert.Equal(t, goaccepts.Encodings(nil), []string{"identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("*"), []string{"*", "identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("*, gzip"), []string{"*", "gzip", "identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("*, gzip;q=0"), []string{"*", "identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("*;q=0"), []string{}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("*;q=0, identity;q=1"), []string{"identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("identity"), []string{"identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("identity;q=0"), []string{}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("gzip"), []string{"gzip", "identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("gzip, compress;q=0"), []string{"gzip", "identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("gzip, deflate"), []string{"gzip", "deflate", "identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("gzip;q=0.8, deflate"), []string{"deflate", "gzip", "identity"}, testMessageE)
	assert.Equal(t, goaccepts.Encodings("gzip;q=0.8, identity;q=0.5, *;q=0.3"), []string{"gzip", "identity", "*"}, testMessageE)
}

func testCompareE(a, b []goaccepts.EncodingResult) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i].Encoding != b[i].Encoding {
			return false
		}

		if a[i].Weight != b[i].Weight {
			return false
		}
	}

	return true
}

func TestEncodingsDetails(t *testing.T) {
	asterisk := goaccepts.EncodingResult{
		Encoding: "*",
		Weight:   100,
	}
	gzip := goaccepts.EncodingResult{
		Encoding: "gzip",
		Weight:   1.0,
	}
	deflate := goaccepts.EncodingResult{
		Encoding: "deflate",
		Weight:   1.0,
	}
	identity := goaccepts.EncodingResult{
		Encoding: "identity",
		Weight:   1.0,
	}
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails(nil), []goaccepts.EncodingResult{identity}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("*"), []goaccepts.EncodingResult{asterisk, identity}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("*, gzip"), []goaccepts.EncodingResult{asterisk, gzip, identity}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("*, gzip;q=0"), []goaccepts.EncodingResult{asterisk, identity}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("*;q=0"), []goaccepts.EncodingResult{}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("*;q=0, identity;q=1"), []goaccepts.EncodingResult{identity}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("identity"), []goaccepts.EncodingResult{identity}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("identity;q=0"), []goaccepts.EncodingResult{}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("gzip"), []goaccepts.EncodingResult{gzip, identity}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("gzip, compress;q=0"), []goaccepts.EncodingResult{gzip, identity}), true, testMessageE)
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("gzip, deflate"), []goaccepts.EncodingResult{gzip, deflate, identity}), true, testMessageE)

	gzip8 := goaccepts.EncodingResult{
		Encoding: "gzip",
		Weight:   0.8,
	}
	identity8 := goaccepts.EncodingResult{
		Encoding: "identity",
		Weight:   0.8,
	}
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("gzip;q=0.8, deflate"), []goaccepts.EncodingResult{deflate, gzip8, identity8}), true, testMessageE)
	identity5 := goaccepts.EncodingResult{
		Encoding: "identity",
		Weight:   0.5,
	}
	asterisk3 := goaccepts.EncodingResult{
		Encoding: "*",
		Weight:   0.3,
	}
	assert.Equal(t, testCompareE(goaccepts.EncodingsDetails("gzip;q=0.8, identity;q=0.5, *;q=0.3"), []goaccepts.EncodingResult{gzip8, identity5, asterisk3}), true, testMessageE)
}
