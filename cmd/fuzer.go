package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	ac "github.com/arjit95/fuzer/autocomplete"
	pq "github.com/arjit95/fuzer/queue"
)

func printResults(results []*pq.Item) {
	println("Matches: ")
	for _, match := range results {
		fmt.Printf("%s: %d\n", match.Value, match.Priority)
	}
}

var commands = [6][2]string{
	{"Add", "Adds an entry to the dictionary"},
	{"Remove", "Removes an entry from the dictionary"},
	{"Clear", "Clears the dictionary"},
	{"List", "Lists all the entries present in dictionary"},
	{"Search", "Performs a fuzzy search on the list of available words"},
	{"Exit", "Exits the program"},
}

func printInfo() {
	fmt.Printf("Please enter any one command to proceed\n\n")
	for _, command := range commands {
		fmt.Printf("%s -- %s\n", command[0], command[1])
	}
}

func processCmd(line string) (string, string) {
	parts := strings.Split(line, " ")
	if len(parts) == 1 {
		return parts[0], ""
	}

	return parts[0], strings.Join(parts[1:], " ")
}

func main() {
	printInfo()

	instance := ac.Create()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(">")

	for scanner.Scan() {
		command, arg := processCmd(scanner.Text())

		switch command {
		case "Add":
			instance.Dict.Add(arg)
			break
		case "Remove":
			instance.Dict.Remove(arg)
			break
		case "List":
			fmt.Printf("%v \n", instance.Dict.List())
			break
		case "Clear":
			instance.Dict.Clear()
			break
		case "Search":
			start := time.Now()
			printResults(instance.GetMatches(arg, 5))
			fmt.Printf("Search took %s\n", time.Since(start))
			break
		case "Exit":
			os.Exit(0)
			break
		default:
			fmt.Printf("Unknown command %s \n", command)
			fmt.Printf("Args: %s \n", arg)
			printInfo()
		}

		fmt.Printf(">")
	}
}
