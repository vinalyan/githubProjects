package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processo interface {
}

type Event struct {
}
