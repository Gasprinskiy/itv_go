package external

import (
	"itv_go/internal/entity/global"
	appuser "itv_go/internal/entity/user"
	"itv_go/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserExternal struct {
	jwtUsecase   *usecase.JwtUsecase
	userUsercase *usecase.UserUsecase
	gin          *gin.Engine
}

func NewUserExternal(
	jwtUsecase *usecase.JwtUsecase,
	userUsercase *usecase.UserUsecase,
	gin *gin.Engine,
) *UserExternal {
	ext := UserExternal{
		jwtUsecase,
		userUsercase,
		gin,
	}

	ext.gin.POST("/register", ext.Register)
	ext.gin.POST("/auth", ext.Auth)

	return &ext
}

// @Summary      User registration
// @Description  Creates a new user with a login and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body appuser.CreateUserParams true "User credentials"
// @Success      201  {object}  appuser.UserResponse
// @Failure      400  {object}  global.ErrorResponse
// @Failure      500  {object}  global.ErrorResponse
// @Router       /register [post]
func (e *UserExternal) Register(c *gin.Context) {
	param := appuser.CreateUserParams{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	id, err := e.userUsercase.Register(param)
	if err != nil {
		c.JSON(global.StatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (e *UserExternal) Auth(c *gin.Context) {
	param := appuser.CreateUserParams{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	id, err := e.userUsercase.Auth(param)
	if err != nil {
		c.JSON(global.StatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	token, err := e.jwtUsecase.GenerateToken(id)
	if err != nil {
		c.JSON(global.StatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	// TO DO move token life time to .env
	c.SetCookie("access_token", token, 3600*24, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"id": id})
}
