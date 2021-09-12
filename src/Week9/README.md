# Solutions of TCP Sticky packets

## 什么是粘包？

粘包（Sticky packet） 是指当 client  端发送了两条消息 `ABC` 和 `DEF`，但在 server 端接收到的却是 `ABCDE`，像这种一次性读取了两条数据的情况就叫粘包。

## 出现粘包的原因

TCP 是面向连接的传输协议，TCP 传输的数据是以流的形式，而流数据是没有明确的开始结尾边界，所以 TCP 也没法判断哪一段流属于一条消息。
粘包产生的原因：
1. client 每次发送的数据 < socket 缓冲区大小
2. server 读取 socket 缓冲区数据不够及时

server_sticky.go 中的读取方式会出现粘包

```bash
received msg from client, the msg is  0aa0b the length of msg is  5
received msg from client, the msg is  bbb0c the length of msg is  5
received msg from client, the msg is  cccc1 the length of msg is  5
received msg from client, the msg is  aa1bb the length of msg is  5
received msg from client, the msg is  bb1cc the length of msg is  5
received msg from client, the msg is  ccc2a the length of msg is  5
received msg from client, the msg is  a2bbb the length of msg is  5
```

## 粘包的解决方案

### FIX Length

> client 和 server 规定固定大小的缓冲区，当字符长度不够时使用空字符弥补

server_fixL.go 和 ficL_test.go 为实现代码

### Delimiter Based

> 以特殊字符结尾

server_delimiter.go 为代码实现

### Length field based frame decoder
