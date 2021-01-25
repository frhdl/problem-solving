package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	var bigger, smaller, sum int64

	// get useful values
	bigger = math.MinInt64
	smaller = math.MaxInt64

	for _, item := range arr {
		temp := int64(item)

		if temp > bigger {
			bigger = temp
		}
		if temp < smaller {
			smaller = temp
		}

		sum += temp
	}

	minSum := sum - bigger
	maxSum := sum - smaller

	fmt.Printf("%v %v", minSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
