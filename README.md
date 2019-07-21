# God

#### 介绍
go God 

启动： go run main.go -conf ./config/${fileName}.json

eg: go run main.go -conf ./config/local.json

#### 软件架构
本想目采用 go gin 框架搭建，关于gin的具体使用请参照：
 https://github.com/gin-gonic/gin/README.md


#### 安装教程

1. 安装 go sdk 1.8+
2. 安装 govendor
    `go get -u -v github.com/kardianos/govendor`
3. xxxx

#### 使用说明

1. 将外部依赖全部加载到 vendor
    `govendor add +e`
2. 指定加载单个外部依赖到 vendor
    `govendor add xxxxx`
3. 列出当前项目的包引用情况 (可以将 e 的包都add 到vendor中)
    `govendor list -v fmt`

4. 安装docker

5. docker 安装 mysql: 5.7
    ` docker pull mysql:5.7`
6. 启动 docker 中的mysql:5.7
    `docker run --name mysql -p 127.0.0.1:3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d ${镜像Id}`
7.  进入容器
    `docker exec -ti ${容器Id} bash`
    或
    `winpty docker exec -ti ${容器Id} bash`       
8.  进入mysql
    ` mysql -h${Ip} -u${账户} -p${密码}`  
9. 安装redis
    `docker pull  redis:4.0`
10. 启动 redis
    ` docker run --name redis -p 6379:6379 -d 598a6f110d01 redis-server --appendonly yes --requirepass "123456"` 
11. 进入容器 (用 redis-cli 登入)
    `winpty  docker exec -ti b617859f4065 redis-cli`    
     
    
