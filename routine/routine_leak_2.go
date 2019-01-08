// Reference: Goroutine leak -> https://medium.com/golangspec/goroutine-leak-400063aef468
/*
// Writing to nil channel blocks forever

package main

func main() {
	var ch chan struct{}
	ch <- struct{}{}
}
*/
/*
Output:
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send (nil chan)]:
main.main()
        /root/go/src/go_json/routine/routine_leak_2.go:5 +0x34
		exit status 2
*/

/*
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	if false {
		ch = make(chan int, 1)
		ch <- 1
	}
	go func(ch chan int) {
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}
*/

/*
Output:
goroutines: 2
goroutines: 2
goroutines: 2
goroutines: 2
goroutines: 2
goroutines: 2
goroutines: 2
goroutines: 2
*/

// Analysis

/*
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.Println(http.ListenAndServe("10.206.156.128:6060", nil))
}
*/

/*
Output:
goroutine profile: total 3
1 @ 0x42e01a 0x42933a 0x4289b7 0x48d4bb 0x48d53d 0x48e39d 0x565a1f 0x576b2a 0x6571ba 0x45aa41
#	0x4289b6	internal/poll.runtime_pollWait+0x56		/usr/local/go/src/runtime/netpoll.go:173
#	0x48d4ba	internal/poll.(*pollDesc).wait+0x9a		/usr/local/go/src/internal/poll/fd_poll_runtime.go:85
#	0x48d53c	internal/poll.(*pollDesc).waitRead+0x3c		/usr/local/go/src/internal/poll/fd_poll_runtime.go:90
#	0x48e39c	internal/poll.(*FD).Read+0x17c			/usr/local/go/src/internal/poll/fd_unix.go:157
#	0x565a1e	net.(*netFD).Read+0x4e				/usr/local/go/src/net/fd_unix.go:202
#	0x576b29	net.(*conn).Read+0x69				/usr/local/go/src/net/net.go:176
#	0x6571b9	net/http.(*connReader).backgroundRead+0x59	/usr/local/go/src/net/http/server.go:668

1 @ 0x42e01a 0x42933a 0x4289b7 0x48d4bb 0x48d53d 0x48f938 0x566332 0x5804ae 0x57ea89 0x661a3f 0x6607d5 0x660529 0x66156a 0x6d7e7e 0x42dbc2 0x45aa41
#	0x4289b6	internal/poll.runtime_pollWait+0x56		/usr/local/go/src/runtime/netpoll.go:173
#	0x48d4ba	internal/poll.(*pollDesc).wait+0x9a		/usr/local/go/src/internal/poll/fd_poll_runtime.go:85
#	0x48d53c	internal/poll.(*pollDesc).waitRead+0x3c		/usr/local/go/src/internal/poll/fd_poll_runtime.go:90
#	0x48f937	internal/poll.(*FD).Accept+0x1a7		/usr/local/go/src/internal/poll/fd_unix.go:372
#	0x566331	net.(*netFD).accept+0x41			/usr/local/go/src/net/fd_unix.go:238
#	0x5804ad	net.(*TCPListener).accept+0x2d			/usr/local/go/src/net/tcpsock_posix.go:136
#	0x57ea88	net.(*TCPListener).AcceptTCP+0x48		/usr/local/go/src/net/tcpsock.go:246
#	0x661a3e	net/http.tcpKeepAliveListener.Accept+0x2e	/usr/local/go/src/net/http/server.go:3216
#	0x6607d4	net/http.(*Server).Serve+0x1a4			/usr/local/go/src/net/http/server.go:2770
#	0x660528	net/http.(*Server).ListenAndServe+0xa8		/usr/local/go/src/net/http/server.go:2711
#	0x661569	net/http.ListenAndServe+0x79			/usr/local/go/src/net/http/server.go:2969
#	0x6d7e7d	main.main+0x3d					/root/go/src/go_json/routine/routine_leak_2.go:70
#	0x42dbc1	runtime.main+0x211				/usr/local/go/src/runtime/proc.go:198

1 @ 0x6cf728 0x6cf530 0x6cc074 0x6d783d 0x6d7bc1 0x65d764 0x65f3d0 0x66040c 0x65c781 0x45aa41
#	0x6cf727	runtime/pprof.writeRuntimeProfile+0x97	/usr/local/go/src/runtime/pprof/pprof.go:679
#	0x6cf52f	runtime/pprof.writeGoroutine+0x9f	/usr/local/go/src/runtime/pprof/pprof.go:641
#	0x6cc073	runtime/pprof.(*Profile).WriteTo+0x3e3	/usr/local/go/src/runtime/pprof/pprof.go:310
#	0x6d783c	net/http/pprof.handler.ServeHTTP+0x20c	/usr/local/go/src/net/http/pprof/pprof.go:243
#	0x6d7bc0	net/http/pprof.Index+0x1d0		/usr/local/go/src/net/http/pprof/pprof.go:254
#	0x65d763	net/http.HandlerFunc.ServeHTTP+0x43	/usr/local/go/src/net/http/server.go:1947
#	0x65f3cf	net/http.(*ServeMux).ServeHTTP+0x12f	/usr/local/go/src/net/http/server.go:2337
#	0x66040b	net/http.serverHandler.ServeHTTP+0xbb	/usr/local/go/src/net/http/server.go:2694
#	0x65c780	net/http.(*conn).serve+0x650		/usr/local/go/src/net/http/server.go:1830
*/

package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
}

/*
goroutine profile: total 1
1 @ 0x4a75f8 0x4a7400 0x4a3f44 0x4ad8a5 0x429a82 0x450ed1
#       0x4a75f7        runtime/pprof.writeRuntimeProfile+0x97  /usr/local/go/src/runtime/pprof/pprof.go:679
#       0x4a73ff        runtime/pprof.writeGoroutine+0x9f       /usr/local/go/src/runtime/pprof/pprof.go:641
#       0x4a3f43        runtime/pprof.(*Profile).WriteTo+0x3e3  /usr/local/go/src/runtime/pprof/pprof.go:310
#       0x4ad8a4        main.main+0x64                          /root/go/src/go_json/routine/routine_leak_2.go:122
#       0x429a81        runtime.main+0x211                      /usr/local/go/src/runtime/proc.go:198
*/
