本文档是我的苏州大学毕业设计项目的后端说明，后端使用go + gin + gorm + mysql完成，其中mysql配置如下：

```
mysql:
  host: "127.0.0.1"
  port: 3306
  user: "root"
  password: "root1234"
  dbname: "graduation"
  max_open_conns: 200
  max_idle_conns: 50
```

mysql建议使用docker起一个，这样方便快捷，在该文件的终端执行以下命令：
```bash
docker-compose up -d
```
项目监听在8080端口，运行如下指令即可启动项目：
```bash
go run main.go
```


目前后端包含四个接口，分别是注册，登录，情绪分析和情绪统计接口，接口具体说明如下。

1.注册接口：

接口路径：/api/v1/signup

请求方式：POST

接口传参：
```json
{
    "username": "sll",
    "password": "1234",
    "re_password": "1234"
}
```
返回数据：
```json
{
    
}
```




