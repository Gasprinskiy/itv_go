package external

import (
	"fmt"
	"itv_go/config"
	"itv_go/external/middleware"
	"itv_go/internal/entity/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieExternal struct {
	// movieUsecase *usecase.
	gin            *gin.Engine
	cnfg           *config.Config
	authMiddleware *middleware.AuthMiddleware
}

func NewMovieExternal(
	// movieUsecase *usecase.MovieUsecase,
	gin *gin.Engine,
	cnfg *config.Config,
	authMiddleware *middleware.AuthMiddleware,
) *MovieExternal {
	ext := MovieExternal{
		gin,
		cnfg,
		authMiddleware,
	}

	// Роут с проверкой авторизации
	ext.gin.POST("/movie", ext.authMiddleware.CheckAccesToken(), ext.CreateMovie)

	return &ext
}

// @Summary      Create a movie
// @Description  Creates a new movie entry (authorized users only)
// @Tags         Movie
// @Accept       json
// @Produce      json
// @Param        movie body map[string]interface{} true "Movie details"
// @Success      201  {object}  map[string]interface{} "Movie created"
// @Failure      400  {object}  global.MessageResponse "Invalid request parameters"
// @Failure      401  {object}  global.MessageResponse "Unauthorized"
// @Failure      500  {object}  global.MessageResponse "Internal server error"
// @Router       /movie [post]
func (e *MovieExternal) CreateMovie(c *gin.Context) {
	var movieData map[string]interface{}
	if err := c.BindJSON(&movieData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	fmt.Println("SUCK")

	// id, err := e.movieUsecase.CreateMovie(movieData)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": global.ErrInternalError.Error()})
	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{"id": id})
}
