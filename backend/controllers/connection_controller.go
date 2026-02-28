package controllers

import (
	"OmniBase/backend/database"
	"OmniBase/backend/services"
	"context"
)

// ConnectionController handles database connection operations
type ConnectionController struct {
	ctx context.Context
	db  *services.DatabaseService
}

// NewConnectionController creates a new ConnectionController
func NewConnectionController(db *services.DatabaseService) *ConnectionController {
	return &ConnectionController{
		db: db,
	}
}

// Startup is called when the Wails app starts
func (c *ConnectionController) Startup(ctx context.Context) {
	c.ctx = ctx
}

// TestConnection tests a database connection with the given config
func (c *ConnectionController) TestConnection(config database.Config) (string, error) {
	return c.db.TestConnection(config)
}
