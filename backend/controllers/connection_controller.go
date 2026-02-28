package controllers

import (
	"OmniBase/backend/database"
	"OmniBase/backend/models"
	"OmniBase/backend/services"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ConnectionController handles database connection operations
type ConnectionController struct {
	ctx      context.Context
	db       *services.DatabaseService
	filePath string
}

// NewConnectionController creates a new ConnectionController
func NewConnectionController(db *services.DatabaseService) *ConnectionController {
	// Determine storage path
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = "."
	}
	storageDir := filepath.Join(configDir, "OmniBase")
	os.MkdirAll(storageDir, 0755)

	return &ConnectionController{
		db:       db,
		filePath: filepath.Join(storageDir, "connections.json"),
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

// SaveConnection saves a new connection or updates an existing one
func (c *ConnectionController) SaveConnection(conn models.ConnectionInfo) (models.ConnectionInfo, error) {
	conns, err := c.loadConnections()
	if err != nil {
		conns = []models.ConnectionInfo{}
	}

	if conn.ID == "" {
		conn.ID = fmt.Sprintf("conn_%d", time.Now().UnixNano())
		conns = append(conns, conn)
	} else {
		found := false
		for i, existing := range conns {
			if existing.ID == conn.ID {
				conns[i] = conn
				found = true
				break
			}
		}
		if !found {
			conns = append(conns, conn)
		}
	}

	if err := c.saveConnections(conns); err != nil {
		return conn, fmt.Errorf("failed to save connection: %w", err)
	}
	return conn, nil
}

// GetConnections returns all saved connections
func (c *ConnectionController) GetConnections() ([]models.ConnectionInfo, error) {
	conns, err := c.loadConnections()
	if err != nil {
		return []models.ConnectionInfo{}, nil
	}
	return conns, nil
}

// DeleteConnection deletes a connection by ID
func (c *ConnectionController) DeleteConnection(id string) error {
	conns, err := c.loadConnections()
	if err != nil {
		return fmt.Errorf("failed to load connections: %w", err)
	}

	filtered := make([]models.ConnectionInfo, 0, len(conns))
	for _, conn := range conns {
		if conn.ID != id {
			filtered = append(filtered, conn)
		}
	}

	return c.saveConnections(filtered)
}

func (c *ConnectionController) loadConnections() ([]models.ConnectionInfo, error) {
	data, err := os.ReadFile(c.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.ConnectionInfo{}, nil
		}
		return nil, err
	}
	var conns []models.ConnectionInfo
	if err := json.Unmarshal(data, &conns); err != nil {
		return nil, err
	}
	return conns, nil
}

func (c *ConnectionController) saveConnections(conns []models.ConnectionInfo) error {
	data, err := json.MarshalIndent(conns, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(c.filePath, data, 0644)
}
