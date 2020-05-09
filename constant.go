package rabbitmqclient

import "github.com/fairyhunter13/amqpwrapper"

// List of all uint constants for boolean
const (
	FalseUint uint64 = iota
	TrueUint
)

// List of default config used for this library.
const (
	DefaultPrefixExchange = "amqp."
	DefaultQueue          = "default"
	DefaultTopic          = "default"
)

const (
	// DefaultKeyInitiator sets the key channel for the initiator of topology in the producer connection.
	DefaultKeyInitiator = "initiator"
	// DefaultKeyProducer specifies the key channel for producer.
	DefaultKeyProducer = "producer.%d"
	// DefaultTypeProducer specifies the default type of channel for the amqpwrapper.
	DefaultTypeProducer = amqpwrapper.Producer
)

// List of all exchange type in rabbitmq
const (
	TypeDirect  = `direct`
	TypeFanout  = `fanout`
	TypeTopic   = `topic`
	TypeHeaders = `headers`
)

// List of all delivery modes for amqp rabbitmqclient
const (
	DeliveryModeTransient  = 1
	DelvieryModePersistent = 2
)

// List of all topology constants
const (
	TopologyExchangeDeclare        = `ExchangeDeclare`
	TopologyExchangeDeclarePassive = `ExchangeDeclarePassive`
	TopologyExchangeBind           = `ExchangeBind`
	TopologyExchangeUnbind         = `ExchangeUnbind`
	TopologyExchangeDelete         = `ExchangeDelete`
	TopologyQueueDeclare           = `QueueDeclare`
	TopologyQueueDeclarePassive    = `QueueDeclarePassive`
	TopologyQueueBind              = `QueueBind`
	TopologyQueueUnbind            = `QueueUnbind`
	TopologyQueueDelete            = `QueueDelete`
)
