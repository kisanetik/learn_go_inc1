package loadjson

import (
	"fmt"
	"math/rand"
	"os"
)

type FileStruct struct {
	fileHandler *os.File
	cache       map[string]string
	iterator    int32
}

type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

const letters = 8

type Mem struct {
	memory map[string]URLData
}

var cache *Mem

func RandomString() string {
	literals := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	l := make([]rune, letters)
	for i := range l {
		l[i] = literals[rand.Intn(len(literals))]
	}

	return string(l)
}

func getConfig() *Mem {
	if nil == cache {
		cache = &Mem{}
	}

	return cache
}

func main() {
	fmt.Println(RandomString())
	fmt.Println(getConfig())
}
