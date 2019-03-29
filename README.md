# Golang

## Serials
* [golangspec][32]
* [50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs][34]
* [gophercises][37]

## Basic Info
* [go env installation offical guide][1]
```
download tarball
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
export go path in /etc/profile
```
* [How to write go code][2]
* [effective go][3]
* [tour of go][4]
* [Go实战--golang中读写文件的几种方式][13]
* [golang database][14]
* [go中文学习站][15]

## Closure
* [Closure in Golang][33]
* [Closure in Go: introduce and gottas][38]

## Share
* [知乎社区核心业务 Golang 化实践][29]

## Http
* [The complete guide to Go net/http timeouts][28]

## Best Practice
* [Practical Go: Real world advice for writing maintainable Go programs][27]

## Reflect
* [Golang的反射reflect深入理解和示例][23]
* [The Laws of Reflection(Go语言反射定律)][26]

## Interface
* [Interfaces in Go (part I)][30]
* [Interfaces in Go (part II)][31]

## Database
* [Configuring sql.DB for Better Performance][22]

## Cache
* [gomemcached][24]

## Microservice
* [go-micro][35]
* [go-kit][36]

## Third-Party Libraries
* [Log: logrus][20]

## Web Server Framework: beego
* [beego quickstart][5]
* [通过beego快速创建一个Restful风格API项目及API文档自动化][6]
* [beego中文文档][7]
* [bee github page][8]
* [beego orm advanced query][9]

## Web Server Framework: gin
* [gin-gonic/gin][12]
* [rsj217/golang-gin-mysql-curd.go][10]
* [Gin实战：Gin+Mysql简单的Restful风格的API][11]
* [reload-changed-without-restart-for-golang-web-gin][17]
* [codegangsta/gin][18]

## Docker Image
* [golang程序打包成容器镜像 ./app: not found错误: 可参考引用文章][21]

## gin Development
* use <codegangsta/gin> to do golang web application live reloading with docker
* docker file
```
# Version: 1.0.0
# Create Time: 2018-07-20
# Author: Jason Tu
# Description: Microservice Audit log img

FROM golang:latest

# Install golang library
RUN go get github.com/gin-gonic/gin && go get github.com/go-sql-driver/mysql \
    && go get github.com/codegangsta/gin

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["gin", "run", "*.go"]
# or CMD ["/bin/bash"]
```
* commands
```
docker run -i -d --rm --name auditlog_ms -p 8080:3000 -v /root/gin_playground/auditlog_ms:/go/src/auditlog_ms -w /go/src/auditlog_ms auditlog_ms:base
or
docker run -it --rm --name auditlog_ms -p 8080:3000 -v /root/gin_playground/auditlog_ms:/go/src/auditlog_ms -w /go/src/auditlog_ms auditlog_ms:base
```
* [refer][17]

## gin Deploy
* [how-to-deploy-a-go-web-application-with-docker][16]
* docker file
```
# Version: 1.0.0
# Create Time: 2018-07-20
# Author: Jason Tu
# Description: Microservice Audit log img

FROM golang:latest

# Create the directory where the application will reside
RUN mkdir /app

ADD auditlog_ms /app/auditlog_ms
ADD conf /app/conf

WORKDIR /app

EXPOSE 8080

# Set the entry point of the container to the application executable
ENTRYPOINT /app/auditlog_ms
```
* commands
```
docker run -i -d --rm --name auditlog_ms_prod -p 8081:3001 auditlog_ms:prod
```

## Dependency Management
* [kardianos/govendor][19]
* [go 依赖管理利器 -- govendor][25]
* govendor practice
```
go get -u github.com/kardianos/govendor
# then build govendor

# goto src code folder
govendor init
govendor add +external
# commit the library code
```
 

[1]: https://golang.org/doc/install
[2]: https://golang.org/doc/code.html#Workspaces
[3]: https://golang.org/doc/effective_go.html
[4]: https://tour.golang.org/
[5]: https://beego.me/quickstart
[6]: https://www.cnblogs.com/huligong1234/p/4707282.html
[7]: https://beego.me/docs/intro/ 
[8]: https://github.com/beego/bee
[9]: https://beego.me/docs/mvc/model/query.md
[10]: https://gist.github.com/rsj217/26492af115a083876570f003c64df118#file-golang-gin-mysql-restful-1-go
[11]: https://www.jianshu.com/p/a3f63b5da74c
[12]: https://github.com/gin-gonic/gin
[13]: https://blog.csdn.net/wangshubo1989/article/details/74777112/
[14]: http://go-database-sql.org/index.html
[15]: https://studygolang.com/pkgdoc
[16]: https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
[17]: http://www.itfanr.cc/2017/08/01/reload-changed-without-restart-for-golang-web-gin/
[18]: https://github.com/codegangsta/gin
[19]: https://github.com/kardianos/govendor
[20]: https://github.com/sirupsen/logrus
[21]: http://windgreen.me/2018/08/14/golang%E7%A8%8B%E5%BA%8F%E6%89%93%E5%8C%85%E6%88%90%E5%AE%B9%E5%99%A8%E9%95%9C%E5%83%8F-app-not-found%E9%94%99%E8%AF%AF/
[22]: https://www.alexedwards.net/blog/configuring-sqldb
[23]: https://juejin.im/post/5a75a4fb5188257a82110544
[24]: https://github.com/bradfitz/gomemcache
[25]: https://blog.csdn.net/yeasy/article/details/65935864
[26]: https://studygolang.com/articles/5834
[27]: https://dave.cheney.net/practical-go/presentations/qcon-china.html
[28]: https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
[29]: https://zhuanlan.zhihu.com/p/48039838?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io
[30]: https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c
[31]: https://medium.com/golangspec/interfaces-in-go-part-ii-d5057ffdb0a6
[32]: https://medium.com/golangspec
[33]: https://www.jianshu.com/p/3934e62d78a1
[34]: http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#closure_for_it_vars
[35]: https://github.com/micro/go-micro
[36]: https://github.com/go-kit/kit
[37]: https://gophercises.com/
[38]: https://www.calhoun.io/closures-in-go/
