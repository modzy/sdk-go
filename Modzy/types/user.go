package types

type KeyDetailResponse struct {
	Prefix         string   `json:"prefix,omitempty"`
	Name           string   `json:"name,omitempty"`
	Description    string   `json:"description,omitempty"`
	Isdefault      bool     `json:"isDefault,omitempty"`
	Ishighpriority bool     `json:"isHighPriority,omitempty"`
	Status         string   `json:"status,omitempty"`
	Createdat      string   `json:"createdAt,omitempty"`
	Submittedjobs  int      `json:"submittedJobs,omitempty"`
	Retrieved      bool     `json:"retrieved,omitempty"`
	Roles          []Roles  `json:"roles,omitempty"`
	Labels         []Labels `json:"labels,omitempty"`
	Account        Account  `json:"account,omitempty"`
	Team           Team     `json:"team,omitempty"`
	User           User     `json:"user,omitempty"`
}
type Entitlements struct {
	Identifier      string `json:"identifier,omitempty"`
	Scope           string `json:"scope,omitempty"`
	Description     string `json:"description,omitempty"`
	Weight          int    `json:"weight,omitempty"`
	Longdescription string `json:"longDescription,omitempty"`
}
type Roles struct {
	Id              int            `json:"id,omitempty"`
	Identifier      string         `json:"identifier,omitempty"`
	Name            string         `json:"name,omitempty"`
	Description     string         `json:"description,omitempty"`
	Longdescription string         `json:"longDescription,omitempty"`
	Weight          int            `json:"weight,omitempty"`
	Type            string         `json:"type,omitempty"`
	Entitlements    []Entitlements `json:"entitlements,omitempty"`
}
type Labels struct {
	Identifier string `json:"identifier,omitempty"`
	Name       string `json:"name,omitempty"`
}
type Account struct {
	Name     string `json:"name,omitempty"`
	Isactive bool   `json:"isActive,omitempty"`
}
type Team struct {
	Identifier  string `json:"identifier,omitempty"`
	Name        string `json:"name,omitempty"`
	Status      string `json:"status,omitempty"`
	Membercount int    `json:"memberCount,omitempty"`
}
type User struct {
	Identifier         string       `json:"identifier,omitempty"`
	Externalidentifier string       `json:"externalIdentifier,omitempty"`
	Firstname          string       `json:"firstName,omitempty"`
	Lastname           string       `json:"lastName,omitempty"`
	Email              string       `json:"email,omitempty"`
	Accesskeys         []Accesskeys `json:"accessKeys,omitempty"`
	Status             string       `json:"status,omitempty"`
	Title              string       `json:"title,omitempty"`
	Visited            bool         `json:"visited,omitempty"`
	Onboarded          bool         `json:"onboarded,omitempty"`
}
