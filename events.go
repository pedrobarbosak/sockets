package sockets

func (c *Channel) OnConnection(fn func(any) any) {
	c.onConnection = fn
}

func (c *Channel) OnDisconnect(fn func(any)) {
	c.onDisconnect = fn
}

func (c *Channel) OnMessage(fn func(any, []byte) any) {
	c.onMessage = fn
}

func (c *Channel) OnBeforeSend(fn func(data any, msg any) (msgToSend any)) {
	c.onBeforeSend = fn
}
