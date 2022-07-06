package database

const (
	host     = "localhost"
	port     = 55000
	user     = "postgres"
	password = "postgrespw"
	dbname   = "postgres"
)

func InitGorm() {
	// https://github.com/jackc/pgx
	// dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// log.Println("PgSQL URL: " + dns)
	// var slowLogger logger.Interface
	// // var writers = make([]io.Writer, 0)
	// var colorful bool
	// var logLeve = logger.Info
	// // writers = append(writers, graylog.GetGraylogClient())
	// slowLogger = logger.New(
	// 	// log.New(io.MultiWriter(writers...), "\r\n", log.LstdFlags),
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		//设定慢查询时间阈值为2s
	// 		SlowThreshold: 2 * time.Second,
	// 		LogLevel:      logLeve,
	// 		Colorful:      colorful,
	// 	},
	// )

	// db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
	// 	Logger: slowLogger,
	// 	// NamingStrategy: schema.NamingStrategy{
	// 	// 	SingularTable: true,
	// 	// },
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// 	NowFunc: func() time.Time {
	// 		return time.Now()
	// 	},
	// })

	// utils.CheckError(err)

	// students := models.Students{}

	// db.AutoMigrate(&models.Students{})

	// db.Create(&models.Students{Name: "a", Roll: "2"})

	// result := db.Debug().First(&students)

	// studentsArr := []models.Students{}

	// db.Debug().Find(&studentsArr, []int{1, 2, 3})

	// errors.Is(result.Error, gorm.ErrRecordNotFound)

}
