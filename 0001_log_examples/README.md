# example_basic.go

``` go
$ go run example_basic.go 
2022/12/01 08:10:42 Jack login, age:18
2022/12/01 08:10:42 Oh, system error when Jack login
panic: Oh, system error when Jack login

goroutine 1 [running]:
log.Panicf({0x4a6d31?, 0x4?}, {0xc00006bf40?, 0xc000104f70?, 0x405359?})
	/home/wwd/Software/go/src/log/log.go:395 +0x67
main.main()
	/home/wwd/all/libwwd/src/tools/go/0001_log_examples/example_basic.go:17 +0xc5
exit status 2

```


# 