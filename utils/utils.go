package utils

import (
	countrycli "github.com/LTSpark/Country-App/internal"
)

func ParseSkipLimit(length int, skip int, limit int) (int, int) {

	if skip > length {
		skip = length
	}

	limit = skip + limit

	if limit > length {
		limit = length
	}

	return skip, limit

}

func IntersectCountrySlices(s1, s2 []countrycli.Country) (s3 []countrycli.Country) {

	m := make(map[string]bool)

	for _, value := range s1 {
		m[value.Name] = true
	}

	for _, value := range s2 {
		if _, ok := m[value.Name]; ok {
			s3 = append(s3, value)
		}
	}

	return
}
