package model

type Entitlement struct {
	Identifier      string `json:"identifier"`
	Description     string `json:"description"`
	LongDescription string `json:"longDescription"`
	Weight          int    `json:"weight"`
}
