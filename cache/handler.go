package cache

import (
	// "fmt"
	"github.com/prestonTao/libp2parea/engine"
	mc "github.com/prestonTao/libp2parea/message_center"
	"github.com/prestonTao/utils"
)

func syncData(c engine.Controller, msg engine.Packet) {
	message, err := mc.ParserMessage(&msg.Data, &msg.Dataplus, msg.MsgID)
	if err != nil {
		// fmt.Println(err)
		return
	}
	// form, _ := utils.FromB58String(msg.Session.GetName())
	form := utils.Bytes2string(msg.Session.GetName())
	if message.IsSendOther(&form) {
		return
	}
	//发送给自己的，自己处理
	if err := message.ParserContent(); err != nil {
		// fmt.Println(err)
		return
	}
	content := []byte("value")
	//回复给发送者
	mhead := mc.NewMessageHead(message.Head.Sender, message.Head.SenderSuperId, true)
	mbody := mc.NewMessageBody(&content, message.Body.CreateTime, message.Body.Hash, message.Body.SendRand)
	message = mc.NewMessage(mhead, mbody)
	message.Reply(MSGID_syncData_recv)
}
func syncData_recv(c engine.Controller, msg engine.Packet) {
	message, err := mc.ParserMessage(&msg.Data, &msg.Dataplus, msg.MsgID)
	if err != nil {
		// fmt.Println("error  1", err)
		return
	}
	// form, _ := utils.FromB58String(msg.Session.GetName())
	form := utils.Bytes2string(msg.Session.GetName())
	if message.IsSendOther(&form) {
		return
	}
	//发送给自己的，自己处理
	if err := message.ParserContent(); err != nil {
		engine.NLog.Error(engine.LOG_file, "%s", err.Error())
		engine.NLog.Error(engine.LOG_file, "%s", string(msg.Dataplus))
		return
	}
	mc.ResponseWait(mc.CLASS_syncdata, utils.Bytes2string(message.Body.Hash), message.Body.Content)
}
