package entities

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//InterfaceEntity is Entity Interface
type InterfaceEntity interface {
	Validate() interface{}
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

func (u BaseEntity) Validate() interface{} {
	return u
}

//User is the entity
type User struct {
	InterfaceEntity
	gorm.Model
	Username    string `gorm:"unique;not null" custom:"name:username"`
	DisplayName string `gorm:"not null"`
	Password    string `gorm:"not null"`
	RoleID      uint16
	Role        *UserRole `gorm:"foreignkey:role_id" custom:"foreignKey:RoleID"`
}

func (u User) Validate() interface{} {
	fmt.Println("Validating User")
	if u.Role == nil {
		u.Role = &UserRole{}
	}
	return u
}

//UserRole is the entity
type UserRole struct {
	InterfaceEntity
	gorm.Model
	Name   string `gorm:"unique"`
	Access string
	Code   string `gorm:"unique"`
}

func (u UserRole) Validate() interface{} {
	return u
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

func (u Menu) Validate() interface{} {
	return u
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

func (u Page) Validate() interface{} {
	return u
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

func (u Profile) Validate() interface{} {
	return u
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

func (u RegisteredRequest) Validate() interface{} {
	return u
}
