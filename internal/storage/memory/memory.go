package memory

import (
	"errors"
	"fmt"
	"sync"
)

type Memory struct {
	Mapping sync.Map
}

func NewMemory() *Memory {
	return &Memory{
		Mapping: sync.Map{},
	}
}

func (m *Memory) GetUrl(aliasUrl string) (string, error){
	val, ok := m.Mapping.Load(aliasUrl)
	if !ok{
		return "", errors.New("this alias does not exist")
	}

	return val.(string), nil
}

func (m *Memory) SaveUrl(urlToSave, aliasUrl string) error{
	m.Mapping.Store(aliasUrl, urlToSave)

	m.Mapping.Range(func(key, value interface{}) bool {
        fmt.Println(key, value)
        return true
    })

	return nil
}


func (m *Memory) IsDublicate(alias string) bool{
	_, ok := m.Mapping.Load(alias)

	return ok
}