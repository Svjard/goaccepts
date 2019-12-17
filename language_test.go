package goaccepts_test

import (
	"testing"

	"github.com/Svjard/goaccepts"

	"github.com/stretchr/testify/assert"
)

var testMessageL = "Expect languages to match parsed header languages"

func TestLanguages(t *testing.T) {
	assert.Equal(t, goaccepts.Languages(nil), []string{"*"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("*"), []string{"*"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("*, en"), []string{"*", "en"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("*, en;q=0"), []string{"*"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("*;q=0.8, en, es"), []string{"en", "es", "*"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("en"), []string{"en"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("en;q=0"), []string{}, testMessageL)
	assert.Equal(t, goaccepts.Languages("en;q=0.8, es"), []string{"es", "en"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("en;q=0.9, es;q=0.8, en;q=0.7"), []string{"en", "es"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("en-US, en;q=0.8"), []string{"en-US", "en"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("en-US, en-GB"), []string{"en-US", "en-GB"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("en-US;q=0.8, es"), []string{"es", "en-US"}, testMessageL)
	assert.Equal(t, goaccepts.Languages("nl;q=0.5, fr, de, en, it, es, pt, no, se, fi, ro"), []string{"fr", "de", "en", "it", "es", "pt", "no", "se", "fi", "ro", "nl"}, testMessageL)
}

func testCompareL(a, b []goaccepts.LanguageResult) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i].Prefix != b[i].Prefix {
			return false
		}

		if a[i].Suffix != b[i].Suffix {
			return false
		}

		if a[i].Weight != b[i].Weight {
			return false
		}

		if a[i].Full != b[i].Full {
			return false
		}
	}

	return true
}

func TestLanguageDetails(t *testing.T) {
	asterisk := goaccepts.LanguageResult{
		Prefix: "*",
		Suffix: "",
		Weight: 100,
		Full:   "*",
	}
	en := goaccepts.LanguageResult{
		Prefix: "en",
		Suffix: "",
		Weight: 1.0,
		Full:   "en",
	}
	es := goaccepts.LanguageResult{
		Prefix: "es",
		Suffix: "",
		Weight: 1.0,
		Full:   "es",
	}
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails(nil), []goaccepts.LanguageResult{asterisk}), true, testMessageL)
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("*"), []goaccepts.LanguageResult{asterisk}), true, testMessageL)
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("*, en"), []goaccepts.LanguageResult{asterisk, en}), true, testMessageL)
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("*, en;q=0"), []goaccepts.LanguageResult{asterisk}), true, testMessageL)

	asterisk8 := goaccepts.LanguageResult{
		Prefix: "*",
		Suffix: "",
		Weight: 0.8,
		Full:   "*",
	}
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("*;q=0.8, en, es"), []goaccepts.LanguageResult{en, es, asterisk8}), true, testMessageL)
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("en"), []goaccepts.LanguageResult{en}), true, testMessageL)
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("en;q=0"), []goaccepts.LanguageResult{}), true, testMessageL)
	en8 := goaccepts.LanguageResult{
		Prefix: "en",
		Suffix: "",
		Weight: 0.8,
		Full:   "en",
	}
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("en;q=0.8, es"), []goaccepts.LanguageResult{es, en8}), true, testMessageL)

	en9 := goaccepts.LanguageResult{
		Prefix: "en",
		Suffix: "",
		Weight: 0.9,
		Full:   "en",
	}
	es8 := goaccepts.LanguageResult{
		Prefix: "es",
		Suffix: "",
		Weight: 0.8,
		Full:   "es",
	}
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("en;q=0.9, es;q=0.8, en;q=0.7"), []goaccepts.LanguageResult{en9, es8}), true, testMessageL)
	enUs := goaccepts.LanguageResult{
		Prefix: "en",
		Suffix: "US",
		Weight: 1.0,
		Full:   "en-US",
	}
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("en-US, en;q=0.8"), []goaccepts.LanguageResult{enUs, en8}), true, testMessageL)
	enGb := goaccepts.LanguageResult{
		Prefix: "en",
		Suffix: "GB",
		Weight: 1.0,
		Full:   "en-GB",
	}
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("en-US, en-GB"), []goaccepts.LanguageResult{enUs, enGb}), true, testMessageL)
	enUs8 := goaccepts.LanguageResult{
		Prefix: "en",
		Suffix: "US",
		Weight: 0.8,
		Full:   "en-US",
	}
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("en-US;q=0.8, es"), []goaccepts.LanguageResult{es, enUs8}), true, testMessageL)
	nl5 := goaccepts.LanguageResult{
		Prefix: "nl",
		Suffix: "",
		Weight: 0.5,
		Full:   "nl",
	}
	fr := goaccepts.LanguageResult{
		Prefix: "fr",
		Suffix: "",
		Weight: 1.0,
		Full:   "fr",
	}
	de := goaccepts.LanguageResult{
		Prefix: "de",
		Suffix: "",
		Weight: 1.0,
		Full:   "de",
	}
	it := goaccepts.LanguageResult{
		Prefix: "it",
		Suffix: "",
		Weight: 1.0,
		Full:   "it",
	}
	pt := goaccepts.LanguageResult{
		Prefix: "pt",
		Suffix: "",
		Weight: 1.0,
		Full:   "pt",
	}
	no := goaccepts.LanguageResult{
		Prefix: "no",
		Suffix: "",
		Weight: 1.0,
		Full:   "no",
	}
	se := goaccepts.LanguageResult{
		Prefix: "se",
		Suffix: "",
		Weight: 1.0,
		Full:   "se",
	}
	fi := goaccepts.LanguageResult{
		Prefix: "fi",
		Suffix: "",
		Weight: 1.0,
		Full:   "fi",
	}
	ro := goaccepts.LanguageResult{
		Prefix: "ro",
		Suffix: "",
		Weight: 1.0,
		Full:   "ro",
	}
	assert.Equal(t, testCompareL(goaccepts.LanguagesDetails("nl;q=0.5, fr, de, en, it, es, pt, no, se, fi, ro"), []goaccepts.LanguageResult{fr, de, en, it, es, pt, no, se, fi, ro, nl5}), true, testMessageL)
}
