package server

// Continuously sends messages form send channel to WS
func (c *Client) writePump() {
	defer c.Conn.Close() // Close connection when finished

	for msg := range c.Send {
		if err := c.Conn.WriteMessage(1, msg); err != nil {
			return // Exit on client disconnect or error
		}
	}
}
