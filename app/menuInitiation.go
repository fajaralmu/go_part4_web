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

func resetAllMenus() {
	menus, _ := repository.Filter(&[]entities.Menu{}, entities.Filter{})

	log.Println("Will reset allMenus")
	for _, menu := range menus {
		repository.Delete(menu.(entities.Menu), false)
	}
	pages, _ := repository.Filter(&[]entities.Page{}, entities.Filter{})
	log.Println("Will reset allPages")

	for _, page := range pages {
		repository.Delete(page.(entities.InterfaceEntity), false)
	}

	log.Println("////////END REMOVING/////////")
	initMenus()
	log.Println("//////////////END RESET//////////////")

}

///////////////////////////////
