package fs

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kisanetik/learn_go_inc1/internal/logger"
	"github.com/kisanetik/learn_go_inc1/internal/utils"
)

type Fs struct {
	fh    *os.File
	cache map[string]string
	count int64
}

type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewFsFromFile(path string) (*Fs, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	return NewFs(file)
}

func NewFs(file *os.File) (*Fs, error) {
	fs := &Fs{
		fh:    file,
		cache: make(map[string]string),
		count: 0,
	}

	urlData := &URLData{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		err := json.NewDecoder(bytes.NewReader([]byte(line))).Decode(&urlData)
		if err != nil {
			logger.Errorf("Error json decode in NewFs: %s", err)
		}

		fs.cache[urlData.ShortURL] = urlData.OriginalURL
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return fs, nil
}

func (m *Fs) Save(long string) (string, error) {
	urlData := &URLData{
		UUID:        fmt.Sprintf("%d", m.count),
		ShortURL:    utils.RandomString(),
		OriginalURL: long,
	}

	jsonData, err := json.Marshal(urlData)
	if err != nil {
		return "", fmt.Errorf("can't marshal json: %w", err)
	}

	_, err = m.fh.Write([]byte("\n"))
	if err != nil {
		return "", fmt.Errorf("can't write to file: %w", err)
	}

	_, err = m.fh.Write(jsonData)
	if err != nil {
		return "", fmt.Errorf("can't write to file: %w", err)
	}

	m.count++
	m.cache[urlData.ShortURL] = urlData.OriginalURL

	return urlData.ShortURL, nil
}

func (m *Fs) Get(short string) string {
	return m.cache[short]
}

func (m *Fs) Close() error {
	return m.fh.Close()
}
