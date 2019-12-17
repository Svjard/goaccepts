package goaccepts

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// AcceptsDef of the accepts header
type AcceptsDef struct {
	encodings []string
	languages []string
	types     []string
	charsets  []string
}

// LanguageResult of the accepts header
type LanguageResult struct {
	prefix string
	suffix string
	q      float64
	i      int
	full   string
}

// LanguagePriority ...
type LanguagePriority struct {
	i int
	o int
	q float64
	s int
}

var simpleLanguageRegExp = "^\\s*([^\\s\\-;]+)(?:-([^\\s;]+))?\\s*(?:;(.*))?$"

/**
 * Parse the Accept-Language header.
 */
func parseAcceptLanguage(accept string) []LanguageResult {
	res := []LanguageResult{}
	accepts := strings.Split(accept, ",")

	fmt.Println("parseAcceptLanguage", accepts)

	j := 0
	for i := 0; i < len(accepts); i++ {
		fmt.Println("1", accepts[i])
		language := parseLanguage(strings.Trim(accepts[i], " "), i)
		fmt.Println("2", language)
		if language != nil {
			res = append(res[:j], append([]LanguageResult{*language}, res[j:]...)...)
			j = j + 1
		}
	}

	return res
}

// parseLanguage parse a language from the Accept-Language header.
func parseLanguage(str string, i int) *LanguageResult {
	if str == "*" {
		return nil
	}

	r, _ := regexp.Compile(simpleLanguageRegExp)
	match := r.FindAllString(str, -1)
	fmt.Println("m", match)

	if match == nil {
		return nil
	}

	prefix := ""
	if len(match) > 0 {
		prefix = match[0]
	}

	suffix := ""
	if len(match) > 1 {
		suffix = match[1]
	}

	full := prefix

	if suffix != "" {
		full += "-" + suffix
	}

	q := 1.0
	if len(match) > 2 {
		params := strings.Split(match[2], ";")
		for j := 0; j < len(params); j++ {
			p := strings.Split(params[j], "=")
			if p[0] == "q" {
				q, _ = strconv.ParseFloat(p[1], 64)
			}
		}
	}

	fmt.Println("m", prefix, suffix, q, i, full)
	return &LanguageResult{
		prefix: prefix,
		suffix: suffix,
		q:      q,
		i:      i,
		full:   full,
	}
}

// specify Get the specificity of the language.
func specify(language string, spec LanguageResult, index int) *LanguagePriority {
	p := parseLanguage(language, index)
	if p == nil {
		return nil
	}

	s := 0
	if strings.ToLower(spec.full) == strings.ToLower(p.full) {
		s |= 4
	} else if strings.ToLower(spec.prefix) == strings.ToLower(p.full) {
		s |= 2
	} else if strings.ToLower(spec.full) == strings.ToLower(p.prefix) {
		s |= 1
	} else if spec.full != "*" {
		return nil
	}

	return &LanguagePriority{
		i: index,
		o: spec.i,
		q: spec.q,
		s: s,
	}
}

// preferredLanguages Get the preferred languages from an Accept-Language header.
func preferredLanguages(accept interface{}) []string {
	// RFC 2616 sec 14.4: no header = *
	val := "*"
	if accept != nil {
		val = accept.(string)
	}

	accepts := parseAcceptLanguage(val)

	// sorted list of all languages
	filteredList := Filter(accepts, func(spec LanguageResult) bool {
		return spec.q > 0
	})
	sort.SliceStable(filteredList, func(i, j int) bool {
		return filteredList[i].q < filteredList[j].q
	})
	fullList := Map(filteredList, func(spec LanguageResult) string {
		return spec.full
	})

	return fullList
}

// Filter ...
func Filter(vs []LanguageResult, f func(LanguageResult) bool) []LanguageResult {
	vsf := make([]LanguageResult, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map ...
func Map(vs []LanguageResult, f func(LanguageResult) string) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		vsf = append(vsf, f(v))
	}
	return vsf
}
