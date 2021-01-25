package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	lengh := float64(len(arr))
	var positive float64
	var negative float64
	var zero float64

	for _, item := range arr {
		if item > 0 {
			positive++
		} else if item < 0 {
			negative++
		} else {
			zero++
		}
	}
	// Positive Ratios
	fmt.Printf("%.6f\n", (positive / lengh))
	// Negative Ratios
	fmt.Printf("%.6f\n", (negative / lengh))
	// Zero Ratios
	fmt.Printf("%.6f\n", (zero / lengh))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
