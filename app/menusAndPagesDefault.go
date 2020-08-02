package app

import "github.com/fajaralmu/go_part4_web/entities"

var config_defaultManagementMenu *entities.Menu = &entities.Menu{
	Code:            "menu",
	Name:            "Menu Management",
	URL:             "/management/menu",
	Color:           "#ffffff",
	BackgroundColor: "#000000",
	Description:     "[Generated] Menu Management",
}
var config_defaultPageManagementMenu *entities.Menu = &entities.Menu{
	Code:            "page",
	Name:            "Page Management",
	URL:             "/management/page",
	Color:           "#ffffff",
	BackgroundColor: "#000000",
	Description:     "[Generated] Page Management",
}
var config_defaultPageSettingMenu *entities.Menu = &entities.Menu{
	Code:            "pagesettings",
	Name:            "Page Setting",
	URL:             "/admin/pagesettings",
	Color:           "#ffffff",
	BackgroundColor: "#000000",
	Description:     "[Generated] Page Setting",
}

var config_defaultSettingPage *entities.Page = &entities.Page{
	Code:        "setting",
	Name:        "Setting",
	Description: "[Generated] Setting Page",
	Link:        "/page/setting",
	NONMenuPage: 0,
	Authorized:  1,
}
var config_defaultManagementPage *entities.Page = &entities.Page{
	Code:        "management",
	Name:        "Management",
	Description: "[Generated] Management Page",
	Link:        "/page/management",
	NONMenuPage: 0,
	Authorized:  1,
}

var config_defaultAdminPage *entities.Page = &entities.Page{
	Code:        "admin",
	Name:        "Admin",
	Description: "[Generated] Admin Page",
	Link:        "/page/admin",
	NONMenuPage: 0,
	Authorized:  1,
}

var config_defaultAboutPage *entities.Page = &entities.Page{
	Code:        "about",
	Name:        "About Us",
	Description: "[Generated] About Us",
	Link:        "/public/about",
	NONMenuPage: 1,
	Authorized:  0,
}
