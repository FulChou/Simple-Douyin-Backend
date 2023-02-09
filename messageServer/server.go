package messageServer

import (
	"Simple-Douyin-Backend/service/db"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
)

//type Message struct {
//	Id         int64  `json:"id,omitempty"`
//	Content    string `json:"content,omitempty"`
//	CreateTime string `json:"create_time,omitempty"`
//}

type SendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content"`
}

type PushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

var chatConnMap = sync.Map{}

func RunMessageServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		var event = SendEvent{}
		fmt.Printf("Receive Message：%+v, n: %d \n", string(buf[:n]), n)

		if err := json.Unmarshal(buf[:n], &event); err != nil {
			//fmt.Printf("Receive Message：%+v\n", event)
			fmt.Printf(" message format error , to json failed: %v\n", err)
			continue
		}
		fmt.Printf("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		err = db.CreateMessage(db.Message{UserID: uint(event.UserId), ToUserID: uint(event.ToUserId), Content: event.MsgContent})
		if err != nil {
			fmt.Println("store message to db fail")
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			fmt.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		pushEvent := PushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}
