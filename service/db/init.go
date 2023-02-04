package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init :init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	//if err = DB.Use(gormopentracing.New()); err != nil {
	//	panic(err)
	//}
}

const (
	NoteTableName           = "note"
	UserTableName           = "user"
	VideoTableName          = "video"
	SecretKey               = "secret key"
	IdentityKey             = "id"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	ApiServiceName          = "demoapi"
	NoteServiceName         = "demonote"
	UserServiceName         = "demouser"
	MySQLDefaultDSN         = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress             = "127.0.0.1:2379"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
)
