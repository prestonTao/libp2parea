package main

import (
	"fmt"
	"time"

	"github.com/prestonTao/libp2parea/message_center/flood"

	"github.com/prestonTao/utils"

	"github.com/prestonTao/keystore"
	"github.com/prestonTao/libp2parea"

	// "github.com/prestonTao/libp2parea/config"
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
	port := uint16(29982)

	libp2parea.Start(true, addr, port, keyPath, addrPre, pwd)
	InitHandler()
	// go sendMsg()
	select {}
}

const msg_text = 1000
const msg_ping = 1001
const msg_pong = 1002

const CLASS_ping_pong = "ping_pong"

func sendMsg() {
	engine.Log.Info("start sendMsg")
	// peer1ID := nodeStore.AddressFromB58String("B2Y23Gj621YbeNQ9yBpo9Cx9H8Hj4XcEm5AfYaMeqLks")
	peer2ID := nodeStore.AddressFromB58String("2KZxDfLnycV65SkgbWPh7AUwwKyvLtpXzRrzDFqssusQ")
	// peer3ID := nodeStore.AddressFromB58String("HfaQgqd1KZuqn1pAQubk7NtDMLRbhGcYPBRkbphPB98i")
	// content := []byte("你好，我是peer3")

	index := uint64(0)

	// <-time.NewTimer(time.Second * 10).C
	// msg, sendOk, isSelf := libp2parea.SendP2pMsgHE(msg_text, &peer2ID, &content)
	// engine.Log.Info("发送消息%v %t %t", msg, sendOk, isSelf)

	// return

	for range time.NewTicker(time.Second * 5).C {
		index++
		bs := utils.Uint64ToBytes(index)

		libp2parea.SendP2pMsgHE(msg_text, &peer2ID, &bs)
		// engine.Log.Info("发送消息%v %t %t", msg, sendOk, isSelf)
	}
}

func InitHandler() {
	libp2parea.Register_p2pHE(msg_text, RecvMsgHandler)

	libp2parea.Register_p2pHE(msg_ping, RecvMsgPing)
	libp2parea.Register_p2pHE(msg_pong, RecvMsgPong)
}

var index = uint64(0)

func RecvMsgHandler(c engine.Controller, msg engine.Packet, message *mc.Message) {
	recvIndex := utils.BytesToUint64(*message.Body.Content)
	if recvIndex-index > 1 {
		engine.Log.Info("收到消息不连续，跳过%d", recvIndex-index)
	}
	index = recvIndex

	if index%10000 == 0 {
		engine.Log.Info("from:%s say:%d", message.Head.Sender.B58String(), recvIndex)
	}
}

func RecvMsgPing(c engine.Controller, msg engine.Packet, message *mc.Message) {
	libp2parea.SendP2pReplyMsgHE(message, msg_pong, message.Body.Content)
}

func RecvMsgPong(c engine.Controller, msg engine.Packet, message *mc.Message) {
	flood.ResponseWait(CLASS_ping_pong, utils.Bytes2string(message.Body.Hash), message.Body.Content)
}
