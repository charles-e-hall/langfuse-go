package transmission

type Queue interface {
	Add(t Trace) error
	Length() int
	Read() Trace
	ReadAll() []Trace
}

type LIFOQueue struct {
	Items           []Trace
	MaxPendingItems int
}
