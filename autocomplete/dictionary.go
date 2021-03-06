package autocomplete

import (
	"runtime"
	"strings"
	"sync"
)

type Dictionary struct {
	words []string
	lower []string
}

func indexOf(words []string, element string) int {
	for i, word := range words {
		if word == element {
			return i
		}
	}

	return -1
}

func remove(words []string, index int) []string {
	words[index] = words[len(words)-1]
	return words[:len(words)-1]
}

func (dict *Dictionary) Count() int {
	return len(dict.words)
}

func (dict *Dictionary) Add(word string) {
	dict.words = append(dict.words, word)
	dict.lower = append(dict.lower, strings.ToLower(word))
}

func (dict *Dictionary) AddAll(words []string) {
	for _, word := range words {
		dict.Add((word))
	}
}

func (dict *Dictionary) Clear() {
	dict.words = make([]string, 0)
	dict.lower = make([]string, 0)
}

func (dict *Dictionary) iterateOverWords(from int, to int, callback func(string, string, int) bool) {
	for ; from < to; from++ {
		proceed := callback(dict.words[from], dict.lower[from], from)
		if !proceed {
			return
		}
	}
}

func (dict *Dictionary) ForSomeWords(callback func(string, string, int) bool) {
	maxThreads := runtime.NumCPU()
	count := dict.Count()

	extra := count % maxThreads
	chunk := count - (count % maxThreads)
	chunk = count / maxThreads
	length := dict.Count()

	start := 0
	var wg sync.WaitGroup

	for start < length {
		wg.Add(1)
		go func(from, to int) {
			dict.iterateOverWords(from, to, callback)
			wg.Done()
		}(start, start+chunk+extra)

		start = start + chunk + extra
		extra = 0
	}

	wg.Wait()
}

func (dict *Dictionary) List() []string {
	return dict.words
}

func (dict *Dictionary) Remove(word string) {
	idx := indexOf(dict.words, word)
	if idx < 0 {
		return
	}

	dict.words = remove(dict.words, idx)
	dict.lower = remove(dict.lower, idx)
}

func createDictionary() *Dictionary {
	dict := &Dictionary{
		words: make([]string, 0),
		lower: make([]string, 0),
	}

	return dict
}
