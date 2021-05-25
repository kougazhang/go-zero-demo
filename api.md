## 用户模块功能及 Api
### 默认正常返回值
```json
{
  "Ok": true
}
```
### 默认异常返回值
```json
{
  "Ok": false,
  "error": ""
}
```

### 1. 用户登录
+ Api: /login
+ Method: POST
+ RequestBody:
```json
{
  "username": "",
  "password": ""
}
```

### 2. 添加用户
+ Api: /users
+ Method: POST
+ RequestBody:
```json
{
  "username": "",
  "password": ""
}
```

### 3. 删除用户
+ Api: /users/:id
+ Method: GET

### 4. 修改用户
+ Api: /users/:id
+ Method: POST
```json
{
  "id": 0,
  "username": "",
  "password": ""
}
```

### 5. 查询用户
+ Api: /users
+ Method: GET
+ Response:
```json
[
  {
  "id": 0,
  "username": "",
  "password": ""
  }
]
```