# Redis-benchmark

## redis-benchmark

```bash
-h 
-p
-s
```

## test

### 10 Bytes

```bash
[root ~]# redis-benchmark -q -d 10
PING_INLINE: 62150.41 requests per second, p50=0.407 msec
PING_MBULK: 61690.31 requests per second, p50=0.415 msec
SET: 62383.03 requests per second, p50=0.407 msec
GET: 61462.82 requests per second, p50=0.407 msec
```

### 20 Bytes

```bash
[root ~]# redis-benchmark -q -d 20
PING_INLINE: 61842.92 requests per second, p50=0.407 msec
PING_MBULK: 60938.45 requests per second, p50=0.407 msec
SET: 61728.39 requests per second, p50=0.407 msec
GET: 63051.70 requests per second, p50=0.407 msec
```

### 50 Bytes

```bash
[root@ ~]# redis-benchmark -q -d 50
PING_INLINE: 62227.75 requests per second, p50=0.407 msec
PING_MBULK: 62695.92 requests per second, p50=0.407 msec
SET: 61349.70 requests per second, p50=0.407 msec
GET: 62266.50 requests per second, p50=0.407 msec
```

### 100 Bytes

```bash
[root ~]# redis-benchmark -q -d 100
PING_INLINE: 61050.06 requests per second, p50=0.415 msec
PING_MBULK: 60496.07 requests per second, p50=0.415 msec
SET: 61614.29 requests per second, p50=0.415 msec
GET: 60532.69 requests per second, p50=0.415 msec
```

### 200 Bytes

```bash
[root ~]# redis-benchmark -q -d 200
PING_INLINE: 61387.36 requests per second, p50=0.407 msec
PING_MBULK: 61012.81 requests per second, p50=0.415 msec
SET: 62111.80 requests per second, p50=0.407 msec
GET: 60204.70 requests per second, p50=0.415 msec
```

### 1000 Bytes

```bash
[root ~]# redis-benchmark -q -d 1000
PING_INLINE: 61500.61 requests per second, p50=0.407 msec
PING_MBULK: 60569.35 requests per second, p50=0.415 msec
SET: 61500.61 requests per second, p50=0.415 msec
GET: 58479.53 requests per second, p50=0.431 msec
```

### 5000 Bytes

```bash
[root ~]# redis-benchmark -q -d 5000
PING_INLINE: 61728.39 requests per second, p50=0.407 msec
PING_MBULK: 60864.27 requests per second, p50=0.415 msec
SET: 57175.53 requests per second, p50=0.439 msec
GET: 55586.44 requests per second, p50=0.455 msec
```

### 统计表格

|      | 10       | 20       | 50      | 100      | 200     | 1000     | 5000     |
| ---- | -------- | -------- | ------- | -------- | ------- | -------- | -------- |
| SET  | 62383.03 | 61728.39 | 61349.7 | 61614.29 | 62111.8 | 61500.61 | 57175.53 |
| GET  | 61462.82 | 63051.7  | 62266.5 | 60532.69 | 60204.7 | 58479.53 | 55586.44 |

### 生成曲线图

![image-20210906084434640](https://github.com/BeGemini/PicBed/blob/main/img/20210906084445.png)

### 结论

随着数据的增大，GET SET 性能下降