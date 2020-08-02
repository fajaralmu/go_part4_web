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
	checkDefaultMenu()
	checkManagementPage()
	checkAdminMenus()

	getPersistenctEntities()
	log.Println("////////////////////END MENU INITIALIZE////////////////////")
}

func defaultAboutPage() entities.Page {
	return getPage("about", config_defaultAboutPage)
}
func defaultSettingPage() entities.Page {
	return getPage("setting", config_defaultSettingPage)

}
func checkAdminMenus() {
	log.Println("STARTS checkAdminMenus")
	methods := []string{
		"/admin/home",
		"/admin/pagesettings",
	}

	baseMapping := "/admin"

	for _, method := range methods {

		checkAdminMenu(baseMapping, method)

	}
	log.Println("ENDS checkAdminMenus")
}

func checkAdminMenu(baseMapping string, path string) {
	menuCode := reflections.GetWordsAfterLastChar(path, "/")
	_, ok := getMenuByCode(menuCode)
	if !ok {

		adminMenu := constructAdminMenu(baseMapping + path)
		repository.SaveWihoutValidation(&adminMenu)
	}

}

func getMenuByCode(code string) (menu entities.Menu, ok bool) {
	list := repository.FilterByKey(&entities.Menu{}, "Code", code)
	if len(list) != 1 {
		return menu, false
	}

	return list[0].(entities.Menu), true
}
func getPageOnlyByCode(code string) (menu entities.Page, ok bool) {
	list := repository.FilterByKey(&entities.Page{}, "Code", code)
	if len(list) != 1 {
		return menu, false
	}

	return list[0].(entities.Page), true
}

func constructAdminMenu(path string) entities.Menu {

	log.Println("constructAdminMenu LINK: ", path)
	menuCode := reflections.GetWordsAfterLastChar(path, "/")
	defAdminPage := defaultAdminPage()
	adminMenu := entities.Menu{
		Code:            menuCode,
		Color:           "#000000",
		BackgroundColor: "#ffffff",
		Description:     "Generated [" + menuCode + "]",
		Name:            strings.ToUpper(menuCode) + "(auto)",
		URL:             path,
		MenuPage:        &defAdminPage,
	}

	return adminMenu
}

func defaultAdminPage() entities.Page {
	return getPage("admin", adminPage())
}

func managementPage() entities.Page {
	return config_defaultManagementPage
}

func adminPage() entities.Page {
	return config_defaultAdminPage
}

func getPage(code string, defaultPageIfNotExist entities.Page) entities.Page {
	page, ok := getPageOnlyByCode(code)
	if ok {
		log.Printf("page with code: %v FOUND! \n", code)
		return page
	}
	log.Printf("WILL SAVE page : %v...", code)
	repository.CreateNewWithoutValidation(&defaultPageIfNotExist)
	return defaultPageIfNotExist
}

func checkManagementPage() {
	log.Println("STARTS _managementPage")
	_managementPage, ok := getPageOnlyByCode("management")
	if ok {
		log.Println("managementPage FOUND")
		return
	}

	log.Println("managementPage NOT FOUND. WILL ADD SETTING")
	_managementPage = managementPage()
	repository.CreateNewWithoutValidation(&_managementPage)
	log.Println("STARTS _managementPage")
}

func defaultManagementPage() entities.Page {

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
	menuCode := t.Name()
	managementPage, _ := getPageOnlyByCode("management")
	menu := entities.Menu{
		Code:            menuCode,
		Name:            reflections.ExtractCamelCase(t.Name()),
		MenuPage:        &managementPage,
		Color:           "#000000",
		BackgroundColor: "#ffffff",
		Description:     "Generated Management Page for: " + t.Name(),
	}

	if commonPage {
		menu.URL = ("/management/" + menuCode)
	} else {
		menu.URL = ("/management/" + menuCode)
	}

	repository.CreateNewWithoutValidation(&menu)

	log.Println("Success Adding Management Menu For: ", menuCode)
}

func checkDefaultMenu() {
	getMenu("management", config_defaultManagementMenu, defaultAdminPage())
}

func getMenu(code string, defaultMenuIfNotExist entities.Menu, menuPage entities.Page) entities.Menu {
	eixsitingPage := getPage(menuPage.Code, menuPage)
	menu, ok := getMenuByCode(code) // menuRepository.findByCode(code);
	if ok {
		log.Printf("menu: %v FOUND!", code)
		return menu
	}

	log.Println("WILL SAVE menu with :", code)

	menu = defaultMenuIfNotExist
	menu.MenuPage = &(eixsitingPage)

	repository.SaveWihoutValidation(&menu)
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
