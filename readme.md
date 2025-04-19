本文档是我的苏州大学毕业设计项目的后端说明，后端使用go + gin + gorm + mysql + redis完成，其中mysql和redis配置如下：

```
mysql:
  host: "127.0.0.1"
  port: 3307
  user: "rootsll"
  password: "root1234"
  dbname: "graduation"
  max_open_conns: 200
  max_idle_conns: 50
```

```
redis:
  host: "127.0.0.1"
  port: 16379
  password: ""
  db: 0
  pool_size: 100
```

mysql和redis建议使用docker起一个，这样方便快捷，在该文件的终端执行以下命令：
```bash
docker-compose up -d
```
项目监听在8080端口，运行如下指令即可启动项目：
```bash
go run main.go
```


目前后端包含四个接口，分别是注册，登录，情绪分析，情绪统计以及打卡接口，接口具体说明如下。
<br />

## 1.注册接口：
\
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
<br />

## 2.登录接口：
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
<br />

## 3.情绪分析接口：
\
接口路径：/api/v1/signup
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
    "data": {
        "sentiment": "这句话表达了正面情感",
        "sentiment_type": 1,
        "text": "今天心情很好"
    },
    "msg": "success"
}
```
<br />

## 4.情绪统计接口：
\
接口路径：/api/v1/emotional/count/:user_id
\
请求方法：GET
\
接口传参(参数包含在url里面)：/api/v1/emotional/count/1234
\
返回数据：
```json
{
    "data": {
        "good_num": 2,
        "bad_num": 0,
        "neutral_num": 0
    },
    "msg": "success"
}
```

## 5.打卡接口：
\
接口路径：/api/v1/checkin/:user_id
\
请求方法：POST
\
接口传参(参数包含在url里面)：/api/v1/checkin/1234
\
返回数据：
```json
{
    "data": "打卡成功",
    "msg": "success"
}
```
