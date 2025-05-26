package filters

import (
	"strings"
)

func IsSuspicious(text string, blacklist []string, whitelist []string) bool {
	lower := strings.ToLower(text)

	// Skip jika mengandung whitelist
	for _, w := range whitelist {
		if strings.Contains(lower, strings.ToLower(w)) {
			return false
		}
	}

	// Deteksi mention
	if strings.Contains(lower, "@") {
		return true
	}

	// Cek blacklist
	for _, bl := range blacklist {
		if strings.Contains(lower, strings.ToLower(bl)) {
			return true
		}
	}

	// Tambahan default keywords
	defaultKeywords := []string{"vcs", "tmo", "vcsan", "join gc"}
	for _, word := range defaultKeywords {
		if strings.Contains(lower, word) {
			return true
		}
	}

	return false
}
