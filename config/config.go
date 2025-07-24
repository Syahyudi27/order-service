package config

import (
	"order-service/common/util"
	"os"

	"github.com/sirupsen/logrus"
	//_ "github.com/spf13/viper/remote"
)

var Config AppConfig

type AppConfig struct {
	Port                   int             `json:"port"`
	AppName                string          `json:"appName"`
	AppEnv                 string          `json:"appEnv"`
	SignatureKey           string          `json:"signatureKey"`
	Database               Database        `json:"database"`
	RateLimiterMaxRequest  int             `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond  int             `json:"rateLimiterTimeSecond"`
	InternalService        InternalService `json:"internalService"`
	GCSType                string          `json:"gcsType"`
	GCSProjectID           string          `json:"gcsProjectID"`
	GCSPrivateKeyID        string          `json:"gcsPrivateKeyID"`
	GCSPrivateKey          string          `json:"gcsPrivateKey"`
	GCSClientEmail         string          `json:"gcsClientEmail"`
	GCSClientID            string          `json:"gcsClientID"`
	GCSAuthURI             string          `json:"gcsAuthURI"`
	GCSTokenURI            string          `json:"gcsTokenURI"`
	GCSAuthProviderCertURL string          `json:"gcsAuthProviderX509CertURL"`
	GCSClientCertURL       string          `json:"gcsClientX509CertURL"`
	GCSUniverseDomain      string          `json:"gcsUniverseDomain"`
	GCSBucketName          string          `json:"gcsBucketName"`
	Kafka                  Kafka           `json:"kafka"`
	Midtrans               Midtrans        `json:"midtrans"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnection     int    `json:"maxOpenConnection"`
	MaxLifetimeConnection int    `json:"maxLifetimeConnection"`
	MaxIdleConnection     int    `json:"maxIdleConnection"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

type InternalService struct {
	User    User    `json:"user"`
	Field   Field   `json:"field"`
	Payment Payment `json:"payment"`
}

type User struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

type Field struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

type Payment struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

type Kafka struct {
	Brokers               []string `json:"brokers"`
	TimouInMs             int      `json:"timouInMs"`
	MaxRetry              int      `json:"maxRetry"`
	Topics                []string `json:"topic"`
	GroupID               string   `json:"groupID"`
	MaxWaitTime           int      `json:"maxWaitTime"`
	MaxProcessingTimeInMs int      `json:"maxProcessingTimeInMs"`
	BackoffTimeInMs       int      `json:"backoffTimeInMs"`

}

type Midtrans struct {
	ServerKey    string `json:"serverKey"`
	ClientKey    string `json:"clientKey"`
	Isproduction bool   `json:"isProduction"`
}

func Init() {
	err := util.BindFromJSON(&Config, "config.json", ".")
	if err != nil {
		logrus.Infof("failed to bind from config: %v", err)
		err = util.BindFromConsul(Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}
	}
}
