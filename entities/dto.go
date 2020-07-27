package entities

import "time"

//Filter is filter object
type Filter struct {
	EntityName   string
	Page         int
	Limit        int
	FieldsFilter map[string]interface{}
	Exact        bool
}

type WebRequest struct {
	Filter    Filter
	User_role *UserRole
	User      *User
	Profile   *Profile
	Menu      *Menu
	Page      *Page
}
type WebResponse struct {
	Date       time.Time
	Code       string
	Message    string
	ResultList interface{}
	Result     interface{}
	TotalData  int
}
