package model

// repository_model
type Cluster struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	APIServer   string `json:"api_server"`
	Config      string `json:"config"`
	Region      string `json:"region"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

func (Cluster) TableName() string {
	return "clusters"
}
