package main

import (
	"fmt"
	"time"

	"github.com/prestonTao/keystore"
	"github.com/prestonTao/libp2parea"
	"github.com/prestonTao/libp2parea/engine"
	mc "github.com/prestonTao/libp2parea/message_center"
	"github.com/prestonTao/libp2parea/message_center/flood"
	"github.com/prestonTao/libp2parea/nodeStore"
	"github.com/prestonTao/utils"
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
	port := uint16(29981)

	libp2parea.Start(true, addr, port, keyPath, addrPre, pwd)
	InitHandler()
	// go sendMsg()
	go pingpong()
	select {}
}

const msg_text = 1000
const msg_ping = 1001
const msg_pong = 1002

const CLASS_ping_pong = "ping_pong"

func pingpong() {
	engine.Log.Info("start ping pong")
	peer3ID := nodeStore.AddressFromB58String("HfaQgqd1KZuqn1pAQubk7NtDMLRbhGcYPBRkbphPB98i")

	index := uint64(0)

	start := time.Now()
	for i := 0; i < 1000; i++ {
		index++
		bs := utils.Uint64ToBytes(index)

		msg, sendOk, _ := libp2parea.SendP2pMsgHE(msg_ping, &peer3ID, &bs)
		if !sendOk {
			engine.Log.Error("消息未发送成功")
			continue
		}
		_, err := flood.WaitRequest(CLASS_ping_pong, utils.Bytes2string(msg.Body.Hash), 1)
		if err != nil {
			engine.Log.Error("消息接收失败:%s", err.Error())
			continue
		}
		// engine.Log.Info("发送消息%v %t %t", msg, sendOk, isSelf)
		if index%1000 == 0 {
			engine.Log.Info("index:%d", index)
		}
	}
	engine.Log.Info("耗时:%s", time.Now().Sub(start))
}

func sendMsg() {
	engine.Log.Info("start sendMsg")
	// peer1ID := nodeStore.AddressFromB58String("B2Y23Gj621YbeNQ9yBpo9Cx9H8Hj4XcEm5AfYaMeqLks")
	// peer2ID := nodeStore.AddressFromB58String("2KZxDfLnycV65SkgbWPh7AUwwKyvLtpXzRrzDFqssusQ")
	peer3ID := nodeStore.AddressFromB58String("HfaQgqd1KZuqn1pAQubk7NtDMLRbhGcYPBRkbphPB98i")
	// content := []byte("你好，我是peer2")
	index := uint64(0)
	// <-time.NewTimer(time.Second * 10).C
	// msg, sendOk, isSelf := libp2parea.SendP2pMsgHE(msg_text, &peer3ID, &content)
	// engine.Log.Info("发送消息%v %t %t", msg, sendOk, isSelf)

	// return

	// index := 0

	for range time.NewTicker(time.Second / 200).C {
		// logicNodes := nodeStore.GetLogicNodes()
		// for _, one := range logicNodes {
		// 	engine.Log.Info("logic id:%s", one.B58String())
		// }

		index++
		bs := utils.Uint64ToBytes(index)

		// msg, sendOk, isSelf :=
		libp2parea.SendP2pMsgHE(msg_text, &peer3ID, &bs)
		// engine.Log.Info("发送消息%v %t %t", msg, sendOk, isSelf)
	}
}

func InitHandler() {
	libp2parea.Register_p2pHE(msg_text, RecvMsgHandler)

	libp2parea.Register_p2pHE(msg_ping, RecvMsgPing)
	libp2parea.Register_p2pHE(msg_pong, RecvMsgPong)
}

func RecvMsgHandler(c engine.Controller, msg engine.Packet, message *mc.Message) {
	engine.Log.Info("from:%s say:%s", message.Head.Sender.B58String(), string(*message.Body.Content))
}

func RecvMsgPing(c engine.Controller, msg engine.Packet, message *mc.Message) {
	libp2parea.SendP2pReplyMsgHE(message, msg_pong, message.Body.Content)
}

func RecvMsgPong(c engine.Controller, msg engine.Packet, message *mc.Message) {
	flood.ResponseWait(CLASS_ping_pong, utils.Bytes2string(message.Body.Hash), message.Body.Content)
}
