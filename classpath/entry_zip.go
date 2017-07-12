package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

// import "fmt"
// import "strings"

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	// fmt.Println(className)
	for _, f := range r.File {
		// debug
		// if strings.Contains(f.Name, "java/lang") {
		// 	fmt.Println("zipEntry    " + f.Name)
		// }
		if f.Name == className {
			// fmt.Println("zipEntry    " + f.Name)
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}

	return nil, nil, errors.New("Class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
