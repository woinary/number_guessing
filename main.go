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

	try_count := 0      // 試行回数
	break_flag := false // 中断フラグ
	history := make([]int, 0, 20)
FOR:
	for {
		// 1行入力
		fmt.Printf(m.Get(PROMPT_MESSAGE), MAX)
		line := readString(sc)
		try_count += 1

		// 中断コマンドのチェック
		if line == "q" {
			break_flag = true
			break
		}

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

		// ここまで来たら履歴に入れておく
		history = append(history, number)

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
	// 結果出力
	if break_flag {
		fmt.Printf(m.Get(BREAK_MESSAGE), try_count-1)
		os.Exit(2)
	}
	fmt.Printf(m.Get(END_MESSAGE), try_count)
	fmt.Print("input history: ")
	for i := range history {
		fmt.Print(history[i], " ")
	}
	fmt.Println("")
}
