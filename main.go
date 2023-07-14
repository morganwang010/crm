package main

import (
    "github.com/gin-gonic/gin"
    "gogin/utils"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "fmt"
)
//type User struct {
    // ...
type User struct {
    EmployeeNumber string `gorm:"column:employee_number" json:"employeeNumber"`
    RoleName       string `gorm:"column:role_name" json:"roleName"`
    CreateTime     string `gorm:"column:create_time" json:"createTime"`
    UpdateTime     string `gorm:"column:update_time" json:"updateTime"`
    Password       string `gorm:"column:password" json:"password"`
    IsActive       string `gorm:"column:is_active" json:"isActive"`
    Type           string `gorm:"column:type" json:"type"`
    Mail           string `gorm:"column:mail" json:"mail"`
    OrgId          int32  `gorm:"column:org_id" json:"orgId"`
    OrgName        string `gorm:"column:org_name" json:"orgName"`
    Id             int32  `gorm:"column:id;primaryKey" json:"id"`
}

//}
// @Summary 获取用户信息
// @Description 根据用户ID获取用户信息
// @Tags Users
// @Accept json
// @Param id path int true "用户ID"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Router /users/{id} [get]
func getUserByID(c *gin.Context) {
    // 处理获取用户信息的逻辑
    utils.UtilityFunc() // 调用来自 utils 包的 UtilityFunc 函数
       c.JSON(200, gin.H{
            "message": "Hello, World,catch you againg!!!!!!!!!!",
        })

}

 



func main() {
    r := gin.Default()
    // db := dbs.Init()
    // 定义路由和处理函数
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })
    r.GET("/aa",getUserByID)
    // dbs.Init()
    r.POST("/users", func(c *gin.Context) {
        // 处理创建用户的逻辑
        // c.JSON() 构建 JSON 响应
    })
    // 查询名为Tom的用户。
    dsn := "host=10.122.66.82 port=5432 user=a_appconnect password=abcd-1234 dbname=drmp"

    // 创建一个`DB`实例。
    db, err := gorm.Open(postgres.Open(dsn))
    if err != nil {
        panic(err)
    }
    user := User{EmployeeNumber: "wangjw30w"}
    // db.Where(&user).First(&user)
    db = db.Table("dr.user").Select(`*`)
    db.Find(&user)

    // 将用户信息输出到控制台。
    fmt.Println(user)
    // 启动服务器
    r.Run()
}

