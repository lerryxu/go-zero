package common

import "fmt"

const (
	DISABLED = 0 // 禁用
	ENABLED  = 1 // 启用
)

type MysqlConfig struct {
	Host          string `json:",default=127.0.0.1"`
	Port          int    `json:",default=3306"`
	User          string `json:",optional"`
	Pass          string `json:",optional"`
	DataBase      string `json:""`
	SlowThreshold int    `json:",default=1000"`
}

func MysqlDsn(c MysqlConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.DataBase)
}
