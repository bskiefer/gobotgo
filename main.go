package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"arc/arc"
)

func main() {
	content, err := ioutil.ReadFile("list.txt")
	if err != nil {
		fmt.Print(err)
	}
	lines := strings.Split(string(content), "\n")

	cookieData, err := ioutil.ReadFile("cookie.txt")
	if err != nil {
		fmt.Print(err)
	}
	cookie := strings.Split(string(cookieData), "\n")[0]

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i, s := range lines {
		// fmt.Println(i, s)
		archiveURL, err := arc.Capture(s, cookie)
		if err != nil {
			fmt.Print(err)
		}

		line := []string{lines[i], archiveURL}
		fmt.Println(line)
		writer.Write(line)
		checkError("Cannot write to file", err)

	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Print(message, err)
	}
}
