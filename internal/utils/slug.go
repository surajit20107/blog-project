package utils

import (
	"github.com/google/uuid"
	"regexp"
	"strings"
)

func GenerateSlug(str string) string {
	s := strings.ToLower(str)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	s = re.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if s == "" {
		s = "post"
	}
	uuidPart := uuid.New().String()[:8]
	return s + "-" + uuidPart
}
