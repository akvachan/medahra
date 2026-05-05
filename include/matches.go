// Copyright (c) 2026 Arsenii Kvachan
// SPDX-License-Identifier: MIT

package medahra

type Match struct {
	MatchID       string         `json:"match_id"`
	CandidateID   string         `json:"candidate_id"`
	PositionItems []PositionItem `json:"positions"`
	Methods       []Method       `json:"methods"`
	CreatedAt     int64          `json:"created_at"`
}

type PositionItem struct {
	PositionID uint32  `json:"position_id"`
	Score      float32 `json:"score"`
	Reasoning  string  `json:"reasoning"`
}

type Method struct {
	Type  string `json:"type"`
	Model string `json:"model"`
}
