package models

// ConnectionInfo represents saved connection information
type ConnectionInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslMode"`
}

// QueryResult represents the result of a database query
type QueryResult struct {
	Columns []string        `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
	Error   string          `json:"error,omitempty"`
}

// TableInfo represents basic table metadata
type TableInfo struct {
	Name    string `json:"name"`
	Schema  string `json:"schema"`
	Type    string `json:"type"` // TABLE, VIEW, etc.
	Comment string `json:"comment,omitempty"`
}
