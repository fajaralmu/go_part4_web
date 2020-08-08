package app

import (
	"log"
	"strings"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/reflections"
	"github.com/fajaralmu/go_part4_web/repository"
)

func defaultAboutPage() *entities.Page {
	return getPageFromDB("about", configDefaultAboutPage)
}
func defaultSettingPage() *entities.Page {
	return getPageFromDB("setting", configDefaultSettingPage)

}

func defaultManagementPage() *entities.Page {
	return getPageFromDB("management", managementPage())
}

func defaultAdminPage() *entities.Page {
	return getPageFromDB("admin", adminPage())
}

func managementPage() *entities.Page {
	return configDefaultManagementPage
}

func adminPage() *entities.Page {
	return configDefaultAdminPage
}

func getMenuByCode(code string) (menu entities.Menu, ok bool) {
	log.Println("getMenuByCode: ", code)

	list := repository.FilterByKey(&[]entities.Menu{}, "Code", code)
	if len(list) != 1 {
		log.Println("Menu (", code, ") Not Found")
		return menu, false
	}
	log.Println("Menu (", code, ") Found")
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

func getPageFromDB(code string, defaultPageIfNotExist *entities.Page) *entities.Page {
	page, ok := getPageOnlyByCode(code)
	if ok {
		log.Printf("page with code: %v FOUND! \n", code)
		return &page
	}
	log.Printf("WILL SAVE page : %v...", code)

	defaultPageObject := defaultPageIfNotExist
	repository.CreateNewWithoutValidation(defaultPageObject)

	log.Println("defaultPageIfNotExist ID: ", defaultPageObject.ID)
	return defaultPageIfNotExist
}

func getMenuFromDB(code string, defaultMenuIfNotExist *entities.Menu, menuPage *entities.Page) *entities.Menu {
	eixsitingPage := getPageFromDB(menuPage.Code, menuPage)
	existingMenu, ok := getMenuByCode(code) // menuRepository.findByCode(code);
	if ok {

		return &existingMenu
	}

	log.Println("WILL SAVE menu with :", code)

	menu := defaultMenuIfNotExist
	menu.MenuPage = &entities.Page{}
	menu.PageID = uint16(eixsitingPage.ID)
	log.Println("menu.PageID:", menu.PageID)

	repository.SaveAndValidate(menu)
	return menu
}

func checkAdminMenu(baseMapping string, path string) {
	menuCode := reflections.GetWordsAfterLastChar(path, "/")
	menuCode = strings.Replace(menuCode, "/", "", -1)
	_, ok := getMenuByCode(menuCode)

	if !ok {
		adminMenu := constructAdminMenu(baseMapping + path)
		repository.CreateNewWithoutValidation(&adminMenu)
	}

}
