package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/irvingdinh/osgarth-api/src/component/logger"
	"github.com/irvingdinh/osgarth-api/src/component/repository"
)

type ItemHandler interface {
	GetItem(c *gin.Context)
}

func NewItemHandler(
	repositoryClient repository.Repository,
) ItemHandler {
	return &itemHandlerImpl{
		repositoryClient: repositoryClient,
	}
}

type itemHandlerImpl struct {
	repositoryClient repository.Repository
}

func (i *itemHandlerImpl) GetItem(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.CToL(ctx, "ItemHandler.GetItem")

	slug := c.Param("slug")

	item, err := i.repositoryClient.ItemRepository().FindOneBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		log.WithError(err).Error("FindOneBySlug returns unexpected error")

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"to": item.Payload["to"],
	})
}
