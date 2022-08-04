package gorm

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/caicaispace/gohelper/setting"

	"gorm.io/driver/mysql"
	orm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GromStruct struct {
	dbs map[string]*orm.DB
}

var (
	service *GromStruct
	once    sync.Once
)

func GetInstance() *GromStruct {
	once.Do(func() {
		service = New()
	})
	return service
}

func New() *GromStruct {
	gs := &GromStruct{}
	gs.dbs = make(map[string]*orm.DB)
	return gs
}

func (gs *GromStruct) AddConnWithConfig(conf *setting.DBSetting, connName string) *GromStruct {
	gs.AddConnWithDns(conf.ToDnsString(), connName)
	return gs
}

func (gs *GromStruct) AddConnWithDns(dns string, connName string) *GromStruct {
	db, err := orm.Open(mysql.Open(dns), getGormConfig())
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	gs.AddConn(db, connName)
	return gs
}

func (gs *GromStruct) AddConn(db *orm.DB, connName string) *GromStruct {
	gs.dbs[getConnName(connName)] = db
	return gs
}

func (gs *GromStruct) GetDB(connName string) *orm.DB {
	return gs.GetConn(connName)
}

func (gs *GromStruct) GetConn(connName string) *orm.DB {
	return gs.dbs[getConnName(connName)]
}

func getConnName(connName string) string {
	if connName == "" {
		connName = "default"
	}
	return connName
}

func getGormConfig() *orm.Config {
	c := &orm.Config{}
	if setting.Server.Env == "dev" {
		c.Logger = logger.Default.LogMode(logger.Info)
		// c.Logger slowLogger()
	}
	return c
}

func slowLogger() logger.Interface {
	return logger.New(
		// stdout
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,   // slow sql threshold
			LogLevel:      logger.Silent, // log level
			Colorful:      false,         // disable color outputs
		},
	)
}

// func NewWithDSN(dns string, connName string) {
// 	dialector := mysql.New(mysql.Config{
// 		DSN:                       dns,   // data source name
// 		DefaultStringSize:         256,   // default size for string fields
// 		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
// 		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
// 		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
// 		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
// 	})
// 	db, err := orm.Open(dialector, getGormConfig())
// }
