package handler

//go:generate mockery --name=Handler --case=snake
type Handler interface {
	PingHandler() PingHandler
}

func New() Handler {
	return &handlerImpl{
		pingHandler: NewPingHandler(),
	}
}

type handlerImpl struct {
	pingHandler PingHandler
}

func (i *handlerImpl) PingHandler() PingHandler {
	return i.pingHandler
}
