package libp2parea

import (
	"github.com/prestonTao/libp2parea/config"
	"github.com/prestonTao/libp2parea/message_center"
	"github.com/prestonTao/libp2parea/message_center/flood"
	"github.com/prestonTao/libp2parea/nodeStore"
	"github.com/prestonTao/libp2parea/utils"
)

func GetNodeMachineID(recvid *nodeStore.AddressNet) int64 {
	message, ok, isSendSelf := message_center.SendP2pMsg(config.MSGID_search_node, recvid, nil)
	if isSendSelf || !ok {
		return 0
	}

	// bs := flood.WaitRequest(config.CLASS_get_MachineID, hex.EncodeToString(message.Body.Hash), 0)
	bs, _ := flood.WaitRequest(config.CLASS_get_MachineID, utils.Bytes2string(message.Body.Hash), 0)
	if bs == nil {
		return 0
	}
	return utils.BytesToInt64(*bs)
}
