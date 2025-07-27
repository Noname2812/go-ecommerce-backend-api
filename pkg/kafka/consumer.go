package kafka

import (
	"context"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

// MessageHandler -.
type MessageHandler func(ctx context.Context, key, value []byte) error

type Consumer struct {
	reader       *kafka.Reader
	logger       *zap.Logger
	handler      MessageHandler
	workerCount  int
	jobQueue     chan kafka.Message
	wg           sync.WaitGroup
	ctx          context.Context
	cancel       context.CancelFunc
	manualCommit bool
}

// ConfigOption applies to kafka.ReaderConfig
type ConfigOption func(*kafka.ReaderConfig)

// WithMinBytes sets the minimum bytes per fetch.
func WithMinBytes(minBytes int) ConfigOption {
	return func(cfg *kafka.ReaderConfig) {
		cfg.MinBytes = minBytes
	}
}

// WithMaxBytes sets the maximum bytes per fetch.
func WithMaxBytes(maxBytes int) ConfigOption {
	return func(cfg *kafka.ReaderConfig) {
		cfg.MaxBytes = maxBytes
	}
}

// WithMaxWait sets the maximum wait time for a fetch.
func WithMaxWait(timeout time.Duration) ConfigOption {
	return func(cfg *kafka.ReaderConfig) {
		cfg.MaxWait = timeout
	}
}

// WithStartOffset sets the start offset (Earliest or Latest).
func WithStartOffset(offset int64) ConfigOption {
	return func(cfg *kafka.ReaderConfig) {
		cfg.StartOffset = offset
	}
}

// WithCommitInterval sets the commit interval (0 = manual commit)
func WithCommitInterval(interval time.Duration) ConfigOption {
	return func(cfg *kafka.ReaderConfig) {
		cfg.CommitInterval = interval
	}
}
func NewConsumer(
	brokers []string,
	topic, groupID string,
	handler MessageHandler,
	logger *zap.Logger,
	workerCount int,
	configOpts []ConfigOption,
) *Consumer {
	// Default ReaderConfig
	readerConfig := kafka.ReaderConfig{
		Brokers:        brokers,
		Topic:          topic,
		GroupID:        groupID,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        1 * time.Second,
		CommitInterval: 0,
		StartOffset:    kafka.LastOffset,
		ErrorLogger: kafka.LoggerFunc(func(msg string, args ...interface{}) {
			logger.Sugar().Errorf(msg, args...)
		}),
	}

	// Apply Kafka ReaderConfig options
	for _, opt := range configOpts {
		opt(&readerConfig)
	}

	reader := kafka.NewReader(readerConfig)

	consumer := &Consumer{
		reader:       reader,
		logger:       logger,
		handler:      handler,
		workerCount:  workerCount,
		jobQueue:     make(chan kafka.Message, workerCount*10),
		manualCommit: readerConfig.CommitInterval == 0,
	}
	return consumer
}

// Start consumer
func (c *Consumer) Start(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)

	c.logger.Info("starting worker pool consumer",
		zap.String("topic", c.reader.Config().Topic),
		zap.String("group_id", c.reader.Config().GroupID),
		zap.Int("worker_count", c.workerCount),
	)

	// start workers
	for i := range c.workerCount {
		c.wg.Add(1)
		go c.worker(i)
	}

	// Message reader goroutine
	go c.messageReader()

	return nil
}

// messageReader - Read message from Kafka
func (c *Consumer) messageReader() {
	defer close(c.jobQueue)

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Info("stopping message reader")
			return
		default:
			m, err := c.reader.ReadMessage(c.ctx)
			if err != nil {
				if c.ctx.Err() != nil || err == context.DeadlineExceeded || err == context.Canceled {
					return
				}
				c.logger.Error("unexpected error while reading message", zap.Error(err))
				time.Sleep(10000 * time.Millisecond)
				continue
			}

			c.logger.Debug("received message",
				zap.String("topic", m.Topic),
				zap.Int("partition", m.Partition),
				zap.Int64("offset", m.Offset),
				zap.ByteString("key", m.Key),
			)

			// send message to job queue
			select {
			case c.jobQueue <- m:
			case <-c.ctx.Done():
				return
			}
		}
	}
}

// worker - Worker process
func (c *Consumer) worker(id int) {
	defer c.wg.Done()

	c.logger.Info("worker started", zap.Int("worker_id", id))

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Info("worker stopping", zap.Int("worker_id", id))
			return
		case msg, ok := <-c.jobQueue:
			if !ok {
				c.logger.Info("job queue closed, worker stopping", zap.Int("worker_id", id))
				return
			}

			start := time.Now()
			err := c.handler(c.ctx, msg.Key, msg.Value)
			duration := time.Since(start)

			if err != nil {
				c.logger.Error("failed to handle message",
					zap.Int("worker_id", id),
					zap.String("topic", msg.Topic),
					zap.Int("partition", msg.Partition),
					zap.Int64("offset", msg.Offset),
					zap.Duration("duration", duration),
					zap.Error(err),
				)
			} else {
				if c.manualCommit {
					if err := c.reader.CommitMessages(c.ctx, msg); err != nil {
						c.logger.Error("failed to commit message manually", zap.Error(err))
					}
					c.logger.Debug("message committed manually",
						zap.Int("worker_id", id),
						zap.String("topic", msg.Topic),
						zap.Int("partition", msg.Partition),
						zap.Int64("offset", msg.Offset),
					)
				}
				if !c.manualCommit {
					c.logger.Debug("message processed successfully",
						zap.Int("worker_id", id),
						zap.String("topic", msg.Topic),
						zap.Int("partition", msg.Partition),
						zap.Int64("offset", msg.Offset),
						zap.Duration("duration", duration),
					)
				}
			}
		}
	}
}

// Close
func (c *Consumer) Close() error {
	if c.cancel != nil {
		c.cancel()
	}

	// wait all workers complete
	c.wg.Wait()

	c.logger.Info("all workers stopped")
	return c.reader.Close()
}
