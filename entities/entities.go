package entities

import (
	"github.com/jinzhu/gorm"
)

//User is the entity
type User struct {
	gorm.Model
	Username    string `gorm:"unique;not null"`
	DisplayName string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Role        UserRole
}

func (User) TableName() string {
	return "user"
}

//UserRole is the entity
type UserRole struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Access string
	Code   string `gorm:"unique"`
}

func (UserRole) TableName() string {
	return "user_role"
}

//Menu is the entity
type Menu struct {
	gorm.Model
	Code        string `gorm:"unique"`
	Name        string `gorm:"unique"`
	Description string
	URL         string `gorm:"unique"`
	MenuPage    Page
	IconURL     string
}

func (Menu) TableName() string {
	return "menu"
}

//Page is the entity
type Page struct {
	gorm.Model
	Code        string `gorm:"unique"`
	Name        string `gorm:"unique"`
	Authorized  int    `gorm:"not null"`
	NONMenuPage int
	Link        string `gorm:"unique"`
	Description string
	ImageURL    string
	Sequence    int
}

func (Page) TableName() string {
	return "page"
}

//Profile is the entity
type Profile struct {
	gorm.Model
	Name             string
	APPCode          string `gorm:"unique"`
	ShortDescription string
	About            string
	WelcomingMessage string
	Address          string
	Contact          string
	Website          string
	IconURL          string
	BackgroundURL    string
}

func (Profile) TableName() string {
	return "profile"
}

//RegisteredRequest is the entity
type RegisteredRequest struct {
	gorm.Model
	RequestID string
	Value     string
	Referrer  string
	UserAgent string
	IPAddress string
}

func (RegisteredRequest) TableName() string {
	return "registered_request"
}
