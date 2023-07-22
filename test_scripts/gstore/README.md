# 测试脚本

**目录结构**:

```
│  go.mod
│  main.go   // 调用Test()和Train()函数运行脚本
│  README.md
│
├─model
│      model.go   // 数据模型
│
├─service
│      do_http.go   // 发送http请求
|      get_values.go   // 解析返回值
│      test.go   // 测试生成的sparql是否正确
│      train.go   // 测试根据train.txt中的sparql能否查到相同的答案
│
└─txt
        generated_predictions.jsonl // 预测生成的sparql
        generated_predictions.txt // 预测生成的sparql
        myans.txt  // 根据train中sparql查询的结果
        my_pred.txt   // 根据预测生成的sparql查询到的结果
        sparql_not_test.txt  // 生成的sparql中为查询到结果的语句
        sparql_not_train.txt  // train中未查询到结果的语句
        train.txt
```

## 启动

**环境配置**: go 1.19.4

**文件路径设置**: `train.go` 和 `test.go` 中的文件路径需要进行配置

**http 请求路径**: `train.go` 和 `test.go` 中的 http 请求路径需要配置(本机 / 服务器)

在 `main.go` 同级目录下运行 `go run main.go` 执行脚本

## 运行结果

控制台输出:

```
train:
all queries:  1000
can't find the answer:  4
right:  972

test:
all queries:  1292
no:  713
ok:  579
```

> 总共运行时间大概是一个多小时
> 如果运行结果相差巨大那可能是数据库连接出错, 需要等一会儿再次运行
