// @title			Shop API
// @version		1.0
// @description	Educational SQL workshop backend API
// @host			localhost:8080
// @BasePath		/api
package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "shop-api/docs"
	"shop-api/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	defer pool.Close()

	r := setupRouter(pool)

	log.Printf("Server starting on port %s", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
