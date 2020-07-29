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

//WebRequest is REST API request payload
type WebRequest struct {
	Filter    Filter
	User_role *UserRole
	User      *User
	Profile   *Profile
	Menu      *Menu
	Page      *Page
}

//WebResponse is REST API response payload
type WebResponse struct {
	Date           time.Time
	Code           string
	Message        string
	ResultList     interface{}
	Result         interface{}
	TotalData      int
	AdditionalData interface{}
	Filter         Filter
}
