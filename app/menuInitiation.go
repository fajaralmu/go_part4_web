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

func initMenus() {

	log.Println("////////////////////MenuInitiationService INITIALIZE////////////////////")
	defaultAdminPage()
	defaultAboutPage()
	checkManagementPage()
	checkAdminMenus()

	getPersistenctEntities()
	log.Println("////////////////////END MENU INITIALIZE////////////////////")
}

func defaultAboutPage() *entities.Page {
	return getPage("about", config_defaultAboutPage)
}
func defaultSettingPage() *entities.Page {
	return getPage("setting", config_defaultSettingPage)

}
func checkAdminMenus() {
	log.Println("STARTS checkAdminMenus")
	methods := []string{
		"/home",
		"/pagesettings",
	}

	baseMapping := "/admin"

	for _, method := range methods {

		checkAdminMenu(baseMapping, method)

	}
	log.Println("ENDS checkAdminMenus")
}

func checkAdminMenu(baseMapping string, path string) {
	menuCode := reflections.GetWordsAfterLastChar(path, "/")
	menuCode = strings.Replace(menuCode, "/", "", -1)
	_, ok := getMenuByCode(menuCode)
	log.Println("getMenuByCode ", menuCode, " OK:", ok)
	if !ok {

		adminMenu := constructAdminMenu(baseMapping + path)
		repository.CreateNewWithoutValidation(&adminMenu)
	}

}

func getMenuByCode(code string) (menu entities.Menu, ok bool) {
	list := repository.FilterByKey(&entities.Menu{}, "Code", code)
	if len(list) != 1 {
		return menu, false
	}

	return list[0].(entities.Menu), true
}
func getPageOnlyByCode(code string) (page entities.Page, ok bool) {
	log.Println("getPageOnlyByCode: ", code)
	list := repository.FilterByKey(&[]entities.Page{}, "Code", code)
	if len(list) != 1 {
		log.Println("fails END getPageOnlyByCode: ", code)
		return page, false
	}
	log.Println("success END getPageOnlyByCode: ", code)
	return list[0].(entities.Page), true
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

func getPage(code string, defaultPageIfNotExist *entities.Page) *entities.Page {
	page, ok := getPageOnlyByCode(code)
	if ok {
		log.Printf("page with code: %v FOUND! \n", code)
		return &page
	}
	log.Printf("WILL SAVE page : %v...", code)
	peg := defaultPageIfNotExist
	repository.CreateNewWithoutValidation(peg)
	log.Println("defaultPageIfNotExist ID: ", peg.ID)
	return defaultPageIfNotExist
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
	_, ok := getMenuByCode(t.Name())
	if !ok {
		addNewManagementMenuPageFor(t)
	}

}

func addNewManagementMenuPageFor(t reflect.Type) {
	log.Println("Will add default menu for: ", t.Name())

	commonPage := true //= dto.commonManagementPage();
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
		menu.URL = ("/management/" + menuCode)
	} else {
		menu.URL = ("/management/" + menuCode)
	}

	repository.CreateNewWithoutValidation(&menu)

	log.Println("Success Adding Management Menu For: ", menuCode)
}

func getMenu(code string, defaultMenuIfNotExist *entities.Menu, menuPage *entities.Page) *entities.Menu {
	eixsitingPage := getPage(menuPage.Code, menuPage)
	existingMenu, ok := getMenuByCode(code) // menuRepository.findByCode(code);
	if ok {
		log.Printf("menu: %v FOUND!", code)
		return &existingMenu
	}

	log.Println("WILL SAVE menu with :", code)

	menu := defaultMenuIfNotExist
	menu.MenuPage = &entities.Page{}
	menu.PageID = uint16(eixsitingPage.ID)
	// menu.Validate()
	log.Println("00000000 menu.PageID:", menu.PageID)
	repository.SaveAndValidate(menu)
	return menu
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
