package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	configCommon "synapsis-challange/common/config"
	httpCommon "synapsis-challange/common/http"
	jwtCommon "synapsis-challange/common/jwt"
	passwordCommon "synapsis-challange/common/password"
	pgCommon "synapsis-challange/common/pg"
	productDelivery "synapsis-challange/internal/delivery/product/http"
	userDelivery "synapsis-challange/internal/delivery/user/http"
	productRepo "synapsis-challange/internal/repository/product/pg"
	userRepo "synapsis-challange/internal/repository/user/pg"
	productUsecase "synapsis-challange/internal/usecase/product"
	userUsecase "synapsis-challange/internal/usecase/user"
	"time"
)

func main() {
	cfg := configCommon.LoadConfig()
	pg, querier := pgCommon.New(cfg.DatabaseURL)
	defer pg.Close()

	jwtManager := jwtCommon.New(cfg.JWTAccessToken)
	passwordManager := passwordCommon.New()

	h := httpCommon.NewHTTPServer()
	api := h.E.Group("/api/v1", middleware.Logger(), middleware.CORS())

	ur := userRepo.NewPGUserRepository(querier)
	uc := userUsecase.NewUserUsecase(ur, jwtManager, passwordManager)
	userDelivery.NewHTTPUserDelivery(api, uc)

	pr := productRepo.NewPGProductRepository(querier)
	pc := productUsecase.NewProductUsecase(pr)
	productDelivery.NewHTTPProductDelivery(api, pc)

	h.E.Logger.Fatal(h.E.StartServer(&http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}))
}
