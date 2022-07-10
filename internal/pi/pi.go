package pi

import (
	"net"
	"sync"
)

type PiBot struct {
	conn net.Conn
}

type PiBotMgr struct {
	pibots map[string]PiBot
}

var piBotMgr *PiBotMgr
var once sync.Once

func GetPiBotMgr() *PiBotMgr {
	once.Do(func() {
		piBotMgr = &PiBotMgr{}
	})
	return piBotMgr
}
