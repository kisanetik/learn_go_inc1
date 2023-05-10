package linker

import (
	"os"
	"path/filepath"
)

func CompressURL(url string) string {
	tFile, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	os.WriteFile(tFile.Name(), []byte(url), 0644)
	return string("http://localhost:8080/" + filepath.Base(tFile.Name()))
}
