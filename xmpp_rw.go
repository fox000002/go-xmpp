package xmpp

import "time"

// SetReadTimeOutInSeconds set ReadDeadline of conn
func (c *Client) SetReadTimeOutInSeconds(seconds int) {
	c.conn.SetReadDeadline(time.Now().Add(time.Second * time.Duration(seconds)))
}

// SetWriteTimeOutInSeconds set WriteDeadline of conn
func (c *Client) SetWriteTimeOutInSeconds(seconds int) {
	c.conn.SetWriteDeadline(time.Now().Add(time.Second * time.Duration(seconds)))
}
