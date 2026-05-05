// Copyright (c) 2026 Arsenii Kvachan
// SPDX-License-Identifier: MIT

package main

import (
	"github.com/akvachan/medahra/include"
)

func main() {
	medahra.ConvertPositionsCSVToJSONL(
		"data/kaggle-positions.csv",
		"data/positions.jsonl",
	)
}
