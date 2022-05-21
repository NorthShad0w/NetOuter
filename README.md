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

```bash
NetOuter d
```

硬编码测试以下公开服务

- windows ntp  
- 114 dns
- local dns
- cloudflare quic
- 百度 tcp 80
- 百度 tcp 443
- 114 tcp dns



#### 测试自定义tcp端口出网可能性，无法硬编码的端口，可以上fofa找端口

文件名随意
```bash
NetOuter a ./targets.txt
```

fofa语法

```
port="69"
```

例子：
targets.txt  

```texinfo
192.168.1.1:21
192.168.1.1:22
192.168.1.1:23
192.168.1.1:25
192.168.1.1:3306
192.168.1.1:3389
```

#### 利用老外公开的一个1-65535全开的服务器测试出网，注意该服务器在美国不是很opsec

文件名随意

```bash
NetOuter b ./targets.txt
```

源码configs目录下有top open ports 可以直接上传使用

例子：
targets.txt  

```texinfo
21
22
23
25
3306
3389
```

#### 测试snmp出网可能性，注意，window下defender可能行为报毒

udp 161 出网 不好硬编码  
fofa可以找公开

```
port="161"
```

不用加端口
```bash
NetOuter snmp 192.168.1.1
```

#### 测试tftp出网

UDP 69出网 不好硬编码  
fofa可以找公开

```
port="69"
```

不用加端口
```
NetOuter snmp 192.168.1.1
```

#### 测试icmp出网

直接发icmp包要管理员权限，ping 默认是setuid的，建议出网测试时顺手ping一下吧，就不集成到软件里了

```bash
ping 114.114.114.114
```

#### TODO

1. UDP 162,623,10161,10162
2. tcp测试 使用golang特性加快速度


#### 还有其他的出网方式就是基于代理出网 sock/http/https

这个需要信息收集，针对不同的机器收集路由和配置的代理ip端口用户密码




