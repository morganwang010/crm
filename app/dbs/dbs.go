package dbs

import (
	"github.com/sirupsen/logrus"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)
var logger = logrus.New()
func Init() {
    // 定义数据库连接字符串。
    dsn := "host=10.122.66.82 port=5432 user=a_appconnect password=abcd-1234 dbname=drmp"

    // 创建一个`DB`实例。
    db, err := gorm.Open(postgres.Open(dsn))
    if err != nil {
        panic(err)
    }
		// Set up connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)

    // // 使用`DB`实例执行SQL查询。
    // var users []User
    // db.Find(&users)

    // // 使用`DB`实例更新数据库。
    // for _, user := range users {
    //     user.Name = "John Doe"
    //     db.Save(&user)
    // }

    // // 关闭`DB`实例。
    // db.Close()
}

// type User struct {
//     gorm.Model
//     Name string `gorm:"column:name"`
// }