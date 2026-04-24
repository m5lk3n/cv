package cmd

import (
	"strings"
	"time"
)

// formatDate converts a YYYY-MM date from resume.json to MM/YYYY for display.
// Non-matching values (e.g., "Present", "") are returned unchanged.
func formatDate(s string) string {
	if len(s) == 7 && s[4] == '-' {
		year, month := s[:4], s[5:]
		var b strings.Builder
		b.Grow(7)
		b.WriteString(month)
		b.WriteByte('/')
		b.WriteString(year)
		return b.String()
	}
	return s
}

// formatEndDate formats an end date like formatDate, but renders an
// empty value as "Present" for ongoing entries.
func formatEndDate(s string) string {
	if s == "" {
		return "Present"
	}
	return formatDate(s)
}

// formatReleaseDate converts a YYYY-MM-DD date from resume.json to "January 2, 2006" format for display.
// Non-matching values are returned unchanged.
func formatReleaseDate(s string) string {
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t.Format("January 2, 2006")
	}
	return s
}
