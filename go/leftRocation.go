package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	arrayNum := len(a)
	// 要素を左にずらす回数
	rotationCount := d % int32(arrayNum)

	if rotationCount == 0 {
		return a
	}

	arr1 := []int32{}
	arr2 := []int32{}

	for i := range a {
		if i < int(rotationCount) {
			arr1 = append(arr1, a[i])
		} else {
			arr2 = append(arr2, a[i])
		}
	}

	return append(arr2, arr1...)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := rotLeft(a, d)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
