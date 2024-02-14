package models

// CheckPortModel represents the structure of the configuration file
type CheckPortModel struct {
	ConsulAddress string          `yaml:"consuladress"`
	Logfile       string          `yaml:"logfile"`
	Exporters     []ExporterModel `json:"exporters"`
}

// ExporterModel represents exporter information
type ExporterModel struct {
	Name       string `json:"name"`
	Port       int    `json:"port"`
	ExportType string `json:"exporttype"`
}
