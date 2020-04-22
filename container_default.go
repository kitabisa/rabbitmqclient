package rabbitmqclient

func (c *Container) setDefaultExchange() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.globalExchange == nil {
		c.globalExchange = new(ExchangeDeclareArgs).Default()
	}
}

func (c *Container) setDefaultTopology() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.topology == nil {
		c.topology = NewTopology()
	}
}