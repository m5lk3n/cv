package resume

import (
	_ "embed"
	"encoding/json"
)

//go:embed resume.json
var resumeData []byte

type Resume struct {
	Education  []Education `json:"education"`
	Languages  []Language  `json:"languages"`
	Volunteer  []Volunteer `json:"volunteer"`
}

type Education struct {
	Institution string   `json:"institution"`
	URL         string   `json:"url"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Score       string   `json:"score"`
	Courses     []string `json:"courses"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type Volunteer struct {
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	URL          string   `json:"url"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

func Load() (Resume, error) {
	var r Resume
	err := json.Unmarshal(resumeData, &r)
	return r, err
}
