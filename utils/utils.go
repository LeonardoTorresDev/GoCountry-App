package utils

import (
	country "github.com/LTSpark/Country-App/internal/domain"
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

func IntersectCountrySlices(s1, s2 []country.Country) (s3 []country.Country) {

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
