# User service 
慕课网 ap 老师 课程购买地址：http://www.imooc.com/t/6512963

当前服务 名称为 User 类型 service 

创建初始化模版请使用

```
sudo docker run --rm -v $(pwd): $(pwd) -w  $(pwd) -e ICODE=xxxxxx cap1573/cap-micro new git.imooc.com/cap1573/user
```
以上命令中 "xxxxxx" 为个人购买的 icode 码，请勿多人使用（会被慕课网检测失效）。

## 快速开始

- [配置信息](#配置信息)
- [使用](#使用)

## 配置信息

- 服务名称: go.micro.service.user
- 类型: service
- 简称: user

 

## 使用
根据 proto 自动生成
```
make proto
```

编译
```
make proto
```

构建镜像
```
make docker
```