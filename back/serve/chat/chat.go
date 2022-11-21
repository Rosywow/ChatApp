package chat

//author: {"name":"auth","email":"XUnion@GMail.com"}
//annotation:chat-service

import (
	"back/cmn"
	"back/serve/login"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

func Enroll(author string) {
	var developer *cmn.ModuleAuthor
	if author != "" {
		var d cmn.ModuleAuthor
		err := json.Unmarshal([]byte(author), &d)
		if err != nil {
			log.Println(err.Error())
			return
		}
		developer = &d
	}

	cmn.AddService(&cmn.ServeEndPoint{
		Fn: chat,

		Path: "/api/chat",
		Name: "chat",

		Developer: developer,
	})
	cmn.AddService(&cmn.ServeEndPoint{
		Fn: getUid,

		Path: "/api/getUid",
		Name: "getUid",

		Developer: developer,
	})
}

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	HandshakeTimeout: 3 * time.Second,  //5->3
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MessageType struct {
	Message string `json:"message"`
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Time string `json:"time"`
}
//
type ConnsType struct {
	Conns *websocket.Conn
	Sender string
	Receiver string
}
var conns []*ConnsType
var mutex sync.Mutex
// authMgmt authenticate/authorization management
func chat(ctx context.Context) {
	q := cmn.GetCtxValue(ctx)
	conn,err :=UP.Upgrade(q.W,q.R,nil)
	if err!=nil {
		fmt.Println(q.R.URL)
		log.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("聊天服务开始时间：",q.BeginTime)
	fmt.Println("聊天")
	fmt.Println("url:",q.R.URL)
	sender := q.R.URL.Query().Get("sender")
	receiver := q.R.URL.Query().Get("receiver")
	if sender==""&&receiver=="" {
		fmt.Println("sender or receiver not define")
		return
	}

	var Conn ConnsType
	var senders string
	senders = sender
	Conn.Conns = conn
	Conn.Receiver = receiver
	Conn.Sender = senders
	//exist := 0
	//for _,v :=range conns{
	//	if v==&Conn{
	//		exist = 1
	//	}
	//}
	//if exist!=1 {
	//	mutex.Lock()
	//	conns = append(conns,&Conn)
	//	mutex.Unlock()
	//}
	mutex.Lock()
	conns = append(conns,&Conn)
	mutex.Unlock()
	//fmt.Println("coons:",conns)
	//if err==nil{
	//	fmt.Println("connect successfully")
	//}
	for {
		//读取消息
		_, p, e :=conn.ReadMessage()
		if e!=nil {
			break
		}
		//fmt.Println("p",string(p))
		var Data MessageType
		err =json.Unmarshal(p,&Data)
		if err!=nil {
			fmt.Println("err",err)
		}
		//fmt.Println(Data)
		mutex.Lock()
		for key,value :=range conns{
			//fmt.Println("key",key)
			//fmt.Println("value",value)
			if Conn.Receiver==value.Sender {
				conns[key].Conns.WriteMessage(1,[]byte(fmt.Sprintf(`{"message":"%s","sender":"%s","time":"%s"}`,Data.Message,Conn.Sender,Data.Time)))
			}
		}
		mutex.Unlock()
	}
	//fmt.Println("connect 数组：",conns)
	//log.Println("关闭服务")
	//删除结束服务的连接
	mutex.Lock()
	for i,v:=range conns{
		if v.Conns == Conn.Conns {
			v.Conns.Close()
			conns = append(conns[:i],conns[i+1:]...)
			break
		}
	}
	//fmt.Println("connect 数组：",conns)
	mutex.Unlock()
}

func getUid(ctx context.Context){
	q := cmn.GetCtxValue(ctx)
	session, _ := login.Store.Get(q.R, "session-name")
	uid := session.Values["userid"]
	q.W.Write([]byte(fmt.Sprintf(`{"uid":%d}`,uid)))
}