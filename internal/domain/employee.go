package domain

type OrganizationsUsers struct {
	Id             int  `json:"id"`
	OrganizationId int  `json:"organization_id"`
	UserId         int  `json:"user_id"`
	RoleId         int  `json:"role_id"`
	Confirmed      bool `json:"confirmed"`
}

type UpdateEmployee struct {
	Id             int  `json:"id"`
	OrganizationId int  `json:"organization_id"`
	UserId         int  `json:"user_id"`
	RoleId         *int `json:"role_id"`
	Confirmed      bool `json:"confirmed"`
}
