package controllers

import (
	"net/http"
	"strconv"

	"github.com/dsperax/management-api-go/internal/domain/entities"
	"github.com/dsperax/management-api-go/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_requests_total",
		Help: "Total de requisições recebidas",
	})
)

// UserController gerencia as requisições relacionadas a usuários.
type UserController struct {
	UserUseCase *usecases.UserUseCase
}

// GetUser godoc
// @Summary      Obter usuário por ID
// @Description  Retorna os dados de um usuário específico
// @Tags         users
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  entities.User
// @Failure      400  {object}  entities.ErrorResponse
// @Failure      500  {object}  entities.ErrorResponse
// @Router       /users/{id} [get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	requestsTotal.Inc()
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{Message: "ID inválido"})
		return
	}

	user, err := ctrl.UserUseCase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
