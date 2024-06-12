package models

type App struct {
	ServerName       string `json:"server_name"`
	PortRun          string `json:"port_run"`
	Secret           string `json:"secret"`
	TokenLifeSeconds int    `json:"token_life_seconds"`
}
