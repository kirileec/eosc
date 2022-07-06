package main

import (
	"context"
	"fmt"
	"github.com/eolinker/eosc/config"
	"github.com/eolinker/eosc/etcd"
	"github.com/eolinker/eosc/log"
	"os"
	"time"

	"net"
	"net/http"
)

func main() {

	transport := log.NewTransport(os.Stderr, log.DebugLevel)
	transport.SetFormatter(&log.LineFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		CallerPrettyfier: nil,
	})
	log.Reset(transport)
	log.SetPrefix(fmt.Sprintf("[demo]"))
	conf, err := config.GetConfig()
	if err != nil {
		fmt.Println(conf)
		return
	}
	listen, err :=net.Listen("tcp",fmt.Sprintf("%s:%d",conf.Admin.IP,conf.Admin.Listen))
	if err != nil {
		fmt.Println(err)
		return
	}

	mux := http.NewServeMux()
	ser:=http.Server{
		Handler: mux,
	}
	go ser.Serve(listen)
	server, err := etcd.NewServer(context.Background(),mux)
	if err != nil {
		fmt.Println(err)
		return
	}
	server.Put("/abc","name")
	server.Watch("/",new(DemoHandler))

	t:=time.NewTicker(time.Second)
	for{
		<-t.C
		server.Put("/ab/time",fmt.Sprint(time.Now().Unix()))
	}

}

type DemoHandler struct {

}

func (d *DemoHandler) Put(key, value string) error {
	fmt.Printf("put:%s=%s\n",key,value)
	return nil
}

func (d *DemoHandler) Delete(key string) error {
	fmt.Printf("delete:%s\n",key)
	return nil
}

func (d *DemoHandler) Reset(values []*etcd.KValue) {
	fmt.Println("reset start===============")
	for _,v:=range values{
		fmt.Printf("\t%s=%s\n",string(v.Key),string(v.Value))
	}
	fmt.Println("reset end=============")
}

