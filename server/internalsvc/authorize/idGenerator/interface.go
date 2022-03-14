package idGenerator

type IDGenerator interface {
	GetIDChan() chan uint64
	OnInit()
}
