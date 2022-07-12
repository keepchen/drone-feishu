package main

type Payload struct {
	//签名相关
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	MsgType   string `json:"msg_type"`

	//卡片相关
	Content Card `json:"card"`
}
