package linker

import (
	"os"
	"path/filepath"
)

func compressUrl(url string) string {
	tFile, err := os.CreateTemp("", "")
	os.WriteFile(tFile.Name(), url, 0644)
	return string("http://localhost:8080/" + filepath.Base(tFile.Name()))
}
