package imagemanagment

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

const IMAGE_PATH = "/home/xavier/DATA_REPTILE/images"

func setup() {

}

func getReptileFolder(id int) string {
	return filepath.Join(IMAGE_PATH, fmt.Sprint(id))
}

func SaveImage(id int, image []byte) {
	reptile_path := getReptileFolder(id)
	err := os.Mkdir(reptile_path, 0760)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	filepath.Walk(reptile_path, walkFuncReptile)
}

func walkFuncReptile(dir string, info fs.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	log.Println(info.Name())

	return nil
}
