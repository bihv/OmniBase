package services

import (
	"OmniBase/backend/database"
	"context"
)

// DatabaseService exposes database operations to the frontend via Wails bindings
type DatabaseService struct {
	ctx  context.Context
	conn *database.Connection
}

// NewDatabaseService creates a new DatabaseService
func NewDatabaseService() *DatabaseService {
	return &DatabaseService{
		conn: database.NewConnection(),
	}
}

// Startup is called when the Wails app starts
func (s *DatabaseService) Startup(ctx context.Context) {
	s.ctx = ctx
	s.conn.SetContext(ctx)
}

// TestConnection tests a database connection
func (s *DatabaseService) TestConnection(config database.Config) (string, error) {
	return s.conn.TestConnection(config)
}
