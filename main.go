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
const PROMPT_MESSAGE = "Please enter a number from 1 to %d\n"

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
		fmt.Printf(PROMPT_MESSAGE, MAX)
		line := readString(sc)

		// 数値に変換
		number, err := strconv.Atoi(line)
		if err != nil {
			// 数値ではない
			fmt.Printf("%s is not number.\n", line)
			continue
		}

		// 入力範囲内か確認する
		if number <= 0 || number > MAX {
			continue
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
