package handlers

import (
	"net/http"
	"test-ozon/internal/service/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Request struct {
	Url string `json:"url" validate:"required,url"`
}

type ResponseOK struct {
	NewAlias string `json:"new-alias,omitempty"`
	Url string `json:"url,omitempty"`
	response.Response
}

func (h *Handlers) GetUrl(c *gin.Context) {
	param := c.Param("alias")

	if param == ""{
		c.JSON(http.StatusBadRequest, response.Error("empty alias"))
		return
	}

	url, err := h.Service.GetUrl(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("no results"))
		return
	}

	c.JSON(http.StatusOK, ResponseOK{
		Url:      url,
		Response: response.OK(),
	})
}

func (h *Handlers) SaveUrl(c *gin.Context) {
	var req Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("request body is empty"))
		return
	}
	if err := validator.New().Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, response.ValidationError(errs))
		return
	}

	newAlias, err := h.Service.SaveUrl(req.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, ResponseOK{
		Response: response.OK(),
		NewAlias: newAlias,
	})
}