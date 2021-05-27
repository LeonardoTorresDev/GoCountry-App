package utils

import (
	country "github.com/LTSpark/Country-App/internal/domain"
)

func ParseCountrySlice(countries []country.Country, skip int, limit int) ([]country.Country, int) {

	length := len(countries)

	if skip > length {
		skip = length
	}

	limit = skip + limit

	if limit > length {
		limit = length
	}

	return countries[skip:limit], length

}

func IntersectCountrySlices(s1, s2 []country.Country) (s3 []country.Country) {

	if s1 != nil && s2 != nil {
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

	if s1 == nil {
		return s2
	}

	return s1

}
