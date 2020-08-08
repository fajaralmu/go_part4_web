package app

import (
	"log"
	"reflect"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
	"github.com/fajaralmu/go_part4_web/repository"
)

var configDefaultManagementMenu *entities.Menu = &entities.Menu{
	Code:            "menu",
	Name:            "Menu Management",
	URL:             "/management/menu",
	Color:           "#ffffff",
	BackgroundColor: "#000000",
	Description:     "[Generated] Menu Management",
}
var configDefaultPageManagementMenu *entities.Menu = &entities.Menu{
	Code:            "page",
	Name:            "Page Management",
	URL:             "/management/page",
	Color:           "#ffffff",
	BackgroundColor: "#000000",
	Description:     "[Generated] Page Management",
}
var configDefaultPageSettingMenu *entities.Menu = &entities.Menu{
	Code:            "pagesettings",
	Name:            "Page Setting",
	URL:             "/admin/pagesettings",
	Color:           "#ffffff",
	BackgroundColor: "#000000",
	Description:     "[Generated] Page Setting",
}

var configDefaultSettingPage *entities.Page = &entities.Page{
	Code:        "setting",
	Name:        "Setting",
	Description: "[Generated] Setting Page",
	Link:        "/page/setting",
	NONMenuPage: 0,
	Authorized:  1,
}
var configDefaultManagementPage *entities.Page = &entities.Page{
	Code:        "management",
	Name:        "Management",
	Description: "[Generated] Management Page",
	Link:        "/page/management",
	NONMenuPage: 0,
	Authorized:  1,
}

var configDefaultAdminPage *entities.Page = &entities.Page{
	Code:        "admin",
	Name:        "Admin",
	Description: "[Generated] Admin Page",
	Link:        "/page/admin",
	NONMenuPage: 0,
	Authorized:  1,
}

var configDefaultAboutPage *entities.Page = &entities.Page{
	Code:        "about",
	Name:        "About Us",
	Description: "[Generated] About Us",
	Link:        "/public/about",
	NONMenuPage: 1,
	Authorized:  0,
}

func addNewManagementMenuPageFor(t reflect.Type) {
	log.Println("Will add default menu for: ", t.Name())

	commonPage := false //= dto.commonManagementPage();
	menuCode := reflections.ToSnakeCase(t.Name(), true)
	managementPage, _ := getPageOnlyByCode("management")
	menu := entities.Menu{
		Code:            menuCode,
		Name:            reflections.ExtractCamelCase(t.Name()),
		MenuPage:        &entities.Page{},
		PageID:          uint16(managementPage.ID),
		Color:           "#000000",
		BackgroundColor: "#ffffff",
		IconURL:         "DefaultIcon.bmp",
		Description:     "Generated Management Page for: " + t.Name(),
	}
	// menu.Validate()
	if commonPage {
		menu.URL = ("/management/common/" + menuCode)
	} else {
		menu.URL = ("/management/" + menuCode)
	}

	repository.CreateNewWithoutValidation(&menu)

	log.Println("Success Adding Management Menu For: ", menuCode)
}
