package goaccepts

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// EncodingResult of the accepts header
type EncodingResult struct {
	// Encoding ...
	Encoding string
	// weight ...
	Weight float64
	// i ...
	i int
}

func identityCheck(results []EncodingResult, encoding *EncodingResult) bool {
	index := findE(results, "*")

	if strings.ToLower(encoding.Encoding) == "identity" {
		return true
	}

	if index != -1 && results[index].Encoding == "*" && results[index].Weight == 0 {
		return true
	}

	return false
}

func parseAcceptEncoding(accept string) []EncodingResult {
	res := []EncodingResult{}
	accepts := strings.Split(accept, ",")
	hasIdentity := false
	minQuality := 1.0

	j := 0
	for i := 0; i < len(accepts); i++ {
		encoding := parseEncoding(strings.Trim(accepts[i], " "), i)
		if encoding != nil {
			n := findE(res, encoding.Encoding)
			if n == -1 {
				res = append(res[:j], append([]EncodingResult{*encoding}, res[j:]...)...)
				hasIdentity = hasIdentity || identityCheck(res, encoding)
				minQuality = math.Min(minQuality, encoding.Weight)
				j = j + 1
			}
		}
	}

	if hasIdentity == false {
		/*
		 * If identity doesn't explicitly appear in the accept-encoding header,
		 * it's added to the list of acceptable encoding with the lowest q
		 */
		identWeight := minQuality
		if minQuality == 0 {
			identWeight = 1.0
		}
		ident := &EncodingResult{
			Encoding: "identity",
			Weight:   identWeight,
			i:        j + 1,
		}
		res = append(res[:j], append([]EncodingResult{*ident}, res[j:]...)...)
	}

	return res
}

func parseEncoding(str string, i int) *EncodingResult {
	if str == "*" {
		return &EncodingResult{
			Encoding: str,
			Weight:   100.0,
			i:        i,
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

	encods := strings.Split(params[0], "-")
	return &EncodingResult{
		Encoding: encods[0],
		Weight:   weight,
		i:        i,
	}
}

func parseEncodings(accept interface{}) []EncodingResult {
	// no header = *
	val := "identity"
	if accept != nil {
		val = accept.(string)
	}

	accepts := parseAcceptEncoding(val)

	// sorted list of all languages
	filteredList := filterE(accepts, func(spec EncodingResult) bool {
		return spec.Weight > 0
	})
	sort.SliceStable(filteredList, func(i, j int) bool {
		return filteredList[i].Weight > filteredList[j].Weight
	})

	return filteredList
}

// Encodings - Get the encording from an Accept-Encoding header.
// https://tools.ietf.org/html/rfc2616#section-14.2
func Encodings(accept interface{}) []string {
	encodings := parseEncodings(accept)

	fullList := mapE(encodings, func(spec EncodingResult) string {
		return spec.Encoding
	})

	return fullList
}

// EncodingsDetails - Get the encording from an Accept-Encoding header.
// https://tools.ietf.org/html/rfc2616#section-14.2
func EncodingsDetails(accept interface{}) []EncodingResult {
	return parseEncodings(accept)
}

func filterE(vs []EncodingResult, f func(EncodingResult) bool) []EncodingResult {
	vsf := make([]EncodingResult, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func mapE(vs []EncodingResult, f func(EncodingResult) string) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		vsf = append(vsf, f(v))
	}
	return vsf
}

func findE(a []EncodingResult, x string) int {
	for i, n := range a {
		if x == n.Encoding {
			return i
		}
	}
	return -1
}
