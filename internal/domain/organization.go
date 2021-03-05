package domain

type Organization struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Site         string `json:"site"`
	Coordinates  string `json:"coordinates"`
	Office       string `json:"office"`
	DateCreation string `json:"date_creation"`
	DateUpdate   string `json:"date_update"`
	IsActive     bool   `json:"is_active"`
	IsEnable     bool   `json:"is_enable"`
}

type OrganizationDocuments struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	File           string
	OrganizationId int  `json:"organization_id"`
	IsSecure       bool `json:"is_secure"`
}

type UpdateOrganization struct {
	Id          int     `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	Site        *string `json:"site"`
	Coordinates *string `json:"coordinates"`
	Office      *string `json:"office"`
	IsActive    *string `json:"is_active"`
}
