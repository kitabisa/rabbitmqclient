package rabbitmqclient

import (
	"sync"
	"time"

	"github.com/fairyhunter13/rabbitmqclient/args"
)

// Topology contains all declarations needed to define the topology in the rabbitmq.
type Topology struct {
	mutex *sync.RWMutex
	// Mutex protects the following fields
	exchangeDeclareArgs        []args.ExchangeDeclare
	exchangeDeclarePassiveArgs []args.ExchangeDeclarePassive
	queueDeclareArgs           []args.QueueDeclare
	queueBindArgs              []args.QueueBind
	currentTime                *time.Time
	lastTime                   *time.Time
}

// NewTopology creates a new topology
func NewTopology() *Topology {
	now := time.Now()
	return &Topology{
		mutex:       new(sync.RWMutex),
		currentTime: &now,
		lastTime:    &now,
	}
}

func (t *Topology) update() *Topology {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	*t.currentTime = time.Now()
	return t
}

func (t *Topology) syncTime() {
	t.lastTime = t.currentTime
}

// IsUpdated checks if the topology has been updated or not.
// IsUpdated also automatically sync the time of last time to the current time if it is updated.
func (t *Topology) IsUpdated() (result bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	result = t.currentTime.After(*t.lastTime)
	if result {
		t.syncTime()
	}
	return
}

// AddExchangeDeclare add the exchange declare args to the topology.
func (t *Topology) AddExchangeDeclare(arg args.ExchangeDeclare) *Topology {
	t.mutex.Lock()
	t.exchangeDeclareArgs = append(t.exchangeDeclareArgs, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// GetExchangeDeclare return the exchange declare args inside the topology.
func (t *Topology) GetExchangeDeclare() []args.ExchangeDeclare {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.exchangeDeclareArgs[:0:0], t.exchangeDeclareArgs...)
}

// AddQueueDeclare adds the queue declaration into the topology
func (t *Topology) AddQueueDeclare(arg args.QueueDeclare) *Topology {
	t.mutex.Lock()
	t.queueDeclareArgs = append(t.queueDeclareArgs, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// GetQueueDeclare gets the queue declaration inside the topology.
func (t *Topology) GetQueueDeclare() []args.QueueDeclare {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueDeclareArgs[:0:0], t.queueDeclareArgs...)
}

// AddQueueBind adds the queue bind args to the topology
func (t *Topology) AddQueueBind(arg args.QueueBind) *Topology {
	t.mutex.Lock()
	t.queueBindArgs = append(t.queueBindArgs, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// GetQueueBind gets the queue bind args inside the topology
func (t *Topology) GetQueueBind() []args.QueueBind {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueBindArgs[:0:0], t.queueBindArgs...)
}
