package entities

//Filter is filter object
type Filter struct {
	Page         int
	Limit        int
	FieldsFilter map[string]interface{}
}
