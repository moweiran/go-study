package repository

import (
	"app/utils"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 55000
	user     = "postgres"
	password = "postgrespw"
	dbname   = "postgres"
)

type pgObj struct {
	client *gorm.DB
}

func newPgClients() *pgObj {
	return &pgObj{
		client: getDB(),
	}
}

func (d *pgObj) GetClient() *gorm.DB {
	return d.client
}

func getDB() *gorm.DB {
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Println("PgSQL URL: " + dns)

	var slowLogger logger.Interface
	// var writers = make([]io.Writer, 0)
	var colorful bool
	var logLeve = logger.Info
	// writers = append(writers, graylog.GetGraylogClient())
	slowLogger = logger.New(
		// log.New(io.MultiWriter(writers...), "\r\n", log.LstdFlags),
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//设定慢查询时间阈值为2s
			SlowThreshold: 2 * time.Second,
			LogLevel:      logLeve,
			Colorful:      colorful,
		},
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: slowLogger,
		// NamingStrategy: schema.NamingStrategy{
		// 	SingularTable: true,
		// },
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now()
		},
	})

	utils.CheckError(err)

	return db
}
