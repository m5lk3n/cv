package cmd

import "lttl.dev/cv/resume"

// loadResume is the indirection used by all commands to load resume data.
// Tests override it to inject a fixture.
var loadResume = resume.Load
