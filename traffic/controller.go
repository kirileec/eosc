/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package traffic

import (
	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/utils"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"os"
)


type Traffics []*Out

type IController interface {
	ITraffic
	Encode(startIndex int)([]byte,[]*os.File,error)
	Close()
}

type Controller struct {
	Traffic
}
func (c *Controller) Encode(startIndex int) ([]byte, []*os.File,error) {
	ts:=c.All()
	pts:=new(PbTraffics)
	files:=make([]*os.File,0,len(ts))
	pts.Traffic = make([]*PbTraffic,0,len(ts))
	for i,ln:=range ts{
		file,err:=ln.File()
		if err!= nil{
			continue
		}
		ln.Close()
		addr:= ln.Addr()
		pt:=&PbTraffic{
			FD: uint64(i+startIndex),
			Addr:    addr.String(),
			Network: addr.Network(),
		}
		pts.Traffic = append(pts.Traffic, pt)
		files = append(files, file)
	}

	data, err := proto.Marshal(pts)
	if err!= nil{
		return nil,nil, err
	}

	return utils.EncodeFrame( data),files,nil

}

func (c *Controller) All() []*net.TCPListener {
	c.locker.Lock()
	list := c.data.List()
	c.data = eosc.NewUntyped()
	c.locker.Unlock()

	ts:=make([]*net.TCPListener,0,len(list))
	for _,it:=range list{
		tf,ok:= it.(*net.TCPListener)
		if !ok{
			continue
		}
		ts = append(ts, tf)
	}

	return ts
}

func NewController(r io.Reader) (*Controller) {
	c:= &Controller{
		Traffic:Traffic{
			data: eosc.NewUntyped(),
		},
	}
	c.Read(r)
	return c
}


func (c *Controller) ListenTcp(network string, addr string)(net.Listener,error) {

	tcp, err := c.Traffic.ListenTcp(network, addr)
	if err != nil {
		return nil, err
	}
	if tcp == nil{
		c.locker.Lock()
		defer c.locker.Unlock()
		tcpAddr, err := net.ResolveTCPAddr(network, addr)
		if err != nil {
			return nil, err
		}

		l,err:=net.ListenTCP(network,tcpAddr)
		if err!= nil{
			return nil,err
		}

		c.Traffic.add(l)
		tcp = l
	}
	return tcp,nil
}


func (c *Controller) Close() {
	list := c.data.List()
	c.data = eosc.NewUntyped()

	for _,it:=range list{
		tf,ok:=it.(*Out)
		if !ok{
			continue
		}
		tf.Listener.Close()
		tf.File.Close()
	}
}

