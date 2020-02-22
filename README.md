![Go](https://github.com/x1nchen/decimal/workflows/Go/badge.svg)

# Decimal
统一接口 Decimal，分别有 shopspring 实现与 tidb 实现。
shopspring 实现思路简单，内存分配开销大，计算力大。兼容性好。
tidb 实现思路复杂，内存分配小，计算力约是 shopspring 的五分之一。有精度限制（1e-30）


## 注意

tidb-decimal 官方只维护 linux/64bit 平台。经测试 windows 32bit 下测试有溢出问题

参考

1. https://gitlab.followme-internal.com/go-common/decimal/-/jobs/86272
2. https://github.com/pingcap/tidb/issues/5224