package idGenerator

import "time"

type DefaultIDGenerator struct {
	idChan chan uint64
}

func (d *DefaultIDGenerator) OnInit() {
	go func() {
		for {
			d.idChan <- newID()
		}
	}()
}

func (d *DefaultIDGenerator) GetIDChan() chan uint64 {
	return d.idChan
}

func newID() uint64 {
	return uint64(time.Now().Unix())
}

func InitDefaultIDGenerator() IDGenerator {
	g := DefaultIDGenerator{
		idChan: make(chan uint64),
	}
	g.OnInit()
	return &g
}
