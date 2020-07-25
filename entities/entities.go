package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

//InterfaceEntity is Entity Interface
type InterfaceEntity interface {
}

//BaseEntity is the entity
type BaseEntity struct {
	InterfaceEntity
	gorm.Model
	ID           int
	CreatedDate  time.Time
	ModifiedDate time.Time
	Deleted      bool
	Color        string
	FontColor    string
}

//User is the entity
type User struct {
	InterfaceEntity
	gorm.Model
	Username    string   `gorm:"unique;not null"`
	DisplayName string   `gorm:"not null"`
	Password    string   `gorm:"not null"`
	Role        UserRole `gorm:"foreignkey:role_id"`
}

//UserRole is the entity
type UserRole struct {
	InterfaceEntity
	gorm.Model
	Name   string `gorm:"unique"`
	Access string
	Code   string `gorm:"unique"`
}

//Menu is the entity
type Menu struct {
	InterfaceEntity
	gorm.Model
	Code        string `gorm:"unique"`
	Name        string `gorm:"unique"`
	Description string
	URL         string `gorm:"unique"`
	MenuPage    Page
	IconURL     string
}

//Page is the entity
type Page struct {
	InterfaceEntity
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

//Profile is the entity
type Profile struct {
	InterfaceEntity
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

//RegisteredRequest is the entity
type RegisteredRequest struct {
	InterfaceEntity
	gorm.Model
	RequestID string
	Value     string
	Referrer  string
	UserAgent string
	IPAddress string
}
