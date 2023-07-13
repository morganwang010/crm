package main

import "github.com/gin-gonic/gin"
//type User struct {
    // ...
    type User struct {
    ID       int
    Username string
    Email    string
}
type ErrorResponse struct {
    Error string

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
       c.JSON(200, gin.H{
            "message": "Hello, World,catch you againg!!!!!!!!!!",
        })

}
func main() {
    r := gin.Default()

    // 定义路由和处理函数
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World!",
        })
    })
    r.GET("/aa",getUserByID)

    r.POST("/users", func(c *gin.Context) {
        // 处理创建用户的逻辑
        // c.JSON() 构建 JSON 响应
    })

    // 启动服务器
    r.Run()
}

