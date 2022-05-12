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

这个基本够用吧，一些公开服务有不变地址的就写死在里面包括。
- windows ntp  
- 114 dns
- local dns
- cloudflare quic
- baidu 80
- baidu 443
- 114 tcp dns


```bash
NetOuter d
```

#### 测试自定义tcp端口出网可能性

一些端口比如22，公开的服务器可能不是一直开着22，也不好写死

```bash
NetOuter tcp ./targets.txt
```

端口就fofa上找公开的  
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

udp 161 出网我遇到过，这也不好写死。
windows下调用这个方法 defender报毒不知道为什么。

```bash
NetOuter snmp 192.168.1.1
```

#### 测试tftp出网

udp 69 我还没遇到过

#### 测试icmp出网

发icmp包要管理员权限，ping 默认是setudp的。

```bash
ping 114.114.114.114
```

#### 还有其他的出网方式就是基于代理出网 sock/http/https

这个需要信息收集，针对不同的机器收集路由，配置的代理ip端口用户密码
不是很好写死。




