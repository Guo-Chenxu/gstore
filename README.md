# 使用 [gstore](https://github.com/pkumod/gStore) 的 java api 进行数据库的查询

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

### gstore 安装

#### 环境配置

1. 检查 gcc/g++版本
    ```shell
    g++ --version
    ```
    需要版本大于 5.4, 否则需要安装
    ```shell
    wget http://ftp.tsukuba.wide.ad.jp/software/gcc/releases/gcc-5.4.0/gcc-5.4.0.tar.gz
    tar xvf gcc-5.4.0.tar.gz
    ./contrib/download_prerequisites
    cd ..
    mkdir gcc-build-5.4.0
    cd gcc-build-5.4.0
    ../gcc-5.4.0/configure --prefix=...(你的安装目录) --enable-checking=release --enable-languages=c,c++ --disable-multilib
    make -j4   #允许4个编译命令同时执行，加速编译过程
    make install
    ```
2. 安装 readline
    ```shell
    wget -c ftp://ftp.gnu.org/gnu/readline/readline-7.0.tar.gz
    tar -zxvf readline-7.0.tar.gz
    cd readline-7.0
    ./configure --prefix=/home/guochenxu/soft/readline # 指定软件安装路径
    make && make installl
    ```
3. 安装 boost (1.56-1.59)
    ```shell
     wget http://sourceforge.net/projects/boost/files/boost/1.56.0/boost_1_56_0.tar.gz
     tar -xzvf boost_1_56_0.tar.gz
     cd boost_1_56_0
     ./bootstrap.sh --prefix=/home/guochenxu/soft/boost
     ./b2
     ./b2 install
    ```
4. 安装 curl
    ```shell
    wget https://curl.haxx.se/download/curl-7.55.1.tar.gz
    tar -xzvf  curl-7.55.1.tar.gz
    cd curl-7.55.1
    ./configure --prefix=/home/guochenxu/soft/curl
    make
    make install
    ```
5. 安装 cmake
   检查版本
    ```shell
    cmake --version
    ```
    若没有则需要安装
    ```shell
    wget https://cmake.org/files/v3.6/cmake-3.6.2.tar.gz
    tar -xvf cmake-3.6.2.tar.gz && cd cmake-3.6.2/
    ./bootstrap --prefix=...
    make
    make install
    ```
6. 安装 pkg-config
   检查版本
    ```shell
    pkg-config --version
    ```
    这个系统应该都有安装
7. 安装 uuid
   从[这里](http://www.mirrorservice.org/sites/ftp.ossp.org/pkg/lib/uuid/)下载压缩包

    ```shell
    tar -zxvf libuuid-1.0.3.tar.gz
    cd libuuid-1.0.3
    ./configure --prefix=/home/guochenxu/soft/libuuid
    make
    make install
    ```

8. 安装 jemalloc
   从[这里](https://sourceforge.net/projects/jemalloc.mirror/files/5.2.1/)下载 jemalloc

    ```shell
    tar -zxvf jemalloc-5.2.1.tar.gz
    cd jemalloc-5.2.1
    ./configure --prefix=/home/guochenxu/soft/jemalloc
    make
    make install
    ```

9. 安装 openssl
   从[这里](https://github.com/openssl/openssl/releases)下载 openssl
    ```shell
    tar -zxvf openssl-3.1.0.tar.gz
    cd openssl-3.1.0
    ./Configure --prefix=/home/guochenxu/soft/openssl
    make
    make install
    ```
10. 安装 ncurses
    从[这里](https://mirrors.aliyun.com/gnu/ncurses/)下载 ncurses
    ```shell
    tar -zxvf ncurses.6.4.tar.gz
    cd ncurses.6.4
    ./configure --prefix=/home/guochenxu/soft/ncurses
    make
    make install
    ```
    ncurses 装完之后可能会缺一些动态编译文件, 需要从[这里](https://pan.baidu.com/s/1UR5bopukSXSQRSFHIQiwMg?pwd=3419)下载后复制到`/prefix/lib`目录下

全部安装完后需要配置环境变量, 打开 `~/.bashrc`, 添加以下环境变量配置 (路径要改成自己安装的路径)

```shell
export UUID_LIB_PATH=/home/guochenxu/soft/libuuid/lib/
export UUID_INCLUDE_PATH=/home/guochenxu/soft/libuuid/include/
export LD_LIBRARY_PATH=${UUID_LIB_PATH}:$LD_LIBRARY_PATH
export LIBRARY_PATH=${UUID_LIB_PATH}:$LIBRARY_PATH
export C_INCLUDE_PATH=${UUID_INCLUDE_PATH}:$C_INCLUDE_PATH
export CPLUS_INCLUDE_PATH=${UUID_INCLUDE_PATH}:$CPLUS_INCLUDE_PATH
export PKG_CONFIG_PATH=/home/guochenxu/soft/libuuid/lib/pkgconfig/:$PKG_CONFIG_PATH

export READLINE_LIB_PATH=/home/guochenxu/soft/readline/lib/
export READLINE_INCLUDE_PATH=/home/guochenxu/soft/readline/include/
export LD_LIBRARY_PATH=${READLINE_LIB_PATH}:$LD_LIBRARY_PATH
export LIBRARY_PATH=${READLINE_LIB_PATH}:$LIBRARY_PATH
export C_INCLUDE_PATH=${READLINE_INCLUDE_PATH}:$C_INCLUDE_PATH
export CPLUS_INCLUDE_PATH=${READLINE_INCLUDE_PATH}:$CPLUS_INCLUDE_PATH


export BOOST_LIB_PATH=/home/guochenxu/soft/boost/lib/
export BOOST_INCLUDE_PATH=/home/guochenxu/soft/boost/include/
export LD_LIBRARY_PATH=${BOOST_LIB_PATH}:$LD_LIBRARY_PATH
export LIBRARY_PATH=${BOOST_LIB_PATH}:$LIBRARY_PATH
export C_INCLUDE_PATH=${BOOST_INCLUDE_PATH}:$C_INCLUDE_PATH
export CPLUS_INCLUDE_PATH=${BOOST_INCLUDE_PATH}:$CPLUS_INCLUDE_PATH

export CURL_LIB_PATH=/home/guochenxu/soft/curl/lib/
export CURL_INCLUDE_PATH=/home/guochenxu/soft/curl/include/
export LD_LIBRARY_PATH=${CURL_LIB_PATH}:$LD_LIBRARY_PATH
export LIBRARY_PATH=${CURL_LIB_PATH}:$LIBRARY_PATH
export C_INCLUDE_PATH=${CURL_INCLUDE_PATH}:$C_INCLUDE_PATH
export CPLUS_INCLUDE_PATH=${CURL_INCLUDE_PATH}:$CPLUS_INCLUDE_PATH
export PKG_CONFIG_PATH=/home/guochenxu/soft/curl/lib/pkgconfig/:$PKG_CONFIG_PATH

export OPENSSL_LIB_PATH=/home/guochenxu/soft/openssl/lib64/
export OPENSSL_INCLUDE_PATH=/home/guochenxu/soft/openssl/include/
export LD_LIBRARY_PATH=${OPENSSL_LIB_PATH}:$LD_LIBRARY_PATH
export LIBRARY_PATH=${OPENSSL_LIB_PATH}:$LIBRARY_PATH
export C_INCLUDE_PATH=${OPENSSL_INCLUDE_PATH}:$C_INCLUDE_PATH
export CPLUS_INCLUDE_PATH=${OPENSSL_INCLUDE_PATH}:$CPLUS_INCLUDE_PATH
export PKG_CONFIG_PATH=/home/guochenxu/soft/openssl/lib64/pkgconfig/:$PKG_CONFIG_PATH

export NCURSES_LIB_PATH=/home/guochenxu/soft/ncurses/lib/
export NCURSES_INCLUDE_PATH=/home/guochenxu/soft/ncurses/include/
export LD_LIBRARY_PATH=${NCURSES_LIB_PATH}:$LD_LIBRARY_PATH
export LIBRARY_PATH=${NCURSES_LIB_PATH}:$LIBRARY_PATH
export C_INCLUDE_PATH=${NCURSES_INCLUDE_PATH}:$C_INCLUDE_PATH
export CPLUS_INCLUDE_PATH=${NCURSES_INCLUDE_PATH}:$CPLUS_INCLUDE_PATH

export JEMALLOC_LIB_PATH=/home/guochenxu/soft/jemalloc/lib/
export JEMALLOC_INCLUDE_PATH=/home/guochenxu/soft/jemalloc/include/
export LD_PRELOAD=${JEMALLOC_LIB_PATH}/libjemalloc.so:$LD_PRELOAD
export LD_LIBRARY_PATH=${JEMALLOC_LIB_PATH}:$LD_LIBRARY_PATH
export LIBRARY_PATH=${JEMALLOC_LIB_PATH}:$LIBRARY_PATH
export C_INCLUDE_PATH=${JEMALLOC_INCLUDE_PATH}:$C_INCLUDE_PATH
export CPLUS_INCLUDE_PATH=${JEMALLOC_INCLUDE_PATH}:$CPLUS_INCLUDE_PATH
```

运行 `source ~/.bashrc` 让配置的环境变量生效

#### 获取 gstore

在[这里](https://pan.baidu.com/s/19teS1sO3jyQCDvDhahQPQg?pwd=7qns)下载 gstore 并解压

```shell
cd gstore
```

先改一个 bug, 在 `makefile`文件的第 59 行加一个 `-lncurse` 参数, 修改完以后应该是这样:

```shell
library = -L/usr/lib64 -L./lib -L/usr/local/lib -L/usr/lib -I/usr/local/include/boost -ljemalloc -lreadline -lncurses -lantlr4-runtime -lgcov -lboost_thread -lboost_system -lboost_regex -lpthread -lcurl -llog4cplus
```

然后开始编译

```shell
# 以下命令可以通过 -j4 加速
make pre # 等待漫长的时候后, 如果没有报错的话就是成功了
make # 若编译顺利完成，最后会出现 Compilation ends successfully! 结果
```

### 创建数据库

在 gstore 根目录下执行

```shell
./bin/gbuild -db dbname -f filename
```

参数含义:

```
dbname：数据库名称
filename：数据集文件名称
```

例如:

```shell
./bin/gbuild -db CCKS -f /home/guochenxu/CCKS/triple.txt
```

### 启动数据库

在 gstore 根目录下执行

```shell
./bin/gserver -s
nohup ./bin/ghttp -p 9020 > ghttp.log & #启动 ghttp 服务, 端口号为9020, 并将日志输出到 ghttp.log
```

### 启动 java 程序

**环境配置**: java 1.8

**数据库配置**: GstoreConfiuration 中的配置信息

```shell
mvn package
java -jar ./target/gstore-0.0.1-SNAPSHOT.jar
```

## 程序执行流程

1. 接收请求, 加载数据库, 若加载失败直接返回 null
2. 在数据库中进行查询
    1. 若查询顺利返回格式为解析后的实体 (bindings 部分考虑到参数不固定, 所以未解析)
    2. 若查询超时 (设定定时器为 30 秒), 则返回 null

## 返回状态码

| 状态码 | 含义                     |
| ------ | ------------------------ |
| 200    | 查询正常                 |
| 500    | 数据库连接错误或查询出错 |
| 1001   | 查询结果为空             |

接口文档可以查看[这里](https://apifox.com/apidoc/shared-45a359ee-2428-4504-b9b7-3d20f3f71a20)
