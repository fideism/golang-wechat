package config

//Config config for pay
type Config struct {
	Sandbox bool   `json:"sandbox"`
	AppID   string `json:"app_id"`
	MchID   string `json:"mch_id"`
	Key     string `json:"key"`
}

type Cert struct {
	Path string `json:"path"`
}
