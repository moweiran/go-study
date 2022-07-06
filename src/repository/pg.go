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

	students "app/domain/students/models"
)

const (
	host     = "localhost"
	port     = 55000
	user     = "postgres"
	password = "postgres"
	dbname   = "study"
)

type MigrateTable interface {
	TableName() string
}

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

func (d *pgObj) AutoMigrate() {
	if d.client.Migrator().HasTable(&students.Students{}) {
		log.Printf("======> DB Auto Migrate: Migrate Table %s\n", (&students.Students{}).TableName())
	}

	if err := d.client.AutoMigrate(&students.Students{}); err != nil {
		log.Printf("======> DB Auto Migrate: Create table %s failed\n", (&students.Students{}).TableName())
		panic(err.Error())
	}
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
