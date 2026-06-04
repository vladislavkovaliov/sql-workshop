package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"shop-api/http/handlers"
	categoryrepo "shop-api/repository/category"
	orderrepo "shop-api/repository/order"
	productrepo "shop-api/repository/product"
	userrepo "shop-api/repository/user"
	categoryservice "shop-api/service/category"
	orderservice "shop-api/service/order"
	productservice "shop-api/service/product"
	userservice "shop-api/service/user"
)

func setupRouter(pool *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		wireProducts(api, pool)
		wireOrders(api, pool)
		wireCategories(api, pool)
		wireUsers(api, pool)
	}

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func wireProducts(rg *gin.RouterGroup, pool *pgxpool.Pool) {
	repo := productrepo.NewPgxRepository(pool)
	svc := productservice.New(repo)
	h := handlers.NewProductHandler(svc)

	rg.GET("/products", h.ListProducts)
	rg.POST("/products", h.CreateProduct)
	rg.GET("/products/cursor", h.ListCursorProducts)
	rg.GET("/products/revenue", h.TotalRevenue)
}

func wireOrders(rg *gin.RouterGroup, pool *pgxpool.Pool) {
	repo := orderrepo.NewPgxRepository(pool)
	svc := orderservice.New(repo)
	h := handlers.NewOrderHandler(svc)

	rg.GET("/orders", h.ListOrder)
}

func wireCategories(rg *gin.RouterGroup, pool *pgxpool.Pool) {
	repo := categoryrepo.NewPgxRepository(pool)
	svc := categoryservice.New(repo)
	h := handlers.NewCategoryHandler(svc)

	rg.GET("/category", h.ListCategory)
}

func wireUsers(rg *gin.RouterGroup, pool *pgxpool.Pool) {
	repo := userrepo.NewPgxRepository(pool)
	svc := userservice.New(repo)
	h := handlers.NewUserHandler(svc)

	rg.GET("/users", h.ListUsers)
	rg.GET("/users/cursor", h.ListCursorUsers)
	rg.GET("/users/search", h.SearchByEmail)

}
