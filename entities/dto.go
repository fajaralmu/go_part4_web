package entities

//Filter is filter object
type Filter struct {
	EntityName   string
	Page         int
	Limit        int
	FieldsFilter map[string]interface{}
}

type WebRequest struct {
	Filter Filter
}
type WebResponse struct {
	ResultList interface{}
	TotalData  int
}
