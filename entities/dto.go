package entities

import "time"

//Filter is filter object
type Filter struct {
	EntityName   string
	Page         int
	Limit        int
	FieldsFilter map[string]interface{}
	Exacts       bool
	OrderBy      string
	OrderType    string
}

//WebRequest is REST API request payload
type WebRequest struct {
	Filter Filter

	//MODEL NAMES
	//convention:
	//1. Starts with uppercase
	//2. Snake case of model type name
	User_role          *UserRole
	User               *User
	Profile            *Profile
	Menu               *Menu
	Page               *Page
	Registered_request *RegisteredRequest
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
