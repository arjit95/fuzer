package main

import (
	"syscall/js"

	ac "github.com/arjit95/fuzer/autocomplete"
)

var autoComplete *ac.AutoComplete

func main() {
	c := make(chan struct{}, 0)
	autoComplete = ac.Create()

	js.Global().Set("_Fuzer", map[string]interface{}{
		"addAll":     js.FuncOf(addWord),
		"removeWord": js.FuncOf(removeWord),
		"search":     js.FuncOf(search),
		"add":        js.FuncOf(addAll),
		"list":       js.FuncOf(list),
		"clear":      js.FuncOf(clear),
		"count":      js.FuncOf(count),
	})

	<-c
}

func count(this js.Value, args []js.Value) interface{} {
	return autoComplete.Dict.Count()
}

func clear(this js.Value, args []js.Value) interface{} {
	autoComplete.Dict.Clear()
	return nil
}

func list(this js.Value, args []js.Value) interface{} {
	words := autoComplete.Dict.List()
	result := make([]interface{}, 0)

	for _, word := range words {
		result = append(result, word)
	}

	return result
}

func addAll(this js.Value, args []js.Value) interface{} {

	for _, word := range args {
		autoComplete.Dict.Add(word.String())
	}

	return js.ValueOf(autoComplete.Dict.Count())
}

func addWord(this js.Value, args []js.Value) interface{} {
	autoComplete.Dict.Add(args[0].String())
	return js.ValueOf(autoComplete.Dict.Count())
}

func removeWord(this js.Value, args []js.Value) interface{} {
	autoComplete.Dict.Remove(args[0].String())
	return js.ValueOf(autoComplete.Dict.Count())
}

func convertToJSArray(arr []int) []interface{} {
	result := make([]interface{}, len(arr))

	for i, val := range arr {
		result[i] = val
	}

	return result
}

func search(this js.Value, args []js.Value) interface{} {
	handler := js.FuncOf(func(this js.Value, promiseArgs []js.Value) interface{} {
		resolve := promiseArgs[0]
		reject := promiseArgs[1]
		if len(args) != 2 {
			reject.Invoke("Insufficient args")
			return nil
		}

		pattern := args[0].String()
		count := args[1].Int()

		go func() {
			result := autoComplete.GetMatches(pattern, count)
			response := make([]interface{}, len(result))
			for i, r := range result {
				response[i] = map[string]interface{}{
					"rank":    r.Priority,
					"word":    r.Value,
					"matches": convertToJSArray(r.Matches),
				}
			}

			resolve.Invoke(response)
		}()

		return nil
	})

	promiseConstructor := js.Global().Get("Promise")
	return promiseConstructor.New(handler)
}
