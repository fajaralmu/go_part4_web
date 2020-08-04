package app

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fajaralmu/go_part4_web/reflections"
)

type customStaticHandler struct {
	root http.FileSystem
}

func (c *customStaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// log.Println("ServeHTTP ", r.RequestURI)
	justFilePath := strings.Replace(r.RequestURI, "/static/", "", 1)

	justFilePath = reflections.RemoveCharsAfter(justFilePath, "?")
	// log.Println("GOTO FILE PATH: ", justFilePath)

	file, err := c.root.Open(justFilePath)

	if err != nil {
		writeErrorMsgBadRequest(w, "Invalid path")
		log.Println("[ERROR] Open()", err.Error())
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		writeErrorMsgBadRequest(w, "Invalid path")
		log.Println("[ERROR] Stat()", err.Error())
		return
	}
	isDirectory := fileInfo.IsDir()
	if isDirectory {
		writeErrorMsgBadRequest(w, "Invalid path BROO")
		return
	}

	////file info////
	// log.Println("File Name: ", fileInfo.Name(), "SIZE: ", fileInfo.Size())

	///reads file///
	b, err := ioutil.ReadAll(file)

	if err != nil {
		writeErrorMsgBadRequest(w, err.Error())
		return
	}

	extension := reflections.GetFileExtention(fileInfo.Name())
	if strings.ToLower(extension) == "js" {
		extension = "text/javascript"
	} else if strings.ToLower(extension) == "css" {
		extension = "text/css"
	}
	w.Header().Add("Content-Type", extension)
	w.Header().Add("inf0", "123-static")
	w.Write((b))

}
