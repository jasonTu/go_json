## Upgrade go to version 1.1.1+


### 1. Uninstall the existing version

As mentioned [here](https://golang.org/doc/install#install), to update a go version you will first need to uninstall the original version.

To uninstall, delete the `/usr/local/go` directory by:

```
$ sudo rm -rf /usr/local/go
```

### 2. Install the new version

Go to the [downloads](https://golang.org/dl/) page and download the binary release suitable for your system.

### 3. Extract the archive file

To extract the archive file:

```
$ sudo tar -C /usr/local -xzf /home/nikhita/Downloads/go1.8.1.linux-amd64.tar.gz
```

### 4. Make sure that your PATH contains `/usr/local/go/bin`

```
$ echo $PATH | grep "/usr/local/go/bin"
```

## Remove vendor folder
```
$ rm -rf service/vendor/
```

## Initial the go mod
```
go mod init auditlog_ms
go mod download  # deps will download to $GOPATH/pkg/mod.
GOOS=linux CGO_ENABLED=0 go build -tags netgo -a -v -o auditlog_ms *.go
```

## Refer
* [go mod 使用][1]
* [使用go mod结合docker分层缓存进行自动CI/CD][2]

[1]: https://juejin.im/post/5c887c105188257edb45e5b1.
[2]: https://juejin.im/post/5c8e503a6fb9a070d878184a
