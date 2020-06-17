package filter

import (
	"io/ioutil"
	"strings"

	"gitee.com/yutou_core/jarvis/utils/derr"
)

// Inst word filter instance
var Inst *WordFilter

// Init init Inst with word read from file
func Init() (err error) {
	Inst = NewWordFilter()
	words, err := LoadAllWords()
	if err != nil {
		return err
	}
	Inst.BatchAdd(words)
	return nil
}

// LoadAllWords load file from project sensitive word dir
func LoadAllWords() (words []string, err error) {
	words = []string{}
	files, err := ioutil.ReadDir(defaultWordsDirPath)
	if err != nil {
		return nil, derr.Errorf("ReadDir failed. Path:%v", defaultWordsDirPath).WithCause(err)
	}
	for _, file := range files {
		path := defaultWordsDirPath + "/" + file.Name()
		text, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, derr.Errorf("ReadFile failed. Path:%v", path).WithCause(err)
		}
		lines := strings.Split(string(text), "\n")
		words = append(words, lines...)
	}
	return words, nil
}
