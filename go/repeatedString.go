package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	sLength := int64(len(s))
	aCountByArgGroup := aCountInArg(s)
	groupCount := int64(n / sLength)
	// groupCountで割り切れない場合に剰余で、残りのループカウント数を計算
	remainingCount := n % sLength
	sumACount := aCountByArgGroup * groupCount

	for i := int64(0); i < remainingCount; i++ {
		if s[i] == 'a' {
			sumACount++
		}
	}

	return sumACount
}

func aCountInArg(s string) int64 {
	var count int64
	for i := range s {
		if s[i] == 'a' {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

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
