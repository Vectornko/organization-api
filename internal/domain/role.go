package domain

type Role struct {
	Id                 int    `json:"id"`
	OrganizationId     int    `json:"organization_id"`
	Name               string `json:"name"`
	EditOrganization   bool   `json:"edit_organization"`
	DeleteOrganization bool   `json:"delete_organization"`
	CreateService      bool   `json:"create_service"`
	EditService        bool   `json:"edit_service"`
	DeleteService      bool   `json:"delete_service"`
	CreateRole         bool   `json:"create_role"`
	EditRole           bool   `json:"edit_role"`
	DeleteRole         bool   `json:"delete_role"`
	CreateEmployee     bool   `json:"create_employee"`
	EditEmployee       bool   `json:"edit_employee"`
	DeleteEmployee     bool   `json:"delete_employee"`
}

type UpdateRole struct {
	Id                 int     `json:"id"`
	OrganizationId     int     `json:"organization_id"`
	Name               *string `json:"name"`
	EditOrganization   *bool   `json:"edit_organization"`
	DeleteOrganization *bool   `json:"delete_organization"`
	CreateService      *bool   `json:"create_service"`
	EditService        *bool   `json:"edit_service"`
	DeleteService      *bool   `json:"delete_service"`
	CreateRole         *bool   `json:"create_role"`
	EditRole           *bool   `json:"edit_role"`
	DeleteRole         *bool   `json:"delete_role"`
	CreateEmployee     *bool   `json:"create_employee"`
	EditEmployee       *bool   `json:"edit_employee"`
	DeleteEmployee     *bool   `json:"delete_employee"`
}
