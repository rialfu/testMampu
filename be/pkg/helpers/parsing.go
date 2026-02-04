package helpers

import (
	"regexp"
	"strconv"
)

func parseInt(v []string, def int) int {
	if len(v) == 0 {
		return def
	}
	i, err := strconv.Atoi(v[0])
	if err != nil {
		return def
	}
	return i
}

func parseString(v []string) string {
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func IsEmailValid(e string) bool {
	// Regex standar untuk validasi email
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
