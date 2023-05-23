package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	configCommon "synapsis-challange/common/config"
	httpCommon "synapsis-challange/common/http"
	jwtCommon "synapsis-challange/common/jwt"
	pgCommon "synapsis-challange/common/pg"
	"time"
)

func main() {
	cfg := configCommon.LoadConfig()
	pg, querier := pgCommon.New(cfg.DatabaseURL)
	defer pg.Close()

	jwtManager := jwtCommon.New(cfg.JWTAccessToken)

	h := httpCommon.NewHTTPServer()
	api := h.E.Group("/api/v1", middleware.Logger(), middleware.CORS())

	h.E.Logger.Fatal(h.E.StartServer(&http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}))
}
