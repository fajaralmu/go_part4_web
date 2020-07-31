package entities

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//InterfaceEntity is InterfaceEntity
type InterfaceEntity interface {
	Validate() interface{}
}

//BaseEntity is the entity
type BaseEntity struct {
	InterfaceEntity
	gorm.Model
	ID           int `custom:"type:FIELD_TYPE_TEXT"`
	CreatedDate  time.Time
	ModifiedDate time.Time
	Deleted      bool
	Color        string `custom:"type:FIELD_TYPE_COLOR;lableName:Background_Color;defaultValue:#ffffff"`
	FontColor    string `custom:"type:FIELD_TYPE_COLOR;defaultValue:#000000"`
}

//User is the entity
type User struct {
	InterfaceEntity
	gorm.Model
	Username    string    `gorm:"unique;not null" custom:"type:FIELD_TYPE_TEXT;emptyAble:FALSE"`
	DisplayName string    `gorm:"not null" custom:"type:FIELD_TYPE_TEXT;emptyAble:FALSE"`
	Password    string    `gorm:"not null" custom:"type:FIELD_TYPE_TEXT;emptyAble:FALSE"`
	Role        *UserRole `custom:"foreignKey:UserRoleID;foreignKeyType:number;type:FIELD_TYPE_FIXED_LIST;optionItemName:Name"`
	UserRoleID  uint16
}

//UserRole is the entity
type UserRole struct {
	InterfaceEntity
	gorm.Model
	Name   string `gorm:"unique" custom:"type:FIELD_TYPE_TEXT"`
	Access string `custom:"type:FIELD_TYPE_TEXT"`
	Code   string `gorm:"unique" custom:"type:FIELD_TYPE_TEXT"`
}

//Menu is the entity
type Menu struct {
	InterfaceEntity
	gorm.Model
	Code            string `gorm:"unique" custom:"type:FIELD_TYPE_TEXT"`
	Name            string `gorm:"unique" custom:"type:FIELD_TYPE_TEXT"`
	Description     string `custom:"type:FIELD_TYPE_TEXTAREA"`
	URL             string `gorm:"unique" custom:"type:FIELD_TYPE_TEXT"`
	MenuPage        *Page  `custom:"foreignKey:PageID;foreignKeyType:number;type:FIELD_TYPE_DYNAMIC_LIST;lableName:Page;optionItemName:Name"`
	PageID          uint16
	IconURL         string `custom:"type:FIELD_TYPE_IMAGE;required:FALSE;defaultValue:DefaultIcon.bmp"`
	Color           string `custom:"type:FIELD_TYPE_COLOR;lableName:Font_Color"`
	BackgroundColor string `custom:"type:FIELD_TYPE_COLOR"`
}

//Page is the entity
type Page struct {
	InterfaceEntity
	gorm.Model
	Code        string `gorm:"unique" custom:"type:FIELD_TYPE_TEXT"`
	Name        string `gorm:"unique" custom:"type:FIELD_TYPE_TEXT"`
	Authorized  int    `orm:"not null" custom:"type:FIELD_TYPE_PLAIN_LIST;lableName:Authorized_(yes_1_or_no_0);availableValues:0,1"`
	NONMenuPage int    `custom:"type:FIELD_TYPE_PLAIN_LIST;lableName:Is_Non_Menu_Page_(yes_1_or_no_0);availableValues:0,1"`
	Link        string `gorm:"unique" custom:"type:FIELD_TYPE_TEXT;lableName:Link_for_non_menu_page"`
	Description string `custom:"type:FIELD_TYPE_TEXTAREA"`
	ImageURL    string `custom:"type:FIELD_TYPE_IMAGE;multiple:true;required:FALSE;defaultValue:DefaultIcon.bmp"`
	Sequence    int    `custom:"type:FIELD_TYPE_NUMBER;lableName:Urutan_Ke"`
	Menus       []Menu `gorm:"-"`
}

//Profile is the entity
type Profile struct {
	InterfaceEntity
	gorm.Model
	Name             string `custom:"type:FIELD_TYPE_TEXT"`
	APPCode          string `gorm:"unique" custom:"type:FIELD_TYPE_HIDDEN"`
	ShortDescription string `custom:"type:FIELD_TYPE_TEXTAREA"`
	About            string `custom:"type:FIELD_TYPE_TEXTAREA"`
	WelcomingMessage string `custom:"type:FIELD_TYPE_TEXTAREA"`
	Address          string `custom:"type:FIELD_TYPE_TEXTAREA"`
	Contact          string `custom:"type:FIELD_TYPE_TEXTAREA"`
	Website          string `custom:"type:FIELD_TYPE_TEXT"`
	IconURL          string `custom:"type:FIELD_TYPE_IMAGE;required:FALSE;defaultValue:DefaultIcon.bmp"`
	BackgroundURL    string `custom:"type:FIELD_TYPE_IMAGE;required:FALSE;defaultValue:DefaultBackground.bmp"`
	Color            string `custom:"type:FIELD_TYPE_COLOR"`
	BackgroundColor  string `custom:"type:FIELD_TYPE_COLOR"`
}

//RegisteredRequest is the entity
type RegisteredRequest struct {
	InterfaceEntity
	gorm.Model
	RequestID string `custom:"type:FIELD_TYPE_TEXT"`
	Value     string `custom:"type:FIELD_TYPE_TEXT"`
	Referrer  string
	UserAgent string
	IPAddress string
}

//IMplementations//////////

// Validate validates model properties //
func (u BaseEntity) Validate() interface{} {
	return u
}

// Validate validates model properties //
func (u User) Validate() interface{} {
	fmt.Println("Validating User")
	if u.Role == nil {
		u.Role = &UserRole{}
	}
	return u
}

// Validate validates model properties //
func (u UserRole) Validate() interface{} {
	return u
}

// Validate validates model properties //
func (u Menu) Validate() interface{} {
	fmt.Println("Validating Menu")
	if u.MenuPage == nil {
		u.MenuPage = &Page{}
	}
	return u
}

// Validate validates model properties //
func (u Page) Validate() interface{} {
	return u
}

// Validate validates model properties //
func (u Profile) Validate() interface{} {
	return u
}

// Validate validates model properties //
func (u RegisteredRequest) Validate() interface{} {
	return u
}
