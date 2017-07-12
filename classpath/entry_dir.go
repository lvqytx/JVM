package classpath

import "io/ioutil"
import "path/filepath"

// import "fmt"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	// fmt.Println("entry_dir     " + fileName)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}
func (self *DirEntry) String() string {
	return self.absDir
}
