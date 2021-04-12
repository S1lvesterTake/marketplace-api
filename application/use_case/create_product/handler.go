package create_product

import (
	"context"
	"net/http"

	"github.com/S1lvesterTake/marketplace-api/application/helper"
	"github.com/S1lvesterTake/marketplace-api/application/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/refactory-id/go-core-package/response"
)

type CreateProductHandler struct {
	request    infrastructure.Request
	repository infrastructure.ProductRepository
}

func NewCreateProductHandler(request infrastructure.Request, repo infrastructure.ProductRepository) CreateProductHandler {
	return CreateProductHandler{
		request:    request,
		repository: repo,
	}
}

func (handler *CreateProductHandler) CreateProductHandler(c *gin.Context) {
	request := CreateProductRequest{}
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, response.SetMessage(err.Error(), false))
		return
	}

	if ok, err := ValidateRequest(&request); !ok {
		errRequest := helper.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, response.SetMessage(errRequest, false))
		return
	}

	product, err := handler.repository.CreateProduct(ctx, RequestMapper(request))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.SetMessage(err.Error(), false))
		return
	}

	c.JSON(http.StatusCreated, NewCreateProductResponse(product, "product created", true))

}
