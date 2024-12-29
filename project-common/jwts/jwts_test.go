package jwts

import (
	"testing"
)

func TestParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzYwODM5NTYsInRva2VuIjoiMTAyMCJ9.5ppYb9YSNerk3jB_CA7RTixXFK8qNs66XfKsX6VGV-Q"
	ParseToken(tokenString, "msproject")
}
