package cmd

import "testing"

func TestFormatDate(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"2024-06", "06/2024"},
		{"1999-12", "12/1999"},
		{"", ""},
		{"Present", "Present"},
		{"2024", "2024"},
		{"2024-06-15", "2024-06-15"},
	}
	for _, c := range cases {
		if got := formatDate(c.in); got != c.want {
			t.Errorf("formatDate(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFormatEndDate(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", "Present"},
		{"2024-06", "06/2024"},
		{"Present", "Present"},
	}
	for _, c := range cases {
		if got := formatEndDate(c.in); got != c.want {
			t.Errorf("formatEndDate(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFormatReleaseDate(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"2024-03-15", "March 15, 2024"},
		{"2001-01-01", "January 1, 2001"},
		{"", ""},
		{"not a date", "not a date"},
		{"2024-03", "2024-03"},
	}
	for _, c := range cases {
		if got := formatReleaseDate(c.in); got != c.want {
			t.Errorf("formatReleaseDate(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
