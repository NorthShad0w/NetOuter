# NetOuter

## 目的

一个无依赖可执行文件，能快速测试受害者端口出网情况，为后续建立代理隧道收集必要信息。

## 原理

### tcp出网探测

根据tcp的特性，程序尝试对指定外网的ip:port建立tcp全连接，若连接建立，则认为该端口出网。  

### udp出网探测

根据udp的特性，基本无法从连接的角度确认出网情况。程序内置了一些常用udp端口的客户端库，针对不同的端口构造不同的基于udp的应用层协议的数据包并发送，如果接受到服务端的回复，则认为该端口出网。

### dns出网探测

除了53端口直接出网，dns可以通过递归或者迭代查询间接出网构造dns隧道。程序会尝试本地dns解析，若解析成功，则可能构造dns隧道。

## 使用

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

#### 测试常规出网可能性

```bash
NetOuter d
```

#### 测试自定义tcp端口出网可能性

```bash
NetOuter tcp ./targets.txt
```

targets.txt

```texinfo
192.168.1.1:21
192.168.1.1:22
192.168.1.1:23
192.168.1.1:25
192.168.1.1:3306
192.168.1.1:3389
```

#### 测试snmp出网可能性

```bash
NetOuter snmp 192.168.1.1
```


```


