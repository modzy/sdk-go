package model

type Entitlement struct {
	Identifier      string `json:"identifier"`
	Description     string `json:"description"`
	LongDescription string `json:"longDescription"`
	Weight          int    `json:"weight"`
}

type AccountingUser struct {
	Identifier         string    `json:"identifier"`
	FirstName          string    `json:"firstName"`
	LastName           string    `json:"lastName"`
	Email              string    `json:"email"`
	Status             string    `json:"status"`
	LastActiveDateTime ModzyTime `json:"lastActiveDateTime"`
	Visited            bool      `json:"visited"`
	Onboarded          bool      `json:"onboarded"`
}

type License struct {
	CompanyName       string `json:"companyName"`
	ProcessingEngines string `json:"processingEngines"`
}
