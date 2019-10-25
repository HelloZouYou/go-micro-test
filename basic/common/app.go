package common

import "strconv"

// AppCfg common config
type AppCfg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// Addr Addr
func (a *AppCfg) Addr() string {
	return a.Address + ":" + strconv.Itoa(a.Port)
}
