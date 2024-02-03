## 💡 简介

基于gin-gorm框架实现的一个简单电商后端

## 🚀 功能

### 认证系统

+ 颁发token
+ 对令牌进行解系获取认证信息与用户信息

### 用户系统

+ 用户登录与注册,用到正则表达式规范化密码和用户名
+ 用户改名
+ 用户改密码
+ 用户搜索商品/查看商品列表
+ 用户选择商品加入购物查
+ 用户提交/取消订单

### 管理员系统

+ 管理员登录,获得授权
+ 管理员创建商品分类
+ 管理员创建商品
+ 管理员发布/删除商品
+ 管理员同样拥有用户系统功能

## 🌟 技术栈

<img src="https://gorm.io/gorm.svg" alt="GORM" style="float: left; zoom: 11000%;">

+ [gorm](https://gorm.io)

>Gorm 是一个用于 Go 语言的对象关系映射（ORM）库，它提供了一种方便的方式来操作数据库，使得开发者可以使用 Go 语言来进行数据库的 CRUD（创建、读取、更新、删除）操作，而无需编写原始的 SQL 语句

尝试用了gorm而不是sql语句

<img src="https://github.com/sgoware/gohu/raw/master/manifest/image/mysql.svg" alt="img" style="float: left; zoom: 150%;">

+ [mysql](https://www.mysql.com/)

>一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统关系型数据库管理系统之一，在 WEB 应用方面，MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一

用MySQL在存储相关的信息

<img src="https://static1.smartbear.co/swagger/media/assets/images/swagger_logo.svg" alt="Swagger Logo" style="zoom:9000%;" />

+ [swagger](https://swagger.io/)

>Swagger 是一个用于设计、构建、记录和使用 RESTful API 的工具集。它包括一组开源工具和技术，帮助开发者设计、构建和文档化 RESTful Web服务

用swagger帮助生成接口文档

+ **分页技术**

>分页,是一种将所有数据分段展示给用户的技术.用户每次看到的不是全部数据,而是其中的一部分,如果在其中没有找到自习自己想要的内容,用户可以通过制定页码或是翻页的方式转换可见内容,直到找到自己想要的内容为止.其实这和我们阅读书籍很类似.

+ [gin]([gin框架 · Go语言中文文档 (topgoer.com)](https://www.topgoer.com/gin框架/))

>- Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点

项目借助gin框架开发

# 基本介绍

#### 1.1项目介绍

> 本项目实现了一个电商项目的核心功能，内容包括：用户登录、用户注册、jwt鉴权、商品分类管理、商品管理、订单管理、购物车管理

#### 1.2使用说明

>MySQL版本 > v8.0
>
>golang版本 >= v1.21
>
>IDE推荐：Goland
>
>开发环境:wsl2-Ubuntu

#### 1.2下载项目

>go clone https://github.com/wmlhy2324/bluework.git

#### 2.1web项目

>```
># 安装依赖
>npm install
>```

>如果直接导入文件包Goland_version>1.21自动安装所需依赖

#### 2.2swagger自动化API文档

##### 2.2.1 安装 swagger

```
go install github.com/swaggo/swag/cmd/swag@latest
```

##### 2.2.2 生成API文档

```
swag init
```

>修改了swagger之后需要重新删除docs包,再在终端执行此命令生成新的接口文档

##### 2.2.3配置文件注意

![image-20240202143136812](C:\Users\17219\AppData\Roaming\Typora\typora-user-images\image-20240202143136812.png)

>在项目运行前,确保修改了配置,提前创建好数据库,否则报错

#### 3.1表结构说明

在运行项目后会自行创建表,分类表和用户表分别会插入两条测试数据

#### 3.2项目结构

.
├── api                                 //存放与api相关的代码
│   ├── cart
│   │   ├── controller.go    //控制器,控制的是请求调用的函数
│   │   └── types.go            //存放请求参数和响应参数的结构体
│   ├── category
│   │   ├── controller.go
│   │   └── types.go
│   ├── order
│   │   ├── controller.go
│   │   └── types.go
│   ├── product
│   │   ├── controller.go
│   │   └── types.go
│   ├── router.go //路由层,注册控制器
│   └── user
│       ├── controller.go
│       └── types.go
├── config              //存放配置文件
│   ├── config.go
│   └── config.yaml
├── docs       //使用swagger自动化api工具生成的接口文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── domain     //存放与业务领域相关的代码
│   ├── cart
│   │   ├── entity.go
│   │   ├── errors.go   //用于错误处理
│   │   ├── hooks.go  //存放钩子函数,比如在保存,更新数据之前执行
│   │   ├── repository.go  //仓库层,主要负责与数据进行交互,比如增删改查
│   │   └── service.go  //服务层, 服务层承载了业务逻辑的实现，负责处理业务规则、协调不同的仓库操作
│   ├── category
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── reposotory.go
│   │   └── service.go
│   ├── order
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── hooks.go
│   │   ├── order_item_repository.go
│   │   ├── order_repository.go
│   │   └── service.go
│   ├── product
│   │   ├── entity.go
│   │   ├── errors.go
│   │   ├── hooks.go
│   │   ├── repository.go
│   │   └── service.go
│   └── user
│       ├── entity.go
│       ├── hooks.go
│       ├── repository.go
│       ├── servers.go
│       ├── user.go
│       └── validation.go
├── go.mod                        //管理依赖关系和版本信息
├── go.sum
├── main.go //主包
└── utils                             //工具层,用于对各种特殊情况下处理
    ├── api_helper //
    │   ├── error_handler.go      //错误处理返回
    │   ├── query_helper.go      //从上下文获取用户id
    │   └── types.go //请求参数的结构体
    ├── database_handler //处理数据库连接的
    │   └── mysql_handler.go
    ├── graceful                   //优雅退出项目
    │   └── showdown.go
    ├── hash                         //用于对用户的密码加盐加密
    │   └── hash.go
    ├── jwt                             //用于jwt鉴权
    │   └── jwt_helper.go
    ├── middleware             //中间件的使用,主要是给用户和管理员授权
    │   └── auth_middleware.go
    └── pagination               //分页工具的使用
        └── pages.go

#### 4.1联系信息

如果你有任何问题或建议，请随时通过我的电子邮件1721966546@qq.com与我联系。”
