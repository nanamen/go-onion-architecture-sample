package conf

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //mysql
	"github.com/jinzhu/gorm"
)

// NewDBConnection 新規データベースコネクションを取得します.
func NewDBConnection() *gorm.DB {
	return getMysqlConn()
}

func getMysqlConn() *gorm.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		Current.Database.User,
		Current.Database.Password,
		Current.Database.Host,
		Current.Database.Port,
		Current.Database.Database,
	)

	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)
	conn.DB().SetMaxIdleConns(10)
	conn.DB().SetMaxOpenConns(20)

	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}
