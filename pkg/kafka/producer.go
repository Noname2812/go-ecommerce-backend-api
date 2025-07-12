// Package kafka implements Kafka producer and consumer functionality.
package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type ProducerOption func(*kafka.Writer)

// WithBalancer allows setting a custom balancer.
func WithBalancer(balancer kafka.Balancer) ProducerOption {
	return func(w *kafka.Writer) {
		w.Balancer = balancer
	}
}

// WithAsync allows enabling/disabling async sending.
func WithAsync(async bool) ProducerOption {
	return func(w *kafka.Writer) {
		w.Async = async
	}
}

// WithBatchSize allows setting the batch size.
func WithBatchSize(batchSize int) ProducerOption {
	return func(w *kafka.Writer) {
		w.BatchSize = batchSize
	}
}

// WithBatchTimeout allows setting the batch timeout.
func WithBatchTimeout(timeout time.Duration) ProducerOption {
	return func(w *kafka.Writer) {
		w.BatchTimeout = timeout
	}
}

// WithRequiredAcks allows configuring the required acks (0, 1, all).
func WithRequiredAcks(acks kafka.RequiredAcks) ProducerOption {
	return func(w *kafka.Writer) {
		w.RequiredAcks = acks
	}
}

type TopicConfig struct {
	Async        bool               // Enable asynchronous sending. If true, messages are sent without waiting for a response.
	RequiredAcks kafka.RequiredAcks //  0: No response needed (fastest, least safe). 1: Wait for leader broker only. -1 or All: Wait for all in-sync replicas (most reliable).
	BatchSize    int                // Maximum size (in bytes) of a batch before sending. Larger batches improve efficiency.
	BatchTimeout time.Duration      // Maximum wait time to collect messages for a batch before sending, even if BatchSize is not reached. Prevents delays during low traffic.
	Balancer     kafka.Balancer     // Balancer is used to distribute messages across partitions. (kafka.Hash, kafka.RoundRobin, kafka.LeastBytes).
}

type Producer struct {
	writers map[string]*kafka.Writer
	logger  *zap.Logger
	mu      sync.RWMutex
	brokers []string
}

// NewProducer creates a new Kafka producer with optional configs.
func NewProducer(brokers []string, logger *zap.Logger) *Producer {
	return &Producer{
		writers: make(map[string]*kafka.Writer),
		logger:  logger,
		brokers: brokers,
	}
}

// SendMessage -.
func (p *Producer) SendMessage(ctx context.Context, topic string, key []byte, value interface{}) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	p.mu.RLock()
	writer, ok := p.writers[topic]
	p.mu.RUnlock()

	if !ok {
		return fmt.Errorf("no writer registered for topic %s", topic)
	}

	msg := kafka.Message{
		Key:   key,
		Value: valueBytes,
		Time:  time.Now(),
	}

	if err := writer.WriteMessages(ctx, msg); err != nil {
		p.logger.Error("failed to write kafka message",
			zap.String("topic", topic),
			zap.ByteString("key", key),
			zap.Error(err),
		)
		return fmt.Errorf("failed to write message: %w", err)
	}

	p.logger.Info("kafka message sent",
		zap.String("topic", topic),
		zap.ByteString("key", key),
	)
	return nil
}

func (p *Producer) RegisterTopic(topic string, cfg TopicConfig) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Set default config values if not provided
	if cfg.BatchSize <= 0 {
		cfg.BatchSize = 1000 // default batch size
	}
	if cfg.BatchTimeout == 0 {
		cfg.BatchTimeout = 100 * time.Millisecond // default batch timeout
	}
	if cfg.RequiredAcks == 0 {
		cfg.RequiredAcks = kafka.RequireOne // default ack level
	}
	if cfg.Balancer == nil {
		cfg.Balancer = &kafka.RoundRobin{}
	}

	writer := &kafka.Writer{
		Addr:         kafka.TCP(p.brokers...),
		Topic:        topic,
		Async:        cfg.Async,
		RequiredAcks: cfg.RequiredAcks,
		BatchSize:    cfg.BatchSize,
		BatchTimeout: cfg.BatchTimeout,
		Balancer:     cfg.Balancer,
		Logger: kafka.LoggerFunc(func(msg string, args ...interface{}) {
			p.logger.Sugar().Infof(msg, args...)
		}),
	}

	p.writers[topic] = writer
	p.logger.Info("registered producer for topic",
		zap.String("topic", topic),
		zap.Bool("async", cfg.Async),
		zap.Any("acks", cfg.RequiredAcks),
	)
}

// Close - Closes all kafka.Writer instances
func (p *Producer) Close() error {
	var errs []error

	for topic, writer := range p.writers {
		if err := writer.Close(); err != nil {
			p.logger.Error("failed to close writer",
				zap.String("topic", topic),
				zap.Error(err),
			)
			errs = append(errs, fmt.Errorf("topic %s: %w", topic, err))
		} else {
			p.logger.Info("kafka writer closed",
				zap.String("topic", topic),
			)
		}
	}

	if len(errs) > 0 {
		// Return combined error
		return fmt.Errorf("failed to close some writers: %v", errs)
	}

	return nil
}
