package imagemanagment

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func getFolderId(id int) string {
	return filepath.Join(IMAGE_PATH, fmt.Sprint(id))
}

func getPhotoNumber(filename string) string {
	// We have 1_01.png, after fut _ = 1, after = 01.png
	_, after, cuttable := strings.Cut(filename, "_")
	if !cuttable {
		log.Panicln("The filename: ", filename, "cannot be cutted in imagemanagment")
	}
	//now we have to get rid of the suffix, as we will work with a variaety of image formats, we have to cut again
	// number = 01, _ = png
	number, _, cuttable := strings.Cut(after, ".")
	if !cuttable {
		log.Panicln("The filename: ", filename, "cannot be cutted in imagemanagment")
	}
	return number
}

func newReptileImageName(list []string, folderNameString string) string {
	var maxNumberOfPhotos int

	for i := 0; i < len(list); i++ {

		if strings.HasPrefix(list[i], folderNameString) {
			str_number := getPhotoNumber(list[i])
			number, err := strconv.Atoi(str_number)

			if err != nil {
				log.Println("OOPS, the filename: ", list[i], " is not a valid file name")
				panic(err)
			}
			if maxNumberOfPhotos < number {
				maxNumberOfPhotos = number
			}
		}
	}
	return folderNameString + "_" + fmt.Sprint(maxNumberOfPhotos+1)
}
