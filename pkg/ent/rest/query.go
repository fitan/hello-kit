package ent

type AgeEQ struct {
	AgeEQ int `json:"AgeEQ" from:"AgeEQ"`
}

type AgeNEQ struct {
	AgeNEQ int `json:"AgeNEQ" from:"AgeNEQ"`
}

type AgeIn struct {
	AgeIn int `json:"AgeIn" from:"AgeIn"`
}

type AgeNotIn struct {
	AgeNotIn int `json:"AgeNotIn" from:"AgeNotIn"`
}

type AgeGT struct {
	AgeGT int `json:"AgeGT" from:"AgeGT"`
}

type AgeGTE struct {
	AgeGTE int `json:"AgeGTE" from:"AgeGTE"`
}

type AgeLT struct {
	AgeLT int `json:"AgeLT" from:"AgeLT"`
}

type AgeLTE struct {
	AgeLTE int `json:"AgeLTE" from:"AgeLTE"`
}

type NameEQ struct {
	NameEQ string `json:"NameEQ" from:"NameEQ"`
}

type NameNEQ struct {
	NameNEQ string `json:"NameNEQ" from:"NameNEQ"`
}

type NameIn struct {
	NameIn string `json:"NameIn" from:"NameIn"`
}

type NameNotIn struct {
	NameNotIn string `json:"NameNotIn" from:"NameNotIn"`
}

type NameGT struct {
	NameGT string `json:"NameGT" from:"NameGT"`
}

type NameGTE struct {
	NameGTE string `json:"NameGTE" from:"NameGTE"`
}

type NameLT struct {
	NameLT string `json:"NameLT" from:"NameLT"`
}

type NameLTE struct {
	NameLTE string `json:"NameLTE" from:"NameLTE"`
}

type NameContains struct {
	NameContains string `json:"NameContains" from:"NameContains"`
}

type NameHasPrefix struct {
	NameHasPrefix string `json:"NameHasPrefix" from:"NameHasPrefix"`
}

type NameHasSuffix struct {
	NameHasSuffix string `json:"NameHasSuffix" from:"NameHasSuffix"`
}

type NameEqualFold struct {
	NameEqualFold string `json:"NameEqualFold" from:"NameEqualFold"`
}

type NameContainsFold struct {
	NameContainsFold string `json:"NameContainsFold" from:"NameContainsFold"`
}
