# dumpkey-frida

1. [install](https://github.com/frida/frida-go#installation) `frida-core-devkit`

2. `$ CGO_ENABLED=1 go run dumpkey.go`

```bash
$ CGO_ENABLED=1 go run dumpkey.go
0x6exxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxd
```

## 分析文章

- [打造macOS下最强的微信取证工具](https://mp.weixin.qq.com/s/dC8489jQ3jEC1AUrCBE6pw) 文章首发于 **百灵鸟安全团队**。
<img style="width: 150px;" src="https://user-images.githubusercontent.com/26270009/213902716-b10473d4-2408-4c19-8acc-65a4444b2dba.png">

## 扫描器优化版
- [kk 版扫描器](https://github.com/kekeimiku/dumpkey)