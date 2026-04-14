package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func analyzeMessage(a *Analytics, m Message){
	s := m.Success
	if s {
		a.MessagesSucceeded++
		a.MessagesTotal++
	}else {
		a.MessagesFailed++
		a.MessagesTotal++		
	}
}
