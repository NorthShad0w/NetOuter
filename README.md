# NetOuter

## 目的

一个无依赖可执行文件，能快速测试受害者端口出网情况，为后续建立代理隧道收集必要信息。

## 原理

### tcp出网探测

根据tcp的特性，程序尝试对指定外网的ip:port建立tcp全连接，若连接建立，则认为该端口出网。  

### udp出网探测

根据udp的特性，基本无法从连接的角度确认出网情况。程序内置了一些常用udp端口的客户端库，针对不同的端口构造不同的基于udp的应用层协议的数据包并发送，如果接受到服务端的回复，则认为该端口出网。

### dns隧道建立可能性探测

除了53端口直接出网，dns可以通过递归或者迭代查询间接出网构造dns隧道。程序会尝试本地dns解析，若解析成功，则可能构造dns隧道。

## 使用

target目录下文件

### 编译

#### linux版本

```bash
CGO_ENABLED=0 go build -ldflags "-w -s" -o target/release/NetOuter ./cmd/netouter/main.go
```

#### windows版本

```bash
GOOS="windows" GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-w -s" -o target/release/NetOuter.exe ./cmd/netouter/main.go
```

### 命令行用法

#### 测试常规出网可能性 - 推荐

直接运行

```bash
NetOuter
```

#### tcp全端口

端口太多比较慢

NetOuter -tcp

#### TODO

1. UDP 162,623,520,10161,10162
2. tcp测试 使用golang特性加快速度
3. ipv6出网探测和逻辑


#### 还有其他的出网方式就是基于代理出网 sock/http/https

这个需要信息收集，针对不同的机器收集路由和配置的代理ip端口用户密码




