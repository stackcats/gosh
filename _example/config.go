package config

import (
	"log"

	"github.com/Seven-Year/saas-ms-gordon/gosh"
)

// Qiniu 七牛配置
type Qiniu struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	Domain    string `json:"domain"`
	Zone      string `json:"zone"`
}

// SendCloud 邮件
type SendCloud struct {
	APIURL         string `json:"url"`
	APITemplateURL string `json:"template_url"`
	APIKey         string `json:"key"`
	APIUser        string `json:"user"`
	From           string `json:"from"`
}

// YunPian 短信
type YunPian struct {
	URL    map[string]string `json:"url"`
	APIKey string            `json:"key"`
}

// Config ...
type Config struct {
	Qn     *Qiniu     `consul:"qiniu"`
	Sc     *SendCloud `consul:"sendcloud"`
	Yp     *YunPian   `consul:"yunpian"`
	Port   string     `consul:"service/misc/port"`
	QnLive *Qiniu     `consul:"qiniu-live"`
	AesKey string     `consul:"service/misc/aes-key"`
}

var config *Config

func init() {
	config = &Config{}

	c := gosh.DefaultConfig()

	g, err := gosh.NewClient(c)
	if err != nil {
		log.Fatal(err)
	}

	err = g.Unmarshal(config)
	if err != nil {
		log.Fatal(err)
	}

}

// NewConfig ...
func NewConfig() *Config {
	return config
}
