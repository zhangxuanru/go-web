package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"config"
	"router"
	"runtime"
)

func main() {
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir(config.STATIC_DIR))))
    router.InitRouter()
	runtime.GOMAXPROCS(runtime.NumCPU())
	server := &http.Server{
		Addr:    ":80",
		Handler: http.DefaultServeMux,
	}
	quitChan := make(chan os.Signal)
	signal.Notify(
		quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	go func() {
		<-quitChan
		server.Close()
	}()
    go server.ListenAndServe()
	<-quitChan
}


