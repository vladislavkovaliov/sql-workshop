// @title			Shop API
// @version		1.0
// @description	Educational SQL workshop backend API
// @host			localhost:8080
// @BasePath		/api
package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "shop-api/docs"
	"shop-api/http/handlers"
	"shop-api/internal/config"
	productrepo "shop-api/repository/product"
	productservice "shop-api/service/product"
)

func main() {
	cfg := config.LoadConfig()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)

	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	defer pool.Close()

	repo := productrepo.NewPgxRepository(pool)
	svc := productservice.New(repo)
	handler := handlers.NewProductHandler(svc)

	r := gin.Default()
	r.GET("/api/products", handler.ListProducts)
	r.GET("/api/products/cursor", handler.ListCursorProducts)
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server starting on port %s", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
