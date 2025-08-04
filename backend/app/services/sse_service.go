package services

import (
	"github.com/samber/do"
	"sync"
)

const defaultQueueCapacity = 100

// SSEService はSSEメッセージキューを管理します
type SSEService struct {
	queue []string
	mu    sync.Mutex
}

// NewSSEService SSEServiceを生成します
func NewSSEService(i *do.Injector) (*SSEService, error) {
	return &SSEService{
		queue: make([]string, 0, defaultQueueCapacity),
	}, nil
}

// Publish メッセージをキューに追加します
func (s *SSEService) Publish(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.queue = append(s.queue, msg)
}

// Consume キューからメッセージを取り出します
// 存在しない場合はfalseを返します
func (s *SSEService) Consume() (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.queue) == 0 {
		return "", false
	}
	msg := s.queue[0]
	s.queue = s.queue[1:]
	return msg, true
}
