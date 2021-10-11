package traffic

import (
	"errors"
	"net"
	"os"
	"sync"

	"github.com/eolinker/eosc/log"
)

var (
	ErrorNotTcpListener = errors.New("not tcp listener")
)

type iRemove interface {
	remove(name string)
}

type tListener struct {
	net.Listener
	once      sync.Once
	parent    iRemove
	file      *os.File
	fileError error
}

func newTTcpListener(listener net.Listener, parent iRemove) *tListener {

	return &tListener{Listener: listener, parent: parent}
}

func (t *tListener) Close() error {
	log.Debug("tListener close try")
	t.once.Do(func() {

		name := toName(t.Listener)
		log.Info("shutdown listener:", name)
		t.parent.remove(name)
		err := t.Listener.Close()
		if err != nil {
			log.Warn("close listener:", err)
		}
		if t.file != nil {
			t.file.Close()
		}
	})
	log.Debug("tListener close done")
	return nil
}

func (t *tListener) File() (*os.File, error) {
	if t.file == nil && t.fileError == nil {
		if tcp, ok := t.Listener.(*net.TCPListener); ok {

			t.file, t.fileError = tcp.File()
		} else {
			t.fileError = ErrorNotTcpListener
		}
	}
	return t.file, t.fileError
}
