package client

type Config struct {
	Servers map[string]serverCfg
}

type serverCfg  struct {
	Network string
	Address string
}