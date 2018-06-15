package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
//	"time"
)

// TODO: Add configurable path option.
const filename = "~/Desktop/todo.txt"

func main() {

	if (len(os.Args) == 1) {
		fmt.Printf("Usage: todo [list|add|done]\n")

	} else if (os.Args[1] == "list") {

		list()

	} else if (os.Args[1] == "done") {

		done(os.Args[2])

	} else if (os.Args[1] == "add") {

		add(os.Args[2])

	} else {
		fmt.Printf("%s", os.Args[1])
	}

}

func list() {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	count := 1

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		fileLine := fscanner.Text()

		data := strings.Split(fileLine, "\t")

		if (len(data) == 2 && data[0] == "0") {

			fmt.Printf("%d: %s\n", count, data[1])

			count = count + 1
		}


	}

}

func done(number string) {

	lines := []string{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	count := 1
	itemsUpdated := 0

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		fileLine := fscanner.Text()

		data := strings.Split(fileLine, "\t")

		if (data[0] == "0") {
			if (fmt.Sprint(count) == number) {
				data[0] = "1"
				itemsUpdated = itemsUpdated + 1
			}
			count = count + 1
		}
		lines = append(lines, strings.Join(data, "\t"))

	}

	if (itemsUpdated == 0) {
		fmt.Printf("Can not find item %s to mark as done.\n", number)
		return
	}

	file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(file)
	defer file.Close()


	for i := 0; i < len(lines); i++ {

		writer.WriteString(lines[i])
		writer.WriteString("\n")

	}

	writer.Flush()

	fmt.Printf("Marked %d item as done\n", itemsUpdated)

}

func add(item string) {

	filename := "~/Desktop/todo.txt"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(file)
	defer file.Close()

	fmt.Fprintln(writer, "0\t" + item)

	writer.Flush()

	fmt.Printf("Item added.\n")

}
