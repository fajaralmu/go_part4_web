package files

import (
	"encoding/base64"
	"log"
	"os"
	"strings"
)

//WriteBase64Img writes file from base64 Data
func WriteBase64Img(imgRawData string) string {
	log.Println("writeBase64Img")

	imageData := strings.Split(imgRawData, ",")
	if imageData == nil || len(imageData) < 2 {
		return ""
	}
	// create a buffered image
	imageString := imageData[1]

	dec, err := base64.StdEncoding.DecodeString(imageString)
	if err != nil {
		panic(err)
	}

	// BufferedImage image = null;
	// byte[] imageByte;

	// Base64.Decoder decoder = Base64.getDecoder();
	// imageByte = decoder.decode(imageString);
	// ByteArrayInputStream bis = new ByteArrayInputStream(imageByte);
	// image = ImageIO.read(bis);
	// bis.close();

	// write the image to a file
	imageIdentity := imageData[0]
	imageType := strings.Replace(imageIdentity, "data:image/", "", 1)
	imageType = strings.Replace(imageType, ";base64", "", 1)
	//   imageIdentity.replace("data:image/", "").replace(";base64", "");
	imageName := "RAND_IMG"
	//String path = servletContext.getRealPath("/resources/img/upload");
	//String path  ="D:/Development/Files/Web/Shop1/Images";
	path := "./public/img/app/" // webAppConfiguration.getUploadedImageRealPath();

	imageFileName := "APP_" + imageName + "." + imageType
	// File outputfile = new File(path +"/"+imageFileName);
	// log.Println("==========UPLOADED FILE: ", outputfile.getAbsolutePath());
	// ImageIO.write(image, imageType, outputfile);
	// log.Println("==output file: ", outputfile.getAbsolutePath());

	log.Println("Writing to path: ", path+imageFileName)

	f, err := os.Create(path + imageFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	return imageFileName
}
