### 前提需要安装好go
### 设置环境变量 
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go env 
```
project初始化，
go mod init gouser   
会在文件夹下生成一个go.mod的文件，
然后安装所需要的依赖包即可，可以先装gin的框架
go get -u github.com/gin-gonic/gin
创建swag的docker服务化
安装swag
go get -u github.com/swaggo/swag/cmd/swag
此时只是下载了swag的包，并没有实际安装，需要编译一下放到系统的可执行路径里
进入到包安装目录
```bash
cd /data/golang/pkg/mod/github.com/swaggo/swag@v1.16.1/cmd/swag
go build main.go
ln -s `pwd`/main /bin/swag
##再回到项目所在的文件夹，执行
swag init
#此时会将项目下的文件里的api接口都自动生成文档

docker run -p 8080:8080 -e SWAGGER_JSON=/app/docs/swagger.json -v $(pwd)/docs:/app/docs swaggerapi/swagger-ui:v3.52.0

读取数据层使用gorm及sqlgen，接口token验证使用PASETO (Platform-Agnostic SEcurity TOkens)，配置管理使用viper
安装软件包：
go get -u github.com/jinzhu/gorm
go get github.com/spf13/viper
go install github.com/anqiansong/sqlgen@latest
使用sqlgen生成models文件
创建models文件夹，在下面创建一个sql文件，直接在命令行进入到models文件下，执行sqlgen即可。