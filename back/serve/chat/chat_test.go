package chat

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	dl := websocket.Dialer{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	//限制通道最大值为2000
	ch :=make(chan int,5000)
	sumTime := time.NewTicker(5*time.Minute)

FOR:
	for {
		//检查有没有超过测试时长
		select {
		case <-sumTime.C:
			{
				//已经超过总时长
				log.Println("测试结束")
				t.Logf("通过测试")
				break FOR
			}
		default:
			{
				//还没超过总时长，什么都不执行
			}

		}


		t1 := time.NewTicker(3*time.Second)
		ch <- 1
		go func (ticker *time.Ticker){
			select {
			case <-ticker.C:
				{
					//单个响应已经超时了
					log.Fatal("响应时间大于三秒")
					<-ch
					t.Error("单次响应时间大于三秒")
				}
			default:
				{
					//单个响应还没超时
					url := "ws://localhost:9090/api/chat?sender=1&receiver=2"
					c, res, err := dl.Dial(url, nil)
					if err != nil {
						log.Println("连接失败:", err)
						t.Logf("连接失败")
					}else {
						fmt.Printf("响应:%s", fmt.Sprint(res))
						c.Close()
					}
					<-ch
				}
			}
		}(t1)
	}
}



var testConns []ConnsType

func TestSend(t *testing.T) {
	log.Println("开始测试")
	//var connections []*websocket.Conn
	//for i:=0;i<600;i++ {
	//		//url :=fmt.Sprintf("ws://localhost:9090/api/chat?sender=%d&receiver=%d",a,b)
	//		url := "ws://localhost:9090/api/chat"
	//		c, _, err := websocket.DefaultDialer.Dial(url, nil)
	//		if err!=nil {
	//			fmt.Println("connect err:",err)
	//		}
	//		connections = append(connections,c)
	//}

	//defer func() {
	//	for _,v:= range connections {
	//		v.Close()
	//	}
	//}()
	//fmt.Println("暂停一下")
	//for _,v:=range connections {
	//	v.Close()
	//}
	time.Sleep(5*time.Minute)
	//提前建立2000个长连接 for循环建立一个conn数组存储coon，senderId，receiverId
	a,b := 1,2
	for b<=600 {
		url :=fmt.Sprintf("ws://localhost:9090/api/chat?sender=%d&receiver=%d",a,b)
		c, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err!=nil {
			fmt.Println("connect err:",err)
		} else{
			Coon := ConnsType{
				Conns: c,
				Sender: fmt.Sprintf("%d",a),
				Receiver: fmt.Sprintf("%d",b),
			}
			testConns=append(testConns,Coon)
		}
		//fmt.Println("res:",res)
		//长连接存放在数组中
		if a<b {
			temp:=a
			a=b
			b=temp
		}else {
			a+=1
			b+=3
		}
	}

	//清空数组
	defer func() {
		testConns=nil
	}()
	//关闭连接
	defer func() {
		for _,v:=range testConns{
			v.Conns.Close()
		}
	}()



	for _,v:=range testConns{
		go func(connect ConnsType){
			for{
				sendTime := time.Now()
				connect.Conns.WriteMessage(1,[]byte(fmt.Sprintf(`{"message":"你好","senderid":"%s","time":"%s"}`,v.Sender,sendTime.Format("01-02-2006 15:04:05"))))
				_,p,err :=connect.Conns.ReadMessage()
				if err!=nil {
					fmt.Println("1",err)
					break
				}
				var Data MessageType
				err = json.Unmarshal(p,&Data)
				if err!=nil {
					fmt.Println("err",err)
				}

				//取出对方发送的消息的事件
				getSendTime,err := time.ParseInLocation("01-02-2006 15:04:05",Data.Time,time.Local)
				if err!=nil {
					fmt.Println("time.ParseInLocation err:",err)
				}

				//如果发送事件和接收事件超过x秒，则说明失败
				if time.Now().Sub(getSendTime) > 3*time.Second {
					log.Println("单次响应超时")
					return
				}

				fmt.Println("senderId:",Data.Sender,"message:",Data.Message)
				//每三秒发一次消息
				time.Sleep(3*time.Second)
			}
		}(v)
	}
	//开2000个协程轮询
	time.Sleep(1*time.Minute)
}

//func TestName(t *testing.T) {
//	a,b:=0,1
//	for  b<600{
//		go func(a int,b int) {
//			url :=fmt.Sprintf("ws://localhost:9090/api/chat?sender=%d&receiver=%d",a,b)
//			c, _, err := websocket.DefaultDialer.Dial(url, nil)
//			if err!=nil {
//				fmt.Println("connect err:",err)
//			}
//			for{
//				sendTime := time.Now()
//				err = c.WriteMessage(1,[]byte(fmt.Sprintf(`{"message":"你好","senderid":"%d","time":"%s"}`,a,sendTime.Format("01-02-2006 15:04:05"))))
//				if err!=nil {
//					log.Println("WriteMessage err :",err)
//				}
//				_,p,err :=c.ReadMessage()
//				if err!=nil {
//					fmt.Println("1",err)
//					break
//				}
//				var Data MessageType
//				err = json.Unmarshal(p,&Data)
//				if err!=nil {
//					fmt.Println("err",err)
//				}
//
//				//取出对方发送的消息的事件
//				getSendTime,err := time.ParseInLocation("01-02-2006 15:04:05",Data.Time,time.Local)
//				if err!=nil {
//					fmt.Println("time.ParseInLocation err:",err)
//				}
//
//				//如果发送事件和接收事件超过x秒，则说明失败
//				if time.Now().Sub(getSendTime) > 3*time.Second {
//					log.Println("单次响应超时")
//					return
//				}
//
//				fmt.Println("senderId:",Data.Sender,"message:",Data.Message)
//				//每三秒发一次消息
//				time.Sleep(3*time.Second)
//			}
//		}(a,b)
//
//		temp:=a
//		a=b
//		b=temp
//
//		go func(a int,b int) {
//			url :=fmt.Sprintf("ws://localhost:9090/api/chat?sender=%d&receiver=%d",a,b)
//			c, _, err := websocket.DefaultDialer.Dial(url, nil)
//			if err!=nil {
//				fmt.Println("connect err:",err)
//			}
//		}(a,b)
//
//		a+=1
//		b+=3
//	}
//}
//
