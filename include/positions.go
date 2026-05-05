// Copyright (c) 2026 Arsenii Kvachan
// SPDX-License-Identifier: MIT

package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"strconv"
)

type Position struct {
	PositionID             uint16                  `json:"positionId"`
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
	PositionColumnJobID                    PositionColumn = 0
	PositionColumnCompanyName              PositionColumn = 1
	PositionColumnTitle                    PositionColumn = 2
	PositionColumnDescription              PositionColumn = 3
	PositionColumnMaxSalary                PositionColumn = 4
	PositionColumnPayPeriod                PositionColumn = 5
	PositionColumnLocation                 PositionColumn = 6
	PositionColumnCompanyID                PositionColumn = 7
	PositionColumnViews                    PositionColumn = 8
	PositionColumnMedSalary                PositionColumn = 9
	PositionColumnMinSalary                PositionColumn = 10
	PositionColumnFormattedWorkType        PositionColumn = 11
	PositionColumnApplies                  PositionColumn = 12
	PositionColumnOriginalListedTime       PositionColumn = 13
	PositionColumnRemoteAllowed            PositionColumn = 14
	PositionColumnJobPostingURL            PositionColumn = 15
	PositionColumnApplicationURL           PositionColumn = 16
	PositionColumnApplicationType          PositionColumn = 17
	PositionColumnExpiry                   PositionColumn = 18
	PositionColumnClosedTime               PositionColumn = 19
	PositionColumnFormattedExperienceLevel PositionColumn = 20
	PositionColumnSkillsDesc               PositionColumn = 21
	PositionColumnListedTime               PositionColumn = 22
	PositionColumnPostingDomain            PositionColumn = 23
	PositionColumnSponsored                PositionColumn = 24
	PositionColumnWorkType                 PositionColumn = 25
	PositionColumnCurrency                 PositionColumn = 26
	PositionColumnCompensationType         PositionColumn = 27
	PositionColumnNormalizedSalary         PositionColumn = 28
	PositionColumnZipCode                  PositionColumn = 29
	PositionColumnFips                     PositionColumn = 30
)

func ConvertPositionsCSVToJSONL(csvPath string, jsonlPath string) error {
	outFile, err := os.Create(jsonlPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriterSize(outFile, 100)
	defer writer.Flush()

	var positionID uint16 = 0
	return ReadCSV(csvPath, func(record []string) error {
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
