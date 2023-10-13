package imagemanagment

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const IMAGE_PATH = "/home/xavier/DATA_REPTILE/images"

func SaveImage(id int, image []byte) {
	reptile_path := getFolderId(id)
	err := os.Mkdir(reptile_path, 0760)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	var imageNamesList []string
	filepath.Walk(reptile_path, func(dir string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		log.Println(info.Name())
		imageNamesList = append(imageNamesList, info.Name())
		return nil
	})
}
