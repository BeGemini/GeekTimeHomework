# Goim  协议解码器

## Goim  协议设计

主要以包/帧方式：
- Package Length: 包长度 (4bytes)
- Header Length: 头长度 (2bytes)
- Protocol Version: 协议版本 (2bytes)
- Operation: 操作码 (4bytes)
- Sequence: 请求序列 Id (4bytes)
- Body: 包内容 (PackageLen - HeaderLen)

Operation:
- Auth
- HeartBeat
- Message

Sequence:
- 按请求、响应对应递增 Id

编码实现：
Goim/encoder

解码实现：
Goim/decoder

效果：
```bash
packageLen:36
headerLen:16
version:1
operation:3
sequence:10086
body:test goim decoder...
```