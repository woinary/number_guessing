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
	m := Messages{}

	// 乱数で正解を生成
	rand.Seed(time.Now().UnixNano())
	target_number := rand.Intn(MAX-1) + 1

FOR:
	for {
		// 1行入力
		fmt.Printf(m.Get(PROMPT_MESSAGE), MAX)
		line := readString(sc)

		// 数値に変換
		number, err := strconv.Atoi(line)
		if err != nil {
			// 数値ではない
			fmt.Printf(m.Get(NOT_NUMBER), line)
			continue
		}

		// 入力範囲内か確認する
		if number <= 0 || number > MAX {
			continue
		}

		switch {
		case number == target_number:
			// 正解なので抜ける
			fmt.Printf(m.Get(HIT_MESSAGE), target_number)
			break FOR

		case number < target_number:
			fmt.Println(m.Get(TOO_SMALL))

		case number > target_number:
			fmt.Println(m.Get(TOO_BIG))

		default:
			fmt.Println(m.Get(UNDEFINED))
		}
	}
}
