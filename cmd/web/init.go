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
	uuidCommon "synapsis-challange/common/uuid"
	xenditCommon "synapsis-challange/common/xendit"
	cartItemDelivery "synapsis-challange/internal/delivery/cart-item/http"
	orderDelivery "synapsis-challange/internal/delivery/order/http"
	productDelivery "synapsis-challange/internal/delivery/product/http"
	userDelivery "synapsis-challange/internal/delivery/user/http"
	cartItemRepo "synapsis-challange/internal/repository/cart-item/pg"
	orderRepo "synapsis-challange/internal/repository/order/pg"
	productRepo "synapsis-challange/internal/repository/product/pg"
	userRepo "synapsis-challange/internal/repository/user/pg"
	cartItemUsecase "synapsis-challange/internal/usecase/cart-item"
	orderUsecase "synapsis-challange/internal/usecase/order"
	productUsecase "synapsis-challange/internal/usecase/product"
	userUsecase "synapsis-challange/internal/usecase/user"
	"time"
)

func main() {
	cfg := configCommon.LoadConfig()
	store := pgCommon.New(cfg.DatabaseURL)
	defer store.Db.Close()

	jwtManager := jwtCommon.New(cfg.JWTAccessToken)
	passwordManager := passwordCommon.New()
	xenditManager := xenditCommon.NewXenditManager(cfg.XenditSecretKey)
	uuidGenerator := uuidCommon.New()

	h := httpCommon.NewHTTPServer()
	api := h.E.Group("/api/v1", middleware.Logger(), middleware.CORS())

	pr := productRepo.NewPGProductRepository(store.Querier)
	pc := productUsecase.NewProductUsecase(pr)
	productDelivery.NewHTTPProductDelivery(api, pc)

	cir := cartItemRepo.NewPGPCartItemRepository(store.Querier)
	cic := cartItemUsecase.NewCartItemUsecase(cir, pr)
	cartItemDelivery.NewHTTPCartItemDelivery(api, cic, jwtManager)

	ur := userRepo.NewPGUserRepository(store.Querier)
	uc := userUsecase.NewUserUsecase(ur, jwtManager, passwordManager)
	userDelivery.NewHTTPUserDelivery(api, uc, cic, jwtManager)

	or := orderRepo.NewPGOrderRepository(store)
	oc := orderUsecase.NewOrderUsecase(or, cir, pr, xenditManager, uuidGenerator)
	orderDelivery.NewHTTPOrderDelivery(api, oc, jwtManager)

	h.E.Logger.Fatal(h.E.StartServer(&http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}))
}
