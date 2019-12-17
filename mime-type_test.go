package goaccepts_test

import (
	"testing"

	"github.com/Svjard/goaccepts"

	"github.com/stretchr/testify/assert"
)

var testMessageM = "Expect mime type to match parsed header mime type"

func TestMimeTypes(t *testing.T) {
	assert.Equal(t, goaccepts.MimeTypes(nil), []string{"*/*"}, testMessageM)
	assert.Equal(t, goaccepts.MimeTypes("*/*"), []string{"*/*"}, testMessageM)
	assert.Equal(t, goaccepts.MimeTypes("application/json"), []string{"application/json"}, testMessageM)
	assert.Equal(t, goaccepts.MimeTypes("application/json;q=0"), []string{}, testMessageM)
	assert.Equal(t, goaccepts.MimeTypes("application/json;q=0.2, text/html"), []string{"text/html", "application/json"}, testMessageM)
	assert.Equal(t, goaccepts.MimeTypes("text/*"), []string{"text/*"}, testMessageM)
	assert.Equal(t, goaccepts.MimeTypes("text/plain, application/json;q=0.5, text/html, */*;q=0.1"), []string{"text/plain", "text/html", "application/json", "*/*"}, testMessageM)
	assert.Equal(t, goaccepts.MimeTypes("text/plain, application/json;q=0.5, text/html, text/xml, text/yaml, text/javascript, text/csv, text/css, text/rtf, text/markdown, application/octet-stream;q=0.2, */*;q=0.1"), []string{"text/plain", "text/html", "text/xml", "text/yaml", "text/javascript", "text/csv", "text/css", "text/rtf", "text/markdown", "application/json", "application/octet-stream", "*/*"}, testMessageM)
}

func testCompareM(a, b []goaccepts.MimeTypeResult) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i].Type != b[i].Type {
			return false
		}

		if a[i].ParentType != b[i].ParentType {
			return false
		}

		if a[i].ChildType != b[i].ChildType {
			return false
		}

		if a[i].Level != b[i].Level {
			return false
		}

		if a[i].Quality != b[i].Quality {
			return false
		}
	}

	return true
}

func TestMimeTypesDetails(t *testing.T) {
	asterisk := goaccepts.MimeTypeResult{
		Type:       "*/*",
		ParentType: "*",
		ChildType:  "*",
		Quality:    100,
	}
	json := goaccepts.MimeTypeResult{
		Type:       "application/json",
		ParentType: "application",
		ChildType:  "json",
		Quality:    1.0,
	}
	html := goaccepts.MimeTypeResult{
		Type:       "text/html",
		ParentType: "text",
		ChildType:  "html",
		Quality:    1.0,
	}
	assert.Equal(t, testCompareM(goaccepts.MimeTypesDetails(nil), []goaccepts.MimeTypeResult{asterisk}), true, testMessageM)
	assert.Equal(t, testCompareM(goaccepts.MimeTypesDetails("*/*"), []goaccepts.MimeTypeResult{asterisk}), true, testMessageM)
	assert.Equal(t, testCompareM(goaccepts.MimeTypesDetails("application/json"), []goaccepts.MimeTypeResult{json}), true, testMessageM)
	assert.Equal(t, testCompareM(goaccepts.MimeTypesDetails("application/json;q=0"), []goaccepts.MimeTypeResult{}), true, testMessageM)
	json2 := goaccepts.MimeTypeResult{
		Type:       "application/json",
		ParentType: "application",
		ChildType:  "json",
		Quality:    0.2,
	}
	assert.Equal(t, testCompareM(goaccepts.MimeTypesDetails("application/json;q=0.2, text/html"), []goaccepts.MimeTypeResult{html, json2}), true, testMessageM)
	text := goaccepts.MimeTypeResult{
		Type:       "text/*",
		ParentType: "text",
		ChildType:  "*",
		Quality:    1.0,
	}
	assert.Equal(t, testCompareM(goaccepts.MimeTypesDetails("text/*"), []goaccepts.MimeTypeResult{text}), true, testMessageM)
	plain := goaccepts.MimeTypeResult{
		Type:       "text/plain",
		ParentType: "text",
		ChildType:  "plain",
		Quality:    1.0,
	}
	json5 := goaccepts.MimeTypeResult{
		Type:       "application/json",
		ParentType: "application",
		ChildType:  "json",
		Quality:    0.5,
	}
	asterisk1 := goaccepts.MimeTypeResult{
		Type:       "*/*",
		ParentType: "*",
		ChildType:  "*",
		Quality:    0.1,
	}
	assert.Equal(t, testCompareM(goaccepts.MimeTypesDetails("text/plain, application/json;q=0.5, text/html, */*;q=0.1"), []goaccepts.MimeTypeResult{plain, html, json5, asterisk1}), true, testMessageM)
	xml := goaccepts.MimeTypeResult{
		Type:       "text/xml",
		ParentType: "text",
		ChildType:  "xml",
		Quality:    1.0,
	}
	yaml := goaccepts.MimeTypeResult{
		Type:       "text/yaml",
		ParentType: "text",
		ChildType:  "yaml",
		Quality:    1.0,
	}
	javascript := goaccepts.MimeTypeResult{
		Type:       "text/javascript",
		ParentType: "text",
		ChildType:  "javascript",
		Quality:    1.0,
	}
	csv := goaccepts.MimeTypeResult{
		Type:       "text/csv",
		ParentType: "text",
		ChildType:  "csv",
		Quality:    1.0,
	}
	css := goaccepts.MimeTypeResult{
		Type:       "text/css",
		ParentType: "text",
		ChildType:  "css",
		Quality:    1.0,
	}
	rtf := goaccepts.MimeTypeResult{
		Type:       "text/rtf",
		ParentType: "text",
		ChildType:  "rtf",
		Quality:    1.0,
	}
	markdown := goaccepts.MimeTypeResult{
		Type:       "text/markdown",
		ParentType: "text",
		ChildType:  "markdown",
		Quality:    1.0,
	}
	stream2 := goaccepts.MimeTypeResult{
		Type:       "application/octet-stream",
		ParentType: "application",
		ChildType:  "octet-stream",
		Quality:    0.2,
	}
	assert.Equal(t, testCompareM(goaccepts.MimeTypesDetails("text/plain, application/json;q=0.5, text/html, text/xml, text/yaml, text/javascript, text/csv, text/css, text/rtf, text/markdown, application/octet-stream;q=0.2, */*;q=0.1"), []goaccepts.MimeTypeResult{plain, html, xml, yaml, javascript, csv, css, rtf, markdown, json5, stream2, asterisk1}), true, testMessageM)
}
