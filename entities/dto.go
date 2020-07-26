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
	Filter Filter
}
type WebResponse struct {
	Date       time.Time
	Code       string
	Message    string
	ResultList interface{}
	TotalData  int
}
