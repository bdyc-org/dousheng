# dousheng
字节跳动第三届青训营后端专场9003队抖音项目

## 灵感来自
[cloudwego/kitex-examples/bizdemo/easy_note](https://github.com/cloudwego/kitex-examples/tree/main/bizdemo/easy_note)

## 项目介绍
[接口说明文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)

### 项目技术栈
- 语言：go
- 数据库：MySQL
- RPC框架：Kitex
- ORM框架：GORM
- HTTP框架：Gin

### 项目模块介绍
| 服务名称   |  模块介绍    | 技术框架   | 传输协议 | 注册中心 |
| --------- | ------------ | ---------- | ------- | ------- |
| api       | API服务       | kitex/gin | http    | etcd    |
| user      | 用户数据管理  | kitex/gorm | thrift | etcd     |
| video     | 视频数据管理  | kitex/gorm | thrift | etcd     |
| favorite  | 点赞数据管理  | kitex/gorm | thrift | etcd     |
| comment   | 评论数据管理  | kitex/gorm | thrift | etcd     |
| relation  | 关注数据管理  | kitex/gorm | thrift | etcd     |

### 项目代码目录结构介绍

```
├─idl（接口定义文件）
├─kitex_gen（Kitex自动生成的代码）
├─cmd
│  ├─api（api服务的业务代码）
│  ├─user（user服务的业务代码）
│  ├─video（video服务的业务代码）
│  ├─favorite（favorite服务的业务代码）
│  ├─comment（comment服务的业务代码）
│  ├─relation（relation服务的业务代码）
├─public(静态文件存放位置)
│  ├─avatar（用户头像存放位置）
│  ├─background（背景图片存放位置）
├─pkg
│  ├─constants（常量）
│  ├─errno（c错误码和错误信息）
│  ├─middleware（中间件）
├─README.md（项目说明文件）
```

## 下载与使用说明

### 运行环境
- linux(Ubuntu)
- go version: 1.18
- mysql version: 8.0
- etcd version: 3.5.4

### 下载
```shell
git clone https://github.com/bdyc-org/dousheng.git
```

### 使用

#### 1. 准备好运行环境，
修改 `pkg/constants/constant.go` 中的 `MySQLDefaultDSN` ，并启动 `etcd`

#### 2. 运行user服务
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

#### 3. 运行video服务
```shell
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```

#### 4. 运行favorite服务
```shell
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh
```

#### 5. 运行comment服务
```shell
cd cmd/comment
sh build.sh
sh output/bootstrap.sh
```

#### 6. 运行relation服务
```shell
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```

#### 7. 运行api服务
```shell
go run cmd/api/main.go
```
