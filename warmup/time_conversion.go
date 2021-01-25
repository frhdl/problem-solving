package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	var hours []string

	//PM case
	if strings.Contains(s, "PM") {
		hours = strings.Split(s, ":")
		if hours[0] != "12" {
			temp, _ := strconv.Atoi(hours[0])
			temp += 12
			hours[0] = strconv.Itoa(temp)
		}
		// Remove letters
		hours[2] = strings.Replace(hours[2], "PM", "", 2)
	} else {
		// AM case
		hours = strings.Split(s, ":")

		if hours[0] == "12" {
			hours[0] = "00"
		}
		// Remove letters
		hours[2] = strings.Replace(hours[2], "AM", "", 2)
	}

	return fmt.Sprintf("%v:%v:%v", hours[0], hours[1], hours[2])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
