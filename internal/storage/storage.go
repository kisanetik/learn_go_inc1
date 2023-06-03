package storage

import (
	"encoding/json"
	"os"

	"github.com/kisanetik/learn_go_inc1/config"
)

type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type Mem map[string]URLData

var cache Mem

func GetData() Mem {
	if nil == cache {
		cache = *&Mem{}
		serialized, err := os.ReadFile(config.GetConf().FileStoragePath)
		if !os.IsNotExist(err) && len(serialized) > 2 {
			if err := json.Unmarshal(serialized, &cache); err != nil {
				panic(err)
			}

		}
	}

	return cache
}

func AddToData(record URLData) {
	cache := GetData()
	cache[record.UUID] = record
}

func Save() bool {
	strJson, _ := json.Marshal(cache)
	err := os.WriteFile(config.GetConf().FileStoragePath, strJson, 0666)
	if err != nil {
		return false
	}

	return true
}
