package gonet

type MessageInfo struct {
	Topic Topic
	Data  interface{}
}

// SubscribeAll returns a channel that gets information about every message that passes
// through the network
func (c *Context) SubscribeAll() chan MessageInfo {
	n := c.net
	if c.Initialized {
		panic("SubscribeAll called after initialization is complete")
	}
	ch := make(chan MessageInfo, 10)
	n.universalsubscribers = append(n.universalsubscribers, ch)
	return ch
}
