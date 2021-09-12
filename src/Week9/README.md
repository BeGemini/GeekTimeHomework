# Solutions of TCP Sticky packets

## 什么是粘包？

粘包（Sticky packet） 是指当 client  端发送了两条消息 `ABC` 和 `DEF`，但在 server 端接收到的却是 `ABCDE`，像这种一次性读取了两条数据的情况就叫粘包。

## 出现粘包的原因

TCP 是面向连接的传输协议，TCP 传输的数据是以流的形式，而流数据是没有明确的开始结尾边界，所以 TCP 也没法判断哪一段流属于一条消息。
粘包产生的原因：
1. client 每次发送的数据 < socket 缓冲区大小
2. server 读取 socket 缓冲区数据不够及时

server_sticky.go 中的读取方式会出现粘包

## 粘包的解决方案

