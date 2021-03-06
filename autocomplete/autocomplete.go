package autocomplete

type AutoComplete struct {
	Dict *Dictionary
}

func Create() *AutoComplete {
	instance := &AutoComplete{
		Dict: createDictionary(),
	}

	return instance
}
