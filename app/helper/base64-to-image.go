package helper

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

//Base64toJpg Given a base64 string of a JPEG, encodes it into an JPEG image test.jpg
func Base64toJpg(data string, id string) string {

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		return "can not decode image!!!"
	}
	bounds := m.Bounds()
	log.Println("base64toJpg", bounds, formatString)
	//Encode from image format to writer
	// imgFilename := "../uploads/" + id + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
	imgFilename := "uploads/" + id + ".jpg"
	os.Mkdir("uploads", os.ModePerm)
	f, err := os.OpenFile(imgFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return "can not create image file"
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 100})
	if err != nil {
		return "can not create image file"
	}
	log.Println("Jpg file", imgFilename, "created")
	return ""

}

//GetJPEGbase64 Gets base64 string of an existing JPEG file
func GetJPEGbase64(fileName string) string {

	imgFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	//fmt.Println("Base64 string is:", imgBase64Str)
	return imgBase64Str

}