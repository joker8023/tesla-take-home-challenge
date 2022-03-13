package pkg

import "os"

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateFile(file string) {
	ok := PathExists(file)
	if !ok {
		f, err := os.Create(file)
		if err != nil {
			return
		}
		defer f.Close()
	}

}
