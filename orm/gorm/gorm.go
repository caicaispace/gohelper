package gorm

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/caicaispace/gohelper/setting"
	"go.opentelemetry.io/otel"
	"gorm.io/plugin/opentelemetry/tracing"

	"gorm.io/driver/mysql"
	orm "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GromStruct struct {
	dbs    map[string]*orm.DB
	config *orm.Config
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
	db, err := orm.Open(mysql.Open(dns), gs.config)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	gs.AddConn(db, connName)
	return gs
}

func (gs *GromStruct) AddConn(db *orm.DB, connName string) *GromStruct {
	gs.dbs[gs.getConnName(connName)] = db
	return gs
}

func (gs *GromStruct) AddLogger(logger logger.Interface) *GromStruct {
	gs.config.Logger = logger
	return gs
}

func (gs *GromStruct) GetDB(connName string) *orm.DB {
	return gs.GetConn(connName)
}

func (gs *GromStruct) GetConn(connName string) *orm.DB {
	return gs.dbs[gs.getConnName(connName)]
}

// https://github.com/go-gorm/opentelemetry
// https://github.com/go-gorm/opentelemetry/tree/master/examples/demo
func (gs *GromStruct) UseTracer() {
	for _, db := range gs.dbs {
		func(db *orm.DB) {
			if err := db.Use(tracing.NewPlugin()); err != nil {
				panic(err)
			}
			ctx := context.Background()
			tracer := otel.Tracer("gorm")
			ctx, span := tracer.Start(ctx, "gorm")
			defer span.End()
		}(db)
	}
}

func (gs *GromStruct) UseSlowLogger() {
	gs.config.Logger = logger.New(
		// stdout
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,   // slow sql threshold
			LogLevel:      logger.Silent, // log level
			Colorful:      false,         // disable color outputs
		},
	)
}

func (gs *GromStruct) getConnName(connName string) string {
	if connName == "" {
		connName = "default"
	}
	return connName
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
