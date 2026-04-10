
package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	senderok := (mToSend.sender.name != "") &&  (mToSend.sender.number != 0)
	recipientok := (mToSend.recipient.name != "") &&  (mToSend.recipient.number != 0)
	return senderok && recipientok
}
