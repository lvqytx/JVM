package classpath

import "io/ioutil"
import "path/filepath"
import "fmt"

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
	fmt.Println("dir")
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}
func (self *DirEntry) String() string {
	return self.absDir
}
