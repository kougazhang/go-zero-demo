## 用户模块功能及 Api
### 默认返回值
```json
{
  "ok": true,
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
+ Api: /users/id/:id
+ Method: DELETE

### 4. 修改用户
+ Api: /users/id/:id
+ Method: PUT
```json
{
  "username": "",
  "password": ""
}
```

### 5. 查询用户
+ Api: /users/id/:id
+ Method: GET
+ Response:
```json
{
  "id": 0,
  "username": "",
  "password": ""
  }
```

### 6. 查询所有用户
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
