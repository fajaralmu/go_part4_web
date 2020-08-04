package app

import (
	"log"
	"reflect"
	"strings"

	"github.com/fajaralmu/go_part4_web/appConfig"
	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
	"github.com/fajaralmu/go_part4_web/repository"
)

var baseAdminMapping string = "/admin"

func initMenus() {

	log.Println("////////////////////MenuInitiationService INITIALIZE////////////////////")
	defaultAdminPage()
	defaultAboutPage()
	checkManagementPage()
	checkAdminMenus()

	getPersistenctEntities()
	log.Println("////////////////////END MENU INITIALIZE////////////////////")
}

func checkAdminMenus() {
	log.Println("STARTS checkAdminMenus")

	methods := []string{"/home", "/pagesettings"}

	for _, method := range methods {
		checkAdminMenu(baseAdminMapping, method)
	}

	log.Println("ENDS checkAdminMenus")
}

func constructAdminMenu(path string) entities.Menu {

	log.Println("constructAdminMenu LINK: ", path)
	menuCode := reflections.GetWordsAfterLastChar(path, "/")
	defAdminPage := defaultAdminPage()
	log.Println("defAdminPage ID", defAdminPage.ID)

	adminMenu := entities.Menu{
		Code:            menuCode,
		Color:           "#000000",
		BackgroundColor: "#ffffff",
		Description:     "Generated [" + menuCode + "]",
		Name:            strings.ToUpper(menuCode),
		URL:             path,
		PageID:          uint16(defAdminPage.ID),
		IconURL:         "DefaultIcon.bmp",
	}
	adminMenu.MenuPage = defAdminPage
	return adminMenu
}

func defaultAdminPage() *entities.Page {
	return getPage("admin", adminPage())
}

func managementPage() *entities.Page {
	return config_defaultManagementPage
}

func adminPage() *entities.Page {
	return config_defaultAdminPage
}

func checkManagementPage() {
	log.Println("STARTS _managementPage")
	_, ok := getPageOnlyByCode("management")
	if ok {
		log.Println("managementPage FOUND")
		return
	}

	log.Println("managementPage NOT FOUND. WILL ADD SETTING")
	_managementPage := managementPage()
	repository.CreateNewWithoutValidation(_managementPage)
	log.Println("STARTS _managementPage")
}

func defaultManagementPage() *entities.Page {

	return getPage("management", managementPage())
}

func getPersistenctEntities() []reflect.Type {
	log.Println("////////// START Management Page for PersistenceEntities")
	types := appConfig.GetEntitiesTypes()

	for _, t := range types {
		log.Println("------validateManagementPage: ", t.Name())
		validateManagementPage(t)
	}
	log.Println("////////// END SET Management Page for PersistenceEntities")
	return types
}

func validateManagementPage(t reflect.Type) {
	code := reflections.ToSnakeCase(t.Name(), true)
	_, ok := getMenuByCode(code)
	if !ok {
		addNewManagementMenuPageFor(t)
	}

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

func resetAllMenus() {
	allMenus := &[]entities.Menu{}
	menus, _ := repository.Filter(allMenus, entities.Filter{})

	log.Println("Will reset allMenus")
	for _, item := range menus {
		repository.Delete(item.(entities.Menu), false)
	}
	allPages := &[]entities.Page{}
	page, _ := repository.Filter(allPages, entities.Filter{})
	log.Println("Will reset allPages")
	for _, item := range page {
		repository.Delete(item.(entities.InterfaceEntity), false)
	}

	log.Println("////////END REMOVING/////////")
	initMenus()
	log.Println("//////////////END RESET//////////////")

}
