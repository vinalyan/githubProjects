package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processo interface {
	Process(e Event) error
}

type Type int

const (
	Unknow Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
