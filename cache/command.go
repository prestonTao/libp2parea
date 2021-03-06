package cache

import (
	"errors"
	// "fmt"
	mc "github.com/prestonTao/libp2parea/message_center"
	"github.com/prestonTao/libp2parea/nodeStore"
	"github.com/prestonTao/utils"
)

func getQuarterLogicIds() []*utils.Multihash {
	return nodeStore.GetQuarterLogicIds(nodeStore.NodeSelf.IdInfo.Id)
}
func SyncDataToQuarterLogicIds() error {
	ids := getQuarterLogicIds()
	for _, val := range ids {
		// fmt.Println(val)
		sendMsg(val, []byte("okd"))
	}
	return nil
}
func sendMsg(id *utils.Multihash, data []byte) error {
	mhead := mc.NewMessageHead(id, id, false)
	mbody := mc.NewMessageBody(&data, "", nil, 0)
	message := mc.NewMessage(mhead, mbody)
	if message.Send(MSGID_syncData) {
		bs := mc.WaitRequest(mc.CLASS_syncdata, utils.Bytes2string(message.Body.Hash))
		if bs == nil {
			// fmt.Println("发送共享文件消息失败，可能超时")
			return errors.New("发送共享文件消息失败，可能超时")
		}
		return nil
	}
	return nil
}
