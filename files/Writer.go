package files

import (
	"encoding/base64"
	"log"
	"os"
	"strings"

	"github.com/fajaralmu/go_part4_web/reflections"
)

const imgPath string = "./public/img/app/"

//WriteBase64Img writes file from base64 string
func WriteBase64Img(imgRawData string, code string) string {
	log.Println("write Base64Data Img")

	imageData := strings.Split(imgRawData, ",")
	if imageData == nil || len(imageData) < 2 {
		log.Println("imgRawData is empty")
		return ""
	}
	// create a buffered image
	imageString := imageData[1]

	dec, err := base64.StdEncoding.DecodeString(imageString)
	if err != nil {
		log.Println("Error DecodeString: ", err.Error())
		return ""
	}

	// write the image to a file
	imageIdentity := imageData[0]
	imageType := strings.Replace(imageIdentity, "data:image/", "", 1)
	imageType = strings.Replace(imageType, ";base64", "", 1)
	imageName := reflections.RandomNum(10)
	path := imgPath // webAppConfiguration.getUploadedImageRealPath();

	imageFileName := code + "_" + imageName + "." + imageType

	log.Println("Writing to path: ", path+imageFileName)

	f, err := os.Create(path + imageFileName)
	if err != nil {
		log.Println("Error Create File: ", err.Error())
		return ""
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		log.Println("Error Write File: ", err.Error())
		return ""
	}
	if err := f.Sync(); err != nil {
		log.Println("Error Sync File: ", err.Error())
		return ""
	}

	return imageFileName
}
