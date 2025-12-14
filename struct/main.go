package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"practice.com/note/note"
	"practice.com/note/todo"
)

type saver interface {
	Save() error
}

/*
	type displayer interface {
		Display()
	}
*/
type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(1)
	printSomething("Hello world")
	printSomething(1.2)

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	outputData(userNote)

}
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving data failed.")
		return err
	}
	fmt.Println("saving the data succeded.")
	return nil
}
func getNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// "any value allowed" type
func printSomething(value interface{}) {
	/*switch value.(type) {
	case int:
		fmt.Println("integer value: ", value)
	case string:
		fmt.Println("string value: ", value)
	case float64:
		fmt.Println("float value: ", value)
	}*/
	intVal, ok := value.(int)
	if ok {
		fmt.Println("integer: ", intVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("string: ", stringVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("float: ", floatVal)
		return
	}
}
