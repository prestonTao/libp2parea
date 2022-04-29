package libp2parea

import (
	"github.com/prestonTao/libp2parea/engine"
	"github.com/prestonTao/libp2parea/message_center"
	"github.com/prestonTao/libp2parea/nodeStore"
	"github.com/prestonTao/libp2parea/virtual_node"
)

/*
	注册邻居节点消息，不转发
*/
func Register_neighbor(msgid uint64, handler message_center.MsgHandler) {
	message_center.Register_neighbor(msgid, handler)
}

/*
	发送一个新邻居节点消息
	@return    bool    消息是否发送成功
*/
func SendNeighborMsg(msgid uint64, recvid *nodeStore.AddressNet, content *[]byte) (*message_center.Message, error) {
	return message_center.SendNeighborMsg(msgid, recvid, content)
}

/*
	对某个消息回复
*/
func SendNeighborReplyMsg(message *message_center.Message, msgid uint64, content *[]byte, session engine.Session) error {
	return message_center.SendNeighborReplyMsg(message, msgid, content, session)
}

func SendNeighborWithReplyMsg(msgid uint64, recvid *nodeStore.AddressNet, content *[]byte, waitRequestClass string, timeout int64) (*[]byte, error) {
	return message_center.SendNeighborWithReplyMsg(msgid, recvid, content, waitRequestClass, timeout)
}

/*
	注册广播消息
*/
func Register_multicast(msgid uint64, handler message_center.MsgHandler) {
	message_center.Register_multicast(msgid, handler)
}

/*
	发送一个新的广播消息
*/
func SendMulticastMsg(msgid uint64, content *[]byte) bool {
	return message_center.SendMulticastMsg(msgid, content)

}

func BroadcastsAll(msgid, p2pMsgid uint64, whiltlistNodes, superNodes, proxyNodes []nodeStore.AddressNet, hash *[]byte) error {
	return message_center.BroadcastsAll(msgid, p2pMsgid, whiltlistNodes, superNodes, proxyNodes, hash)
}

/*
	注册消息
	从超级节点中搜索目标节点
*/
func Register_search_super(msgid uint64, handler message_center.MsgHandler) {
	message_center.Register_search_super(msgid, handler)
}

/*
	发送一个新的查找超级节点消息
*/
func SendSearchSuperMsg(msgid uint64, recvid *nodeStore.AddressNet, content *[]byte) (*message_center.Message, bool) {
	return message_center.SendSearchSuperMsg(msgid, recvid, content)
}

/*
	对某个消息回复
*/
func SendSearchSuperReplyMsg(message *message_center.Message, msgid uint64, content *[]byte) bool {
	return message_center.SendSearchSuperReplyMsg(message, msgid, content)
}

/*
	注册消息
	从所有节点中搜索节点，包括普通节点
*/
func Register_search_all(msgid uint64, handler message_center.MsgHandler) {
	message_center.Register_search_all(msgid, handler)
}

/*
	发送一个新的查找超级节点消息
*/
func SendSearchAllMsg(msgid uint64, recvid *nodeStore.AddressNet, content *[]byte) (*message_center.Message, bool) {
	return message_center.SendSearchAllMsg(msgid, recvid, content)
}

/*
	对某个消息回复
*/
func SendSearchAllReplyMsg(message *message_center.Message, msgid uint64, content *[]byte) bool {
	return message_center.SendSearchAllReplyMsg(message, msgid, content)
}

/*
	注册点对点通信消息
*/
func Register_p2p(msgid uint64, handler message_center.MsgHandler) {
	message_center.Register_p2p(msgid, handler)
}

/*
	发送一个新消息
	@return    *Message     返回的消息
	@return    bool         是否发送成功
	@return    bool         消息是发给自己
*/
func SendP2pMsg(msgid uint64, recvid *nodeStore.AddressNet, content *[]byte) (*message_center.Message, bool, bool) {
	return message_center.SendP2pMsg(msgid, recvid, content)
}

/*
	发送一个新消息
	是SendP2pMsg方法的定制版本，多了recvSuperId参数。
*/
func SendP2pMsgEX(msgid uint64, recvid, recvSuperId *nodeStore.AddressNet, content *[]byte) (*message_center.Message, bool) {
	return message_center.SendP2pMsgEX(msgid, recvid, recvSuperId, content)
}

/*
	对某个消息回复
*/
func SendP2pReplyMsg(message *message_center.Message, msgid uint64, content *[]byte) bool {
	return message_center.SendP2pReplyMsg(message, msgid, content)
}

/*
	注册点对点通信消息
*/
func Register_p2pHE(msgid uint64, handler message_center.MsgHandler) {
	message_center.Register_p2pHE(msgid, handler)
}

/*
	发送一个加密消息，包括消息头也加密
	@return    *Message     返回的消息
	@return    bool         是否发送成功
	@return    bool         消息是发给自己
*/
func SendP2pMsgHE(msgid uint64, recvid *nodeStore.AddressNet, content *[]byte) (*message_center.Message, bool, bool) {
	return message_center.SendP2pMsgHE(msgid, recvid, content)
}

/*
	对某个消息回复
*/
func SendP2pReplyMsgHE(message *message_center.Message, msgid uint64, content *[]byte) bool {
	return message_center.SendP2pReplyMsgHE(message, msgid, content)
}

/*
	注册虚拟节点搜索节点消息
*/
func Register_vnode_search(msgid uint64, handler message_center.MsgHandler) {
	message_center.Register_vnode_search(msgid, handler)
}

/*
	发送虚拟节点搜索节点消息
*/
func SendVnodeSearchMsg(msgid uint64, sendVnodeid, recvVnodeid *virtual_node.AddressNetExtend, content *[]byte) (*message_center.Message, bool) {
	return message_center.SendVnodeSearchMsg(msgid, sendVnodeid, recvVnodeid, content)
}

/*
	对发送虚拟节点搜索节点消息回复
*/
func SendVnodeSearchReplyMsg(message *message_center.Message, msgid uint64, content *[]byte, session engine.Session) error {
	return message_center.SendVnodeSearchReplyMsg(message, msgid, content, session)
}

/*
	注册虚拟节点之间点对点加密消息
*/
func Register_vnode_p2pHE(msgid uint64, handler message_center.MsgHandler) {
	message_center.Register_vnode_p2pHE(msgid, handler)
}

/*
	发送虚拟节点之间点对点消息
*/
func SendVnodeP2pMsgHE(msgid uint64, sendVnodeid, recvVnodeid *virtual_node.AddressNetExtend, content *[]byte) (*message_center.Message, bool) {
	return message_center.SendVnodeP2pMsgHE(msgid, sendVnodeid, recvVnodeid, content)
}

/*
	对发送虚拟节点之间点对点消息回复
*/
func SendVnodeP2pReplyMsgHE(message *message_center.Message, msgid uint64, content *[]byte, session engine.Session) error {
	return message_center.SendVnodeP2pReplyMsgHE(message, msgid, content, session)
}
