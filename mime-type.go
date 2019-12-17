package goaccepts

import (
	"sort"
	"strconv"
	"strings"
)

// MimeTypeResult of the accepts header
type MimeTypeResult struct {
	// Type ...
	Type string
	// ParentType ...
	ParentType string
	// ChildType
	ChildType string
	// Level ...
	Level int64
	// Quality ...
	Quality float64
	// i ...
	i int
}

func parseAcceptMimeType(accept string) []MimeTypeResult {
	res := []MimeTypeResult{}
	accepts := strings.Split(accept, ",")

	j := 0
	for i := 0; i < len(accepts); i++ {
		mimeType := parseMimeType(strings.Trim(accepts[i], " "), i)
		if mimeType != nil {
			n := findM(res, mimeType.Type)
			if n == -1 {
				res = append(res[:j], append([]MimeTypeResult{*mimeType}, res[j:]...)...)
				j = j + 1
			}
		}
	}

	return res
}

func parseMimeType(str string, i int) *MimeTypeResult {
	if str == "*/*" {
		return &MimeTypeResult{
			Type:       str,
			ParentType: "*",
			ChildType:  "*",
			Quality:    100.0,
			i:          i,
		}
	}

	quality := 1.0
	var level int64 = int64(0)
	params := strings.Split(str, ";")
	mimeType := strings.ToLower(params[0])
	if strings.IndexByte(mimeType, '"') == 0 && strings.LastIndexByte(mimeType, '"') == len(mimeType)-1 {
		mimeType = strings.Trim(mimeType, "\"")
	}
	typeParts := strings.Split(mimeType, "/")
	childType := ""
	if len(typeParts) > 1 {
		childType = typeParts[1]
	}

	if len(params) > 1 {
		for n := 1; n < len(params); n++ {
			p := strings.Split(params[n], "=")
			if p[0] == "q" {
				quality, _ = strconv.ParseFloat(p[1], 64)
			}

			if strings.ToLower(p[0]) == "level" {
				level, _ = strconv.ParseInt(p[1], 10, 64)
			}
		}
	}

	if level > 0 {
		return &MimeTypeResult{
			Type:       mimeType,
			ParentType: typeParts[0],
			ChildType:  childType,
			Level:      level,
			Quality:    quality,
			i:          i,
		}
	}

	return &MimeTypeResult{
		Type:       mimeType,
		ParentType: typeParts[0],
		ChildType:  childType,
		Quality:    quality,
		i:          i,
	}
}

func parseMimeTypes(accept interface{}) []MimeTypeResult {
	// no header = */*
	val := "*/*"
	if accept != nil {
		val = accept.(string)
	}

	accepts := parseAcceptMimeType(val)

	// sorted list of all languages
	filteredList := filterM(accepts, func(spec MimeTypeResult) bool {
		return spec.Quality > 0
	})
	sort.SliceStable(filteredList, func(i, j int) bool {
		return filteredList[i].Quality > filteredList[j].Quality
	})

	return filteredList
}

// MimeTypes - Get the list of accepted mime types from an Accept header.
// https://tools.ietf.org/html/rfc2616#section-14.1
func MimeTypes(accept interface{}) []string {
	charsets := parseMimeTypes(accept)

	fullList := mapM(charsets, func(spec MimeTypeResult) string {
		return spec.Type
	})

	return fullList
}

// MimeTypesDetails - Get the list of accepted mime types from an Accept header.
// https://tools.ietf.org/html/rfc2616#section-14.1
func MimeTypesDetails(accept interface{}) []MimeTypeResult {
	return parseMimeTypes(accept)
}

func filterM(vs []MimeTypeResult, f func(MimeTypeResult) bool) []MimeTypeResult {
	vsf := make([]MimeTypeResult, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func mapM(vs []MimeTypeResult, f func(MimeTypeResult) string) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		vsf = append(vsf, f(v))
	}
	return vsf
}

func findM(a []MimeTypeResult, x string) int {
	for i, n := range a {
		if x == n.Type {
			return i
		}
	}
	return -1
}
