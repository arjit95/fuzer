package autocomplete

import (
	"container/heap"
	"fmt"
	"strings"
	"sync"

	pq "github.com/arjit95/fuzer/queue"
)

func highlightWord(word string, indexes []int, openTag string, closeTag string) string {
	if len(openTag) == 0 && len(closeTag) == 0 {
		return word
	}

	result := ""
	highlightIdx := 0

	for i, ch := range word {
		if highlightIdx < len(indexes) && i == indexes[highlightIdx] {
			result += fmt.Sprintf("%s%c%s", openTag, ch, closeTag)
		} else {
			result += string(ch)
		}
	}

	return result
}

func (ac *AutoComplete) GetMatches(pattern string, count int) []*pq.Item {
	queue := &pq.PriorityQueue{}
	heap.Init(queue)

	originalPattern := []rune(pattern)
	lowerPattern := []rune(strings.ToLower(pattern))

	var mutex = &sync.Mutex{}

	ac.Dict.ForSomeWords(func(word, lower string, idx int) bool {
		wordRank, highlights := rank(originalPattern, lowerPattern, []rune(word), []rune(lower))
		if wordRank != MIN_RANK {
			mutex.Lock()

			queue.Push(&pq.Item{
				Value:    word,
				Priority: wordRank,
				Matches:  highlights,
			})

			if queue.Len() > count {
				heap.Pop(queue)
			}

			mutex.Unlock()
		}

		return true
	})

	result := make([]*pq.Item, queue.Len())

	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(queue).(*pq.Item)
	}

	return result
}
