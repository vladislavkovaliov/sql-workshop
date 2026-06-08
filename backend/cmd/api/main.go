// @title			Shop API
// @version		1.0
// @description	Educational SQL workshop backend API
// @host			localhost:8080
// @BasePath		/api
package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "shop-api/docs"
	"shop-api/internal/config"
	"shop-api/internal/events"

	orderrepo "shop-api/repository/order"
)

func main() {
	cfg := config.LoadConfig()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	defer pool.Close()

	brokers := []string{cfg.KafkaBrockers}

	producer := events.NewProducer(brokers)

	defer producer.Close()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	orderRepo := orderrepo.NewPgxRepository(pool)

	go events.StartConsumer(ctx, brokers, "shop-api", orderRepo)

	r := setupRouter(pool, producer)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Server starting on port %s", cfg.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	<-ctx.Done()

	log.Printf("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
}
