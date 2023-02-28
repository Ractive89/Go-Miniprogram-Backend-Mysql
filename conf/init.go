package conf

import (
	"dayang/models"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	UserName     string
	UserPassWord string
	UserDatabase string
	DbHost       string
	DbPort       string
	AppID        string
	AppSecret    string
	ProjectName  string
	HttpPort     string
	SSLPem       string
	SSLKey       string
	HostName     string
)

func Init() *gorm.DB {
	// 读取配置文件
	err := godotenv.Load("app.env")

	if err != nil {
		log.Fatal("Load Env File Error.", err)
		return nil
	}

	loadDatabase()
	loadAPP()
	loadSSL()

	// 拼接数据库路径
	path := strings.Join([]string{UserName, ":", UserPassWord, "@tcp(", DbHost, ":", DbPort, ")/", UserDatabase, "?charset=utf8mb4&parseTime=true"}, "")

	// 加载Mysql
	dayangDB := models.Database(path)

	return dayangDB
}

func loadDatabase() {
	UserName = os.Getenv("USER_NAME")
	UserPassWord = os.Getenv("USER_PASSWORD")
	UserDatabase = os.Getenv("USER_DATABASE")
	DbHost = os.Getenv("DATABASE_HOST")
	DbPort = os.Getenv("DATABASE_PORT")
}

func loadAPP() {
	ProjectName = os.Getenv("PROJECT_NAME")
	AppID = os.Getenv("APP_ID")
	AppSecret = os.Getenv("APP_SECRET")
	HttpPort = os.Getenv("ROUTE_PORT")
}

func loadSSL() {
	SSLPem = os.Getenv("SSLPEM")
	SSLKey = os.Getenv("SSLKEY")
	HostName = os.Getenv("HOSTNAME")
}
