## openshift-basic-identity-provider

* 可作为openshift basic remote identity provider为openshift提供用户认证
* 目前使用splite作为数据存储，之后会换成mysql或其他数据库

## API

以下是目前支持的api

![](doc/img/api.png)

详细接口信息可以使用项目内`swagger.yaml`文件在**swagger hub**中查看

### 数据表结构

```sqlite
CREATE TABLE IF NOT EXISTS user(
	id integer not null primary key, 
	username text not null unique,
	password text not null,
	email text,
	name text
	)
```

## how to start

### 容器运行

* git clone 本项目，需要本地环境有``docker`及`go`

* `docker build -t imagename:tag .`
* `docker run -d -p port:8080 imagename:tag`

### openshift运行

#### 使用storage class

* 编辑项目下`openshift/template/tmpl-use-sc.json`中的参数
* `oc new-app -f openshift/template/tmpl-use-sc.json`

#### 使用已有pv运行

* `oc new-app -f openshift/template/tmpl.json`

### 本地测试

替换url中的`yourdomain`

* list user

  `curl -X GET "http://yourdomain/openshift-basic-identity-provider/1.0.0/users" -H "accept: application/json"`

* create user

  `curl -X POST "http://yourdomain/openshift-basic-identity-provider/1.0.0/user" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"password\": \"john\", \"username\": \"john\"}"`

* update user

  `curl -X PUT "https://yourdomain/openshift-basic-identity-provider/1.0.0/user/john" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"password\": \"john\", \"name\": \"john\", \"email\": \"somebody@gmail.com\", \"username\": \"john\"}"`

* delete user

  `curl -X DELETE "http://yourdomain/openshift-basic-identity-provider/1.0.0/user/john" -H "accept: application/json"`

* validate login

  `curl -X POST "http://yourdomain/openshift-basic-identity-provider/1.0.0/auth/token" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"password\": \"john\", \"username\": \"john\"}"`

### to do list

- [ ] 错误处理及返回
- [ ] 日志打印
- [ ] 数据库字段加密

