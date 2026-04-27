package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"lttl.dev/cv/resume"
)

// useTestResume swaps loadResume to read from testdata/resume.json for the
// duration of the test.
func useTestResume(t *testing.T) {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "resume.json"))
	if err != nil {
		t.Fatalf("read testdata resume: %v", err)
	}
	orig := loadResume
	loadResume = func() (resume.Resume, error) {
		return resume.LoadFromBytes(data)
	}
	t.Cleanup(func() { loadResume = orig })
}

// runCmd invokes the root command with the given args and returns stdout+stderr.
func runCmd(t *testing.T, args ...string) string {
	t.Helper()
	return RunCommand(strings.Join(args, " "))
}

// assertContainsAll fails the test if any of substrings is missing from out.
func assertContainsAll(t *testing.T, out string, substrings ...string) {
	t.Helper()
	for _, s := range substrings {
		if !strings.Contains(out, s) {
			t.Errorf("output missing %q.\n--- output ---\n%s\n--- end ---", s, out)
		}
	}
}

// assertAliases checks that every alias produces the same output as the
// canonical command.
func assertAliases(t *testing.T, canonical string, aliases ...string) {
	t.Helper()
	want := runCmd(t, canonical)
	for _, alias := range aliases {
		got := runCmd(t, alias)
		if got != want {
			t.Errorf("alias %q output differs from %q.\n--- alias ---\n%s\n--- canonical ---\n%s",
				alias, canonical, got, want)
		}
	}
}

func TestAbout(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "about")
	assertContainsAll(t, out, "Test about summary line.")
	assertAliases(t, "about", "summary", "tldr", "info")
}

func TestBasics(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "basics")
	assertContainsAll(t, out,
		"Test User", "Test Engineer",
		"test@example.com", "+1-555-0100",
		"https://example.com",
		"Testville, TS, US",
		"GitHub: https://github.com/testuser",
	)
	assertAliases(t, "basics", "contact")
}

func TestCertificates(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "certificates")
	assertContainsAll(t, out,
		"Test Cert",
		"Issuer: Test Issuer",
		"Date: 06/2024",
		"https://example.com/cert",
	)
	assertAliases(t, "certificates",
		"certifications", "certs", "cert", "certificate", "certification",
		"licenses", "license", "lic",
	)
}

func TestEducation(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "education")
	assertContainsAll(t, out,
		"Test University",
		"Computer Science",
		"BSc",
		"09/2010 to 06/2014",
		"GPA: 4.0",
		"Algorithms, Databases",
		"https://test-uni.edu",
	)
	assertAliases(t, "education", "edu")
}

func TestFAQs(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "faqs")
	assertContainsAll(t, out, "Q: Why test?", "A: Because bugs.")
	assertAliases(t, "faqs", "faq")
}

func TestHashtags(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "hashtags")
	assertContainsAll(t, out, "#testing #golang")
	assertAliases(t, "hashtags", "hashtag", "tags")
}

func TestLanguages(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "languages")
	assertContainsAll(t, out,
		"English — Native",
		"Klingon — Beginner",
	)
	assertAliases(t, "languages", "language", "langs", "lang")
}

func TestProjects(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "projects")
	assertContainsAll(t, out,
		"Test Project",
		"01/2022 to 12/2023",
		"Project description.",
		"Did the test thing",
		"https://example.com/proj",
	)
	assertAliases(t, "projects", "proj")
}

func TestPublications(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "publications")
	assertContainsAll(t, out,
		"Test Article",
		"Publisher: Test Publisher",
		"Date: March 15, 2024",
		"https://example.com/pub",
		"Pub summary.",
	)
	assertAliases(t, "publications",
		"publication", "pubs", "pub", "articles", "article",
	)
}

func TestQuotes(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "quotes")
	assertContainsAll(t, out, `"Test all the things."`, "— Anon")
	assertAliases(t, "quotes", "quote")
}

func TestReferences(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "references")
	assertContainsAll(t, out, "Jane Doe", "Recommends them highly.")
	assertAliases(t, "references",
		"reference", "recommendations", "recommendation", "refs", "ref",
	)
}

func TestSkills(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "skills")
	assertContainsAll(t, out,
		"Go (Senior)",
		"concurrency, testing",
	)
	assertAliases(t, "skills", "skill")
}

func TestVolunteering(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "volunteering")
	assertContainsAll(t, out,
		"Test Org — Volunteer",
		"01/2018 to Present",
		"https://example.com/vol",
		"Volunteer summary.",
		"Helped out",
	)
	assertAliases(t, "volunteering", "volunteer", "vol")
}

func TestWork(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "work")
	assertContainsAll(t, out,
		"Test Co — Engineer",
		"01/2020 to 12/2023",
		"https://example.com/co",
		"Work summary.",
		"Shipped stuff",
		"Built things",
		"Older Co — Junior Engineer",
		"06/2015 to 12/2019",
		"Learned the ropes",
	)
	assertAliases(t, "work", "experience", "exp")
}

func TestAchievements(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "achievements")
	assertContainsAll(t, out,
		"Test Co:",
		"Shipped stuff",
		"Built things",
		"Older Co:",
		"Learned the ropes",
	)
	assertAliases(t, "achievements", "achieved", "achievement")
}

func TestWorkAchievementsSubcommand(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "work achievements")
	// Must match top-level `achievements` output.
	want := runCmd(t, "achievements")
	if out != want {
		t.Errorf("`work achievements` output differs from `achievements`.\n--- got ---\n%s\n--- want ---\n%s", out, want)
	}
}

func TestVersion(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "version")
	// In tests buildinfo defaults are used.
	assertContainsAll(t, out,
		"CV app v",
		"DEV",
		"CV data last modified on December 1, 2024",
	)
}

func TestUnknownCommand(t *testing.T) {
	useTestResume(t)
	out := runCmd(t, "this-is-not-a-command")
	if !strings.Contains(out, "unknown command") {
		t.Errorf("expected unknown command error, got:\n%s", out)
	}
}

func TestEmptyInput(t *testing.T) {
	if got := RunCommand(""); got != "" {
		t.Errorf("expected empty output for empty input, got %q", got)
	}
	if got := RunCommand("   "); got != "" {
		t.Errorf("expected empty output for whitespace input, got %q", got)
	}
}

func TestStripsLeadingCv(t *testing.T) {
	useTestResume(t)
	withPrefix := runCmd(t, "cv about")
	withoutPrefix := runCmd(t, "about")
	if withPrefix != withoutPrefix {
		t.Errorf("`cv about` should match `about`.\n--- with prefix ---\n%s\n--- without ---\n%s",
			withPrefix, withoutPrefix)
	}
}
