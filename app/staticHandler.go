package app

import (
	"io/ioutil"
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
		writeErrorMsgBadRequest(w, err.Error())
		return
	} else {
		fileInfo, err := file.Stat()
		if err != nil {
			writeErrorMsgBadRequest(w, err.Error())
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

		w.Header().Add("Content-Type", reflections.GetFileExtention(fileInfo.Name()))
		w.Header().Add("inf0", "123-static")
		w.Write((b))

	}
	// h := http.FileServer(c.root)
	// h.ServeHTTP(w, r)
	// log.Println("H: ", reflect.TypeOf(h))

}
