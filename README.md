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



