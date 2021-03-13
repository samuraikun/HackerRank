package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	maxSum := int32(0)

	// 横方向のループ回数
	for i := 0; i < 4; i++ {
		// 縦方向のループ回数
		for j := 1; j < 5; j++ {
			/*
				例. 6x6のサンプル2重配列の場合
				1 1 1 0 0 0
				0 1 0 0 0 0
				1 1 1 0 0 0
				0 0 0 0 0 0
				0 0 0 0 0 0
				0 0 0 0 0 0

				各行各列スタートの砂時計の組み合わせは、
				| 1行1列 | 1行2列 |
				|-------|--------|
				| 1 1 1 | 1 1 0  |
				|   1   |   0    |
				| 1 1 1 | 1 1 0  |
				パターン: 砂時計の配列の組み合わせ
				{
					'砂時計上部分': [n行n-1列, n行n列, n行n+1列],
					'砂時計真ん中部分': n+1行n列,
					'砂時計下部分': [n+2行n-1列, n+2行n列, n+2行n+1列]
				}
			*/
			tempSum := arr[i][j-1] + arr[i][j] + arr[i][j+1] + arr[i+1][j] + arr[i+2][j-1] + arr[i+2][j] + arr[i+2][j+1]

			if i == 0 && j == 1 {
				maxSum = tempSum
			} else {
				maxSum = max(maxSum, tempSum)
			}
		}
	}

	return maxSum
}

func max(a int32, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

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
