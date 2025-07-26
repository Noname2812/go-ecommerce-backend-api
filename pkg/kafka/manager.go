package kafka

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type Manager struct {
	producer  *Producer
	Consumers map[string][]*ConsumerInfo
	logger    *zap.Logger
	mu        sync.RWMutex
	brokers   []string
}

type ConsumerInfo struct {
	Consumer    *Consumer
	Topic       string
	GroupID     string
	WorkerCount int
}

func NewManager(brokers []string, logger *zap.Logger) *Manager {
	return &Manager{
		producer:  NewProducer(brokers, logger),
		Consumers: make(map[string][]*ConsumerInfo),
		logger:    logger,
		brokers:   brokers,
	}
}

// ************* Consumer ******************** //

// Create consumer.
func (m *Manager) AddConsumer(
	topic, groupID string,
	handler MessageHandler,
	workerCount int,
	configOpts []ConfigOption,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// check if consumer already exists
	if consumers, exists := m.Consumers[topic]; exists {
		for _, consumerInfo := range consumers {
			if consumerInfo.GroupID == groupID {
				return fmt.Errorf("worker pool consumer for topic %s with groupID %s already exists", topic, groupID)
			}
		}
	}

	// create consumer
	consumer := NewConsumer(
		m.brokers,
		topic,
		groupID,
		handler,
		m.logger,
		workerCount,
		configOpts,
	)

	// create consumer info
	consumerInfo := &ConsumerInfo{
		Consumer:    consumer,
		Topic:       topic,
		GroupID:     groupID,
		WorkerCount: workerCount,
	}

	// add consumer info to map
	m.Consumers[topic] = append(m.Consumers[topic], consumerInfo)

	m.logger.Info("worker pool consumer added",
		zap.String("topic", topic),
		zap.String("group_id", groupID),
		zap.Int("worker_count", workerCount),
	)

	return nil
}

// Start all consumers
func (m *Manager) StartAllConsumers(ctx context.Context) {
	m.mu.RLock()
	allConsumers := make([]*ConsumerInfo, 0)
	for _, consumers := range m.Consumers {
		allConsumers = append(allConsumers, consumers...)
	}
	m.mu.RUnlock()

	// start all consumers
	for _, consumerInfo := range allConsumers {
		go func(ci *ConsumerInfo) {
			if err := ci.Consumer.Start(ctx); err != nil {
				m.logger.Error("worker pool consumer stopped with error",
					zap.String("topic", ci.Topic),
					zap.String("group_id", ci.GroupID),
					zap.Error(err),
				)
			}
		}(consumerInfo)
	}
}

// ******************** Producer *************

// Register topic
func (m *Manager) RegisterTopic(topic string, cfg TopicConfig) {
	m.producer.RegisterTopic(topic, cfg)
}

// Send message
func (m *Manager) SendMessage(ctx context.Context, topic string, key []byte, value interface{}) error {
	return m.producer.SendMessage(ctx, topic, key, value)
}

// Close
func (m *Manager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Close worker pool consumers
	for topic, consumers := range m.Consumers {
		for _, consumerInfo := range consumers {
			if err := consumerInfo.Consumer.Close(); err != nil {
				m.logger.Error("failed to close worker pool consumer",
					zap.String("topic", topic),
					zap.String("group_id", consumerInfo.GroupID),
					zap.Error(err),
				)
			}
		}
	}
	return m.producer.Close()
}
