package main

type MessageId int

type Messages map[MessageId]string

const (
	UNDEFINED MessageId = iota
	PROMPT_MESSAGE
	HIT_MESSAGE
	END_MESSAGE
	TOO_BIG
	TOO_SMALL
	NOT_NUMBER
)

// 各種メッセージの定義
var messages = map[MessageId]string{
	UNDEFINED:      "undefined error occured.",
	PROMPT_MESSAGE: "Please enter a number from 1 to %d\n",
	HIT_MESSAGE:    "Great! target is %d\n",
	END_MESSAGE:    "The number of attempts was %d.\n",
	TOO_BIG:        "too big.",
	TOO_SMALL:      "too small.",
	NOT_NUMBER:     "%s is not number.\n",
}

func (m *Messages) Get(id MessageId) string {
	message, ok := messages[id]
	if !ok {
		return messages[UNDEFINED]
	}
	return message
}
