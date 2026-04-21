package resume

import (
	_ "embed"
	"encoding/json"
)

//go:embed resume.json
var resumeData []byte

type Meta struct {
	LastModified string `json:"lastModified"`
}

type Resume struct {
	Basics       Basics        `json:"basics"`
	Certificates []Certificate `json:"certificates"`
	Education    []Education   `json:"education"`
	Languages    []Language    `json:"languages"`
	Meta         Meta          `json:"meta"`
	Projects     []Project     `json:"projects"`
	Skills       []Skill       `json:"skills"`
	Volunteer    []Volunteer   `json:"volunteer"`
	Work         []Work        `json:"work"`
	XMK          XMK           `json:"x-mk"`
}

type Basics struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Image    string    `json:"image"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	URL      string    `json:"url"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}

type Location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
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

type Project struct {
	Name        string   `json:"name"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	URL         string   `json:"url"`
}

type Certificate struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Issuer string `json:"issuer"`
	URL    string `json:"url"`
}

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type Work struct {
	Name       string   `json:"name"`
	Position   string   `json:"position"`
	URL        string   `json:"url"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type XMK struct {
	About  string  `json:"about"`
	Quotes []Quote `json:"quotes"`
	FAQs   []FAQ   `json:"faqs"`
}

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type FAQ struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func Load() (Resume, error) {
	var r Resume
	err := json.Unmarshal(resumeData, &r)
	return r, err
}
