// Copyright (c) 2026 Arsenii Kvachan
// SPDX-License-Identifier: MIT

package medahra

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
)

type Position struct {
	PositionID             uint32                  `json:"positionId"`
	Description            string                  `json:"description"`
	ShortDescription       string                  `json:"shortDescription"`
	Modes                  []string                `json:"modes"`
	ContractTypes          []string                `json:"contractTypes"`
	RequiredSkills         []string                `json:"requiredSkills"`
	RequiredQualifications []RequiredQualification `json:"requiredQualifications"`
	RequiredLanguages      []Language              `json:"requiredLanguages"`
	MinSalary              uint32                  `json:"minSalary"`
	MaxSalary              uint32                  `json:"maxSalary"`
	ExactSalary            uint32                  `json:"exactSalary"`
}

type RequiredQualification struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	NQF                 uint8    `json:"nqf"`
	QualificationCustom string   `json:"qualificationCustom"`
	Skills              []string `json:"skills"`
}

func ReadCSV(
	filePath string,
	handler func(record []string) error,
	reserveMB int,
) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReaderSize(f, reserveMB)

	csvReader := csv.NewReader(reader)
	csvReader.ReuseRecord = true

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := handler(record); err != nil {
			return err
		}
	}

	return nil
}

type PositionColumn uint8

const (
	PositionColumnRowID                    PositionColumn = 0
	PositionColumnJobID                    PositionColumn = 1
	PositionColumnCompanyName              PositionColumn = 2
	PositionColumnTitle                    PositionColumn = 3
	PositionColumnDescription              PositionColumn = 4
	PositionColumnMaxSalary                PositionColumn = 5
	PositionColumnPayPeriod                PositionColumn = 6
	PositionColumnLocation                 PositionColumn = 7
	PositionColumnCompanyID                PositionColumn = 8
	PositionColumnViews                    PositionColumn = 9
	PositionColumnMedSalary                PositionColumn = 10
	PositionColumnMinSalary                PositionColumn = 11
	PositionColumnFormattedWorkType        PositionColumn = 12
	PositionColumnApplies                  PositionColumn = 13
	PositionColumnOriginalListedTime       PositionColumn = 14
	PositionColumnRemoteAllowed            PositionColumn = 15
	PositionColumnJobPostingURL            PositionColumn = 16
	PositionColumnApplicationURL           PositionColumn = 17
	PositionColumnApplicationType          PositionColumn = 18
	PositionColumnExpiry                   PositionColumn = 19
	PositionColumnClosedTime               PositionColumn = 20
	PositionColumnFormattedExperienceLevel PositionColumn = 21
	PositionColumnSkillsDesc               PositionColumn = 22
	PositionColumnListedTime               PositionColumn = 23
	PositionColumnPostingDomain            PositionColumn = 24
	PositionColumnSponsored                PositionColumn = 25
	PositionColumnWorkType                 PositionColumn = 26
	PositionColumnCurrency                 PositionColumn = 27
	PositionColumnCompensationType         PositionColumn = 28
	PositionColumnNormalizedSalary         PositionColumn = 29
	PositionColumnZipCode                  PositionColumn = 30
	PositionColumnFips                     PositionColumn = 31
)

func ConvertPositionsCSVToJSONL(csvPath string, jsonlPath string) error {
	outFile, err := os.Create(jsonlPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriterSize(outFile, 100)
	defer writer.Flush()

	var positionID uint32 = 0
	return ReadCSV(csvPath, func(record []string) error {
		if record[0] == "" {
			return nil
		}

		positionID++

		position := Position{
			PositionID:  positionID,
			Description: record[PositionColumnDescription],
		}

		// save position
		line, err := json.Marshal(position)
		if err != nil {
			return err
		}
		if _, err := writer.Write(line); err != nil {
			return err
		}
		if _, err := writer.Write([]byte("\n")); err != nil {
			return err
		}

		return nil
	}, 100)
}
