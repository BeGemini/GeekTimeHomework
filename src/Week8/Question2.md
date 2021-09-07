# INFO Memory

## 数据准备

`batch1.txt`，`batch2.txt` 为准备数据
使用命令批量 `SET` 数据：

```bash
cat batch1.txt | ./redis-cli
```

## 写入前的 INFO memory 信息
```bash
127.0.0.1:6379> INFO memory
# Memory
used_memory:162908720
used_memory_human:155.36M
used_memory_rss:165830656
used_memory_rss_human:158.15M
used_memory_peak:165948352
used_memory_peak_human:158.26M
used_memory_peak_perc:98.17%
used_memory_overhead:830544
used_memory_startup:809784
used_memory_dataset:162078176
used_memory_dataset_perc:99.99%
allocator_allocated:163340936
allocator_active:163741696
allocator_resident:167235584
total_system_memory:3973292032
total_system_memory_human:3.70G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.00
allocator_frag_bytes:400760
allocator_rss_ratio:1.02
allocator_rss_bytes:3493888
rss_overhead_ratio:0.99
rss_overhead_bytes:-1404928
mem_fragmentation_ratio:1.02
mem_fragmentation_bytes:2962960
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

## 写入后的 INFO memory 信息

```bash
127.0.0.1:6379> INFO memory
# Memory
used_memory:163839728
used_memory_human:156.25M
used_memory_rss:166961152
used_memory_rss_human:159.23M
used_memory_peak:165948352
used_memory_peak_human:158.26M
used_memory_peak_perc:98.73%
used_memory_overhead:1361552
used_memory_startup:809784
used_memory_dataset:162478176
used_memory_dataset_perc:99.66%
allocator_allocated:163888880
allocator_active:164179968
allocator_resident:167718912
total_system_memory:3973292032
total_system_memory_human:3.70G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.00
allocator_frag_bytes:291088
allocator_rss_ratio:1.02
allocator_rss_bytes:3538944
rss_overhead_ratio:1.00
rss_overhead_bytes:-757760
mem_fragmentation_ratio:1.02
mem_fragmentation_bytes:3162448
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:20496
mem_aof_buffer:0
mem_allocator:jemalloc-5.1.0
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

## 写入前后发生变化的参数对比

|                          | 写入前    | 写入后    |
| ------------------------ | --------- | --------- |
| used_memory              | 162908720 | 163839728 |
| used_memory_human        | 155.36M   | 156.25M   |
| used_memory_rss          | 165830656 | 166961152 |
| used_memory_rss_human    | 158.15M   | 159.23M   |
| used_memory_peak_perc    | 98.17%    | 98.73%    |
| used_memory_overhead     | 830544    | 1361552   |
| used_memory_dataset_perc | 99.99%    | 99.66%    |
| allocator_allocated      | 163340936 | 163888880 |
| allocator_active         | 163741696 | 164179968 |
| allocator_resident       | 167235584 | 167718912 |
| allocator_frag_bytes     | 400760    | 291088    |
| allocator_rss_bytes      | 3493888   | 3538944   |
| rss_overhead_bytes       | -1404928  | -757760   |
| mem_fragmentation_bytes  | 2962960   | 3162448   |