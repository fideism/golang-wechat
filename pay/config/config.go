package config

//Config config for pay
type Config struct {
	Sandbox bool   `json:"sandbox"`
	AppID   string `json:"app_id"`
	MchID   string `json:"mch_id"`
	Key     string `json:"key"`
}

// Cert 证书配置优先使用 content 如没有内容则读取path路径文件
type Cert struct {
	Path    string `json:"path"`
	Content []byte `json:"content"`
}
