package handler

import "github.com/irvingdinh/osgarth-api/src/component/repository"

//go:generate mockery --name=Handler --case=snake
type Handler interface {
	ItemHandler() ItemHandler
	PingHandler() PingHandler
}

func New(
	repositoryClient repository.Repository,
) Handler {
	return &handlerImpl{
		itemHandler: NewItemHandler(repositoryClient),
		pingHandler: NewPingHandler(),
	}
}

type handlerImpl struct {
	itemHandler ItemHandler
	pingHandler PingHandler
}

func (i *handlerImpl) ItemHandler() ItemHandler {
	return i.itemHandler
}

func (i *handlerImpl) PingHandler() PingHandler {
	return i.pingHandler
}
