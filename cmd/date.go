package cmd

import "strings"

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
