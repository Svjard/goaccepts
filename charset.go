package goaccepts

import (
	"sort"
	"strconv"
	"strings"
)

// CharsetResult of the accepts header
type CharsetResult struct {
	// Charset ...
	Charset string
	// Weight ...
	Weight float64
	// i ...
	i int
}

func parseAcceptCharset(accept string) []CharsetResult {
	res := []CharsetResult{}
	accepts := strings.Split(accept, ",")

	j := 0
	for i := 0; i < len(accepts); i++ {
		charset := parseCharset(strings.Trim(accepts[i], " "), i)
		if charset != nil {
			n := findC(res, charset.Charset)
			if n == -1 {
				res = append(res[:j], append([]CharsetResult{*charset}, res[j:]...)...)
				j = j + 1
			}
		}
	}

	return res
}

func parseCharset(str string, i int) *CharsetResult {
	if str == "*" {
		return &CharsetResult{
			Charset: str,
			Weight:  100.0,
			i:       i,
		}
	}

	weight := 1.0
	params := strings.Split(str, ";")
	if len(params) > 1 {
		p := strings.Split(params[1], "=")
		if strings.ToLower(p[0]) == "q" {
			weight, _ = strconv.ParseFloat(p[1], 64)
		}
	}

	charset := params[0]
	return &CharsetResult{
		Charset: charset,
		Weight:  weight,
		i:       i,
	}
}

func parseCharsets(accept interface{}) []CharsetResult {
	// no header = *
	val := "*"
	if accept != nil {
		val = accept.(string)
	}

	accepts := parseAcceptCharset(val)

	// sorted list of all languages
	filteredList := filterC(accepts, func(spec CharsetResult) bool {
		return spec.Weight > 0
	})
	sort.SliceStable(filteredList, func(i, j int) bool {
		return filteredList[i].Weight > filteredList[j].Weight
	})

	return filteredList
}

// Charsets - Get the list of accepted charsets from an Accept-Charset header.
// https://tools.ietf.org/html/rfc2616#section-14.4
func Charsets(accept interface{}) []string {
	charsets := parseCharsets(accept)

	fullList := mapC(charsets, func(spec CharsetResult) string {
		return spec.Charset
	})

	return fullList
}

// CharsetsDetails - Get the list of accepted charsets from an Accept-Charset header.
// https://tools.ietf.org/html/rfc2616#section-14.2
func CharsetsDetails(accept interface{}) []CharsetResult {
	return parseCharsets(accept)
}

func filterC(vs []CharsetResult, f func(CharsetResult) bool) []CharsetResult {
	vsf := make([]CharsetResult, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func mapC(vs []CharsetResult, f func(CharsetResult) string) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		vsf = append(vsf, f(v))
	}
	return vsf
}

func findC(a []CharsetResult, x string) int {
	for i, n := range a {
		if x == n.Charset {
			return i
		}
	}
	return -1
}
