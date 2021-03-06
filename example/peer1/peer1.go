package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/prestonTao/keystore"
	"github.com/prestonTao/libp2parea"
	"github.com/prestonTao/libp2parea/engine"
	mc "github.com/prestonTao/libp2parea/message_center"
	"github.com/prestonTao/libp2parea/nodeStore"
)

func main() {
	fmt.Println("start")
	start()
}

func start() {
	addrPre := "SELF"
	keyPath := "key.txt"
	pwd := "123456789"
	//加载节点密钥
	// keystore.Load(keyPath)

	err := keystore.Load(keyPath, addrPre)
	if err != nil {
		keystore.CreateKeystore(keyPath, addrPre, pwd)
	}

	//
	// config.SetNetType(config.NetType_test)

	engine.Netid = 554433
	addr := "127.0.0.1"
	port := uint16(29980)

	// var port uint16
	// fmt.Println("请输入一个端口号:")
	// fmt.Scan(&port)
	args := os.Args
	if len(args) > 0 {
		portStr := args[0]
		portInt, err := strconv.Atoi(portStr)
		if err != nil {
			engine.Log.Info(err.Error())
		} else {
			port = uint16(portInt)
		}
	}

	libp2parea.Start(true, addr, port, keyPath, addrPre, pwd)
	InitHandler()
	// go sendMsg()
	// time.Sleep(time.Hour * 999999)
	select {}
}

const msg_text = 1000

func sendMsg() {
	engine.Log.Info("start sendMsg")
	// peer1ID := nodeStore.AddressFromB58String("B2Y23Gj621YbeNQ9yBpo9Cx9H8Hj4XcEm5AfYaMeqLks")
	// peer2ID := nodeStore.AddressFromB58String("2KZxDfLnycV65SkgbWPh7AUwwKyvLtpXzRrzDFqssusQ")
	peer3ID := nodeStore.AddressFromB58String("HfaQgqd1KZuqn1pAQubk7NtDMLRbhGcYPBRkbphPB98i")
	content := []byte("你好")

	<-time.NewTimer(time.Second * 20).C
	msg, sendOk, isSelf := libp2parea.SendP2pMsgHE(msg_text, &peer3ID, &content)
	engine.Log.Info("发送消息%v %t %t", msg, sendOk, isSelf)

	return

	for range time.NewTicker(time.Second * 5).C {
		msg, sendOk, isSelf := libp2parea.SendP2pMsgHE(msg_text, &peer3ID, &content)
		engine.Log.Info("发送消息%v %t %t", msg, sendOk, isSelf)
	}
}

func InitHandler() {
	libp2parea.Register_p2pHE(msg_text, RecvMsgHandler)
}

func RecvMsgHandler(c engine.Controller, msg engine.Packet, message *mc.Message) {
	engine.Log.Info("from:%s say:%s", message.Head.Sender.B58String(), string(*message.Body.Content))
}
