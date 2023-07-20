# gstore

使用 [gstore](https://github.com/pkumod/gStore) 的 java api 进行数据库的操作

**目录结构**:

```
│  .gitignore
│  mvnw
│  mvnw.cmd
│  pom.xml
│  README.md
│
├─src
│  ├─main
│     ├─java
│     │  └─com
│     │      └─guochenxu
│     │          └─gstore
│     │              │  GstoreApplication.java
│     │              │
│     │              ├─config
│     │              │      GstoreConfiguration.java   // 数据库配置信息
│     │              │      WebMvcConfig.java   
│     │              │
│     │              ├─controller
│     │              │      GstoreController.java   // 接口
│     │              │
│     │              ├─entity
│     │              │      GstoreResult.java   // gstore查询结果
│     │              │
│     │              ├─jgsc
│     │              │      GstoreConnector.java
│     │              │
│     │              ├─service
│     │              │  │  GstoreService.java   // 服务类
│     │              │  │
│     │              │  └─impl
│     │              │          GstoreServiceImpl.java
│     │              │
│     │              └─vo
│     │                      R.java
│     │
│     └─resources
│         │  application.properties
│         │  application.yml
│         │
│         ├─static
│         └─templates
```

## 启动

**环境配置**: java 1.8

**数据库配置**: GstoreConfiuration中的配置信息

```shell
mvn package
java -jar ./target/gstore-0.0.1-SNAPSHOT.jar
```

## 程序执行流程

1. 接收请求, 加载数据库, 若加载失败直接返回null
2. 在数据库中进行查询
    1. 若查询顺利返回格式为解析后的实体 (bindings部分考虑到参数不固定, 所以未解析)
    2. 若查询超时 (设定定时器为30秒), 则返回null

## 返回状态码

| 状态码  | 含义           |
|------|--------------|
| 200  | 查询正常         |
| 500  | 数据库连接错误或查询出错 |
| 1001 | 查询结果为空       |

接口文档可以查看[这里](https://apifox.com/apidoc/shared-45a359ee-2428-4504-b9b7-3d20f3f71a20)
