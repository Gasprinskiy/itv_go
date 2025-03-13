package external

import (
	"itv_go/config"
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
	cnfg         *config.Config
}

func NewUserExternal(
	jwtUsecase *usecase.JwtUsecase,
	userUsercase *usecase.UserUsecase,
	gin *gin.Engine,
	cnfg *config.Config,
) *UserExternal {
	ext := UserExternal{
		jwtUsecase,
		userUsercase,
		gin,
		cnfg,
	}

	ext.gin.POST("/user/register", ext.Register)
	ext.gin.POST("/user/auth", ext.Auth)
	ext.gin.POST("/user/logout", ext.Logout)

	return &ext
}

// @Summary      User registration
// @Description  Creates a new user with a login and password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user body appuser.CreateUserParams true "User credentials"
// @Success      201  {object}  appuser.UserResponse "User ID"
// @Failure      400  {object}  global.MessageResponse "Invalid request parameters"
// @Failure      409  {object}  global.MessageResponse "User allready exists"
// @Failure      500  {object}  global.MessageResponse "Internal server error"
// @Router       /user/register [post]
func (e *UserExternal) Register(c *gin.Context) {
	param := appuser.CreateUserParams{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	id, err := e.userUsercase.Register(param)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// @Summary      User authentication
// @Description  Authenticates a user and returns a JWT token
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        credentials body appuser.CreateUserParams true "User login and password"
// @Success      200  {object}  appuser.UserResponse "User ID"
// @Failure      400  {object}  global.MessageResponse "Invalid request parameters"
// @Failure      401  {object}  global.MessageResponse "Unauthorized"
// @Failure      404  {object}  global.MessageResponse "Invalid login or password"
// @Failure      500  {object}  global.MessageResponse "Internal server error"
// @Router       /user/auth [post]
func (e *UserExternal) Auth(c *gin.Context) {
	param := appuser.CreateUserParams{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	id, err := e.userUsercase.Auth(param)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	token, err := e.jwtUsecase.GenerateToken(id)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.SetCookie("access_token", token, 3600*e.cnfg.JwtSecretLifeTime, "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary      User logout
// @Description  Logs out the user by deleting the access token cookie
// @Tags         User
// @Success      200  {object}  global.MessageResponse "Successfully logged out"
// @Router       /user/logout [post]
func (e *UserExternal) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.JSON(global.SuccessStatuses[global.SuccessLogedOut], gin.H{"message": global.SuccessLogedOut})
}
