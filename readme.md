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

接口路径：/api/v1/emotional   \
请求方式：POST
\
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
    "msg": "success"
}
```

2.登录接口：
\
接口路径：/api/v1/login
\
请求方法：POST
\
接口传参：
```json
{
    "username": "sll",
    "password": "1234"
}
```
返回数据：
```json
{
    "user_id": 1,
    "user_name": "sll",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJzbGwiLCJleHAiOjE2NzYwNjQzMjQsImlhdCI6MTY3NTk2MzcyNH0.q9_0_yzqQ1qQ0VJKhq"
}
```

3.情绪分析接口：
\
接口路径：/api/v1/emotional
\
请求方法：POST
\
接口传参：
```json
{
    "user_id": 1,
    "text": "今天是个好天气"
}
```
返回数据：
```json
{
    "msg": "good",
    "data": "。。。"  // 详细解释上述语句
}
```

4.情绪统计接口：
\
接口路径：/api/v1/emotional/count/:user_id
\
请求方法：GET
\
接口传参(参数包含在url里面)：/api/v1/emotional/count/1
\
返回数据：
```json
{
    "good_num": 1,     // 用户好心情数量
    "bad_num": 0,      // 用户坏心情数量
    "neutral_num": 0   // 用户中立心情数量
}
```


