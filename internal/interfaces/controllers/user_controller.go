package controllers

import (
	"net/http"
	"strconv"

	"github.com/dsperax/management-api-go/internal/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase *usecases.UserUseCase
}

// GetUser godoc
// @Summary      Obter usuário por ID
// @Description  Retorna os dados de um usuário específico
// @Tags         users
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  entities.User
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /users/{id} [get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := ctrl.UserUseCase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
