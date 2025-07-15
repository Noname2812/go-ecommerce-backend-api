package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/Noname2812/go-ecommerce-backend-api/cmd/swag/docs"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/initialize"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count_total",
		Help: "Total number of ping requests.",
	},
)

// @title           API Documentation Ecommerce Backend SHOPDEVGO
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/anonystick/go-ecommerce-backend-go

// @contact.name   TEAM TIPSGO
// @contact.url    github.com/anonystick/go-ecommerce-backend-go
// @contact.email  tipsgo@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1
// @schema http
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// handle OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// load configuration
	config := initialize.LoadConfig()

	// init app contaier
	appContainer, err := initialize.NewAppContainer(config)
	if err != nil {
		panic("Init app contaier failed")
	}

	// init grpc server
	router.InitGrpcServer(config, appContainer)

	// init kafka
	initialize.InitKafka(ctx, config, appContainer)

	// init http router
	r := router.InitHttpRouter(config, appContainer)

	go func() {
		<-sigChan
		cancel()
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		var wg sync.WaitGroup
		wg.Add(2)

		// closing kafka
		go func() {
			defer wg.Done()
			if err := appContainer.KafkaManager.Close(); err != nil {
				appContainer.Logger.Error("Failed to close Kafka manager", zap.Error(err))
			} else {
				appContainer.Logger.Info("Kafka manager closed successfully")
			}
		}()

		// closing grpc
		go func() {
			defer wg.Done()
			if err := appContainer.GRPCServerManager.Shutdown(shutdownCtx); err != nil {
				appContainer.Logger.Error("Failed to shut down gRPC", zap.Error(err))
			} else {
				appContainer.Logger.Info("gRPC manager shut down successfully")
			}
		}()

		// wait until both goroutines done
		wg.Wait()

		appContainer.Logger.Info("All services stopped. Exiting.")
		os.Exit(0)
	}()

	// prometheus
	prometheus.MustRegister(pingCounter)

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8002"); err != nil {
		panic(err)
	}
}
