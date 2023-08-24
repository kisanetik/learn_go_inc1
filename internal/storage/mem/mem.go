package mem

import (
	"github.com/kisanetik/learn_go_inc1/internal/utils"
)

type Mem struct {
	memory map[string]string
}

func NewMem() (*Mem, error) {
	m := &Mem{
		memory: make(map[string]string),
	}

	return m, nil
}

func (m *Mem) Save(long string) (string, error) {
	short := utils.RandomString()

	m.memory[short] = long

	return short, nil
}

func (m *Mem) Get(short string) string {
	return m.memory[short]
}

func (m *Mem) Close() error {
	return nil
}
