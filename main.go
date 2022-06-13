package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MAX = 100

func readString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	// 乱数で正解を生成
	rand.Seed(time.Now().UnixNano())
	target_number := rand.Intn(MAX-1) + 1

FOR:
	for {
		// 1行入力
		fmt.Printf("Please input number(1 - %d):", MAX)
		line := readString(sc)

		// 数値に変換
		number, err := strconv.Atoi(line)
		if err != nil {
			// 数値ではない
			fmt.Println("not number: " + err.Error())
			os.Exit(1)
		}

		switch {
		case number == target_number:
			// 正解なので抜ける
			fmt.Printf("Great! target is %d\n", target_number)
			break FOR

		case number < target_number:
			fmt.Println("too small.")

		case number > target_number:
			fmt.Println("too big.")

		}
	}
}
