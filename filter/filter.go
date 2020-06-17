package filter

import (
	"bytes"
	"strings"
)

// WordFilter 敏感词过滤器
type WordFilter struct {
	root        *Node
	StripSpace  bool
	PlaceHolder rune
}

// NewWordFilter return new wordfilter instance
func NewWordFilter() *WordFilter {
	return &WordFilter{
		PlaceHolder: defaultPlaceHolder,
		StripSpace:  defaultStripSpace,
		root:        NewNode(),
	}
}

// BatchAdd convert sensitive word list into sensitive word tree
func (w *WordFilter) BatchAdd(words []string) {
	for _, word := range words {
		if w.StripSpace {
			word = stripSpace(word)
		}
		w.root.add([]rune(word))
	}
}

// Add add new sensitive word into sensitive word tree
func (w *WordFilter) Add(word string) {
	if w.StripSpace {
		word = stripSpace(word)
	}
	w.root.add([]rune(word))
}

// Remove remove sensitive word from sensitive word tree
func (w *WordFilter) Remove(word string) {
	if w.StripSpace {
		word = stripSpace(word)
	}
	w.root.remove([]rune(word))
}

// Filter replace sensitive word with placeholder
func (w *WordFilter) Filter(text string) string {
	if w.StripSpace {
		text = stripSpace(text)
	}

	textRune := []rune(text)
	length := len(textRune)
	if length == 0 {
		return ""
	}
	for i := 0; i < length; i++ {
		mLength := w.root.mlength(textRune[i:], 0)
		if mLength > 0 {
			for j := 0; j < mLength; j++ {
				textRune[i+j] = w.PlaceHolder
			}
			i = i + mLength - 1
		}
	}
	return string(textRune)
}

// Contains return is exist sensitive word in text
func (w *WordFilter) Contains(text string) bool {
	if w.StripSpace {
		text = stripSpace(text)
	}
	return w.root.contains([]rune(text))
}

// Strip space
func stripSpace(str string) string {
	fields := strings.Fields(str)
	var bf bytes.Buffer
	for _, field := range fields {
		bf.WriteString(field)
	}
	return bf.String()
}
