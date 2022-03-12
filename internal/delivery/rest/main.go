package rest

import (
	"github.com/ridwanakf/tms-backend/internal/app"
	"github.com/ridwanakf/tms-backend/internal/delivery/rest/server"
	"github.com/ridwanakf/tms-backend/internal/delivery/rest/service"
)

// Start server
func Start(app *app.ShipmentApp) {
	srv := server.New()
	svc := service.GetServices(app)

	// Init Trucks Handlers
	trucks := srv.Group("/trucks")

	// Init Drivers Handlers
	drivers := srv.Group("/drivers")

	// Init Shipments Handlers
	shipments := srv.Group("/shipments")

	server.Start(srv, &app.Cfg.Server)
}
