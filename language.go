package goaccepts

import (
	"sort"
	"strconv"
	"strings"
)

// LanguageResult of the accepts header
type LanguageResult struct {
	// prefix ...
	Prefix string
	// suffix ...
	Suffix string
	// weight ...
	Weight float64
	// i ...
	i int
	// full ...
	Full string
}

func parseAcceptLanguage(accept string) []LanguageResult {
	res := []LanguageResult{}
	accepts := strings.Split(accept, ",")

	j := 0
	for i := 0; i < len(accepts); i++ {
		language := parseLanguage(strings.Trim(accepts[i], " "), i)
		if language != nil {
			n := findL(res, language.Full)
			if n == -1 {
				res = append(res[:j], append([]LanguageResult{*language}, res[j:]...)...)
				j = j + 1
			}
		}
	}

	return res
}

func parseLanguage(str string, i int) *LanguageResult {
	if str == "*" {
		return &LanguageResult{
			Prefix: str,
			Suffix: "",
			Weight: 100.0,
			i:      i,
			Full:   str,
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

	langs := strings.Split(params[0], "-")
	prefix := langs[0]
	suffix := ""
	if len(langs) > 1 {
		suffix = strings.Join(langs[1:], "-")
	}
	full := params[0]

	return &LanguageResult{
		Prefix: prefix,
		Suffix: suffix,
		Weight: weight,
		i:      i,
		Full:   full,
	}
}

func parseLanguages(accept interface{}) []LanguageResult {
	// no header = *
	val := "*"
	if accept != nil {
		val = accept.(string)
	}

	accepts := parseAcceptLanguage(val)

	// sorted list of all languages
	filteredList := filterL(accepts, func(spec LanguageResult) bool {
		return spec.Weight > 0
	})
	sort.SliceStable(filteredList, func(i, j int) bool {
		return filteredList[i].Weight > filteredList[j].Weight
	})

	return filteredList
}

// Languages - Get the preferred languages from an Accept-Language header.
// https://tools.ietf.org/html/rfc2616#section-14.4
func Languages(accept interface{}) []string {
	languages := parseLanguages(accept)

	fullList := mapL(languages, func(spec LanguageResult) string {
		return spec.Full
	})

	return fullList
}

// LanguagesDetails - Get the preferred languages from an Accept-Language header.
// https://tools.ietf.org/html/rfc2616#section-14.4
func LanguagesDetails(accept interface{}) []LanguageResult {
	return parseLanguages(accept)
}

func filterL(vs []LanguageResult, f func(LanguageResult) bool) []LanguageResult {
	vsf := make([]LanguageResult, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func mapL(vs []LanguageResult, f func(LanguageResult) string) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		vsf = append(vsf, f(v))
	}
	return vsf
}

func findL(a []LanguageResult, x string) int {
	for i, n := range a {
		if x == n.Full {
			return i
		}
	}
	return -1
}
