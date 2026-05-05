// Copyright (c) 2026 Arsenii Kvachan
// SPDX-License-Identifier: MIT

package medahra

type Candidate struct {
	CandidateID            uint16              `json:"candidate_id"`
	About                  string              `json:"about"`
	ShortAbout             string              `json:"short_about"`
	GeneralSkills          []string            `json:"general_skills"`
	PreferredPositions     []string            `json:"preferred_positions"`
	PreferredSkills        []string            `json:"preferred_skills"`
	PreferredMinSalary     uint32              `json:"preferred_min_salary"`
	PreferredModes         []string            `json:"preferred_modes"`
	PreferredContractTypes []string            `json:"preferred_contract_types"`
	WorkingExperience      []WorkingExperience `json:"working_experience"`
	Qualifications         []Qualification     `json:"qualifications"`
	Languages              []Language          `json:"languages"`
}

type WorkingExperience struct {
	DateFrom      string   `json:"date_from"`
	DateTo        string   `json:"date_to"`
	DaysDuration  uint16   `json:"days_duration,omitempty"`
	HoursDuration uint16   `json:"hours_duration,omitempty"`
	Description   string   `json:"description"`
	Seniority     string   `json:"seniority"`
	Skills        []string `json:"skills"`
	Title         string   `json:"title"`
	Type          string   `json:"type"`
}

type Qualification struct {
	DateFrom            string   `json:"date_from"`
	DateTo              string   `json:"date_to"`
	Description         string   `json:"description"`
	HoursDuration       uint16   `json:"hours_duration"`
	Institution         string   `json:"institution"`
	Name                string   `json:"name"`
	NQF                 uint8    `json:"nqf"`
	QualificationCustom string   `json:"qualification_custom"`
	Skills              []string `json:"skills"`
}

type Language struct {
	CEFR     string `json:"cefr"`
	Language string `json:"language"`
}
