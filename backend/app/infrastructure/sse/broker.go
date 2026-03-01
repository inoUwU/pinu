package sse

import (
	"sync"

	"github.com/samber/do"
	"inoUwU/pinu/app/domain/port"
)

const defaultQueueCapacity = 100

// Broker SSEメッセージブローカーの実装（port.SSEBroker を実装）
type Broker struct {
	queue []string
	mu    sync.Mutex
}

// NewSSEBroker SSEBrokerを生成する
func NewSSEBroker(i *do.Injector) (port.SSEBroker, error) {
	return &Broker{
		queue: make([]string, 0, defaultQueueCapacity),
	}, nil
}

// Publish メッセージをキューに追加する
func (b *Broker) Publish(msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.queue = append(b.queue, msg)
}

// Consume キューからメッセージを取り出す
// 存在しない場合はfalseを返す
func (b *Broker) Consume() (string, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if len(b.queue) == 0 {
		return "", false
	}
	msg := b.queue[0]
	b.queue = b.queue[1:]
	return msg, true
}
