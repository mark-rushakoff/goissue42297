This repository is a reproducer for https://github.com/golang/go/issues/42297.

It turns out the actual reproducer is much simpler than I first expected,
but the repository can stick around for posterity's sake.

```
$ go version && go run main.go
go version go1.15.3 darwin/amd64
ok

$ gotip version && gotip run main.go
go version devel +256d729c0b Fri Oct 30 15:26:28 2020 +0000 darwin/amd64
panic: strconv.ParseFloat: parsing "0.5": invalid bit size 10

goroutine 1 [running]:
main.main()
	/Users/mr/tmp/goissue42297/main.go:12 +0xf7
exit status 2
```