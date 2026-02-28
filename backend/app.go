package backend

import (
	"OmniBase/backend/controllers"
	"OmniBase/backend/services"
	"context"
)

// App struct - main application lifecycle controller
type App struct {
	ctx context.Context

	// Controllers (bound to Wails frontend)
	ConnCtrl   *controllers.ConnectionController
	QueryCtrl  *controllers.QueryController
	SchemaCtrl *controllers.SchemaController

	// Internal services
	db *services.DatabaseService
}

// NewApp creates a new App and initializes all controllers
func NewApp() *App {
	db := services.NewDatabaseService()

	return &App{
		db:         db,
		ConnCtrl:   controllers.NewConnectionController(db),
		QueryCtrl:  controllers.NewQueryController(),
		SchemaCtrl: controllers.NewSchemaController(),
	}
}

// Startup is called when the Wails app starts.
// It distributes the context to all controllers and services.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.db.Startup(ctx)
	a.ConnCtrl.Startup(ctx)
	a.QueryCtrl.Startup(ctx)
	a.SchemaCtrl.Startup(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return "Hello " + name + ", welcome to OmniBase!"
}
