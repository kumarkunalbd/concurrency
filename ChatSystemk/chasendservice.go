package ChatSystemk




type chatmessage struct{
	senderId  int
	chatMessager string
	ReceivedID int
}


type storedChatMessage struct{
	senderId  int
	chatMessager string
	ReceivedID int
}


func snedChat(chatMessage chatmessage) {
	urlRestAPIToSendChat := getAPIEndpiointForChat()
	chatBody ::= chat<chatMessage.chatMessager
	messageId := IDgeneratorForMessage()

	//
}



func IDgeneratorForMessage() int {
	// this wil generate id UUID for messaage
}

func storetheMessageinDB() {

}