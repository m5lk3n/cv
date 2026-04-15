package resume

import (
	_ "embed"
	"encoding/json"
)

//go:embed resume.json
var resumeData []byte

type Resume struct {
	Education []Education `json:"education"`
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

func Load() (Resume, error) {
	var r Resume
	err := json.Unmarshal(resumeData, &r)
	return r, err
}
