package main

//ResponseStruct ...
type ResponseStruct struct {
	Software Software  `json:"software"`
	Warnings Warnings  `json:"warnings"`
	Language Language  `json:"language"`
	Matches  []Matches `json:"matches"`
}

//Software ...
type Software struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	APIVersion  int    `json:"apiVersion"`
	Premium     bool   `json:"premium"`
	PremiumHint string `json:"premiumHint"`
	Status      string `json:"status"`
}

//Warnings ...
type Warnings struct {
	IncompleteResults bool `json:"incompleteResults"`
}

//Language ...
type Language struct {
	Name             string           `json:"name"`
	Code             string           `json:"code"`
	DetectedLanguage DetectedLanguage `json:"detectedLanguage"`
}

//DetectedLanguage ...
type DetectedLanguage struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

//Matches ...
type Matches struct {
	Message      string             `json:"message"`
	ShortMessage string             `json:"shortMessage"`
	Replacements []MatchReplacement `json:"replacements"`
	Offset       int                `json:"offset"`
	Length       int                `json:"length"`
	Context      MatchContext       `json:"context"`
	Sentence     string             `json:"sentence"`
	Type         MatchType          `json:"type"`
	Rule         MatchRule          `json:"rule"`
}

//MatchRule ...
type MatchRule struct {
	ID          string       `json:"id"`
	Description string       `json:"description"`
	IssueType   string       `json:"issueType"`
	Category    RuleCategory `json:"category"`
}

//RuleCategory ...
type RuleCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//MatchType ...
type MatchType struct {
	TypeName string `json:"typeName"`
}

//MatchReplacement ...
type MatchReplacement struct {
	Value string `json:"value"`
}

//MatchContext ...
type MatchContext struct {
	Text   string `json:"text"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}
