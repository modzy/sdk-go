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

type AccountingProject struct {
	Identifier  string      `json:"identifier"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Status      string      `json:"status"`
	Visibility  string      `json:"visibility"`
	AccessKeys  []AccessKey `json:"accessKeys"`
	User        UserSummary `json"user"`
	CreatedAt   ModzyTime   `json:"createdAt"`
	UpdatedAt   ModzyTime   `json:"updatedAt"`
}

type AccessKey struct {
	Prefix         string    `json:"prefix"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	IsDefault      bool      `json:"isDefault"`
	IsHighPriority bool      `json:"isHighPriority"`
	Status         string    `json:"status"`
	CreatedAt      ModzyTime `json:"createdAt"`
}

type UserSummary struct {
	Identifier string `json:"identifier"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
}
