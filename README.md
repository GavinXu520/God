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

一、安装 govendor

1. 将外部依赖全部加载到 vendor
    `govendor add +e`
2. 指定加载单个外部依赖到 vendor
    `govendor add xxxxx`
3. 列出当前项目的包引用情况 (可以将 e 的包都add 到vendor中)
    `govendor list -v fmt`
4. 如果使用govendor 下载新包，请用
    `govendor fetch ${包名}`
    如:
    `govendor fetch github.com/go-sql-driver/mysql`   
5. 如果是先修改了 vendor.json 的话，需要用下面命令同步包到vendor目录中
    `govendor sync`    
     

二、docker安装 mysql

1. 安装docker (怎么安装自己查)

2. docker 拉取 mysql: 5.7 镜像
    `docker pull mysql:5.7`
3. 查看镜像ID
    `docker images`    
4. docker启动mysql:5.7
    `docker run --name mysql -p 127.0.0.1:3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d ${镜像Id}`
5. 查看由对应镜像启动的容器ID
    `docker container ls`
6. 进入容器
    `docker exec -ti ${容器Id} bash`
    或者
    `winpty docker exec -ti ${容器Id} bash`       
7. 进入mysql
    `mysql -h${Ip} -u${账户} -p${密码}` 
    如果不行，则
    `mysql -uroot -p`
    然后等待输入密码提示出现后，再输入密码

三、docker安装 redis
 
1. 安装redis
    `docker pull  redis:4.0`
2. 查看镜像ID
    `docker images`
3. 启动 redis
    `docker run --name redis -p 6379:6379 -d ${镜像ID} redis-server --appendonly yes --requirepass "123456"` 
4. 查看由对应镜像启动的容器ID
    `docker container ls`
5. 进入容器 (用 redis-cli 登入)
    `docker exec -ti ${容器ID} redis-cli`
    或者
    `winpty  docker exec -ti ${容器ID} redis-cli`  
    
四、注意： 在连接 mysql后，执行建表语句时，报下面错误时：
    `ERROR 1067 (42000)/ERROR 1292 (22007) Zero Date is not Accepted by Timestamp`   
    请先直接执行：
    `set global  sql_mode = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';`
    然后在执行建表语句即可。   
     
    
