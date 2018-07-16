package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/nanamen/go-echo-rest-sample/domain/model"
	"github.com/nanamen/go-echo-rest-sample/usecase"
	"github.com/labstack/echo"
)

// UserHandler interface
type UserHandler interface {
	GetUsers(c echo.Context) error
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userHandler struct {
	UserUseCase usecase.UserUseCase
}

// NewUserHandler UserHandlerを取得します.
func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	users, err := h.UserUseCase.GetUsers(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Users does not exist.")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.GetUser(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User does not exist.")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) CreateUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.CreateUser(ctx, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Create.")
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := h.UserUseCase.UpdateUser(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Create.")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID must be int")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := h.UserUseCase.DeleteUser(ctx, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "User can not Delete.")
	}

	return c.NoContent(http.StatusNoContent)
}
