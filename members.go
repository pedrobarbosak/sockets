package sockets

func (c *Channel) addMember(conn *Conn) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.members = append(c.members, conn)
}

func (c *Channel) removeMember(conn *Conn) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i, cc := range c.members {
		if cc.id == conn.id {
			l := len(c.members) - 1
			c.members[i] = c.members[l]
			c.members = c.members[:l]
			return
		}
	}
}
