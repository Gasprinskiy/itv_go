package external

import (
	"itv_go/config"
	"itv_go/external/middleware"
	"itv_go/internal/entity/global"
	"itv_go/internal/entity/movie"
	"itv_go/internal/usecase"
	"itv_go/tools/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieExternal struct {
	gin            *gin.Engine
	cnfg           *config.Config
	movieUsecase   *usecase.MovieUsecase
	authMiddleware *middleware.AuthMiddleware
}

func NewMovieExternal(
	gin *gin.Engine,
	cnfg *config.Config,
	movieUsecase *usecase.MovieUsecase,
	authMiddleware *middleware.AuthMiddleware,
) *MovieExternal {
	ext := MovieExternal{
		gin,
		cnfg,
		movieUsecase,
		authMiddleware,
	}

	ext.gin.POST("/movie", ext.authMiddleware.CheckAccesToken(), ext.CreateMovie)

	return &ext
}

// @Summary      Create a movie
// @Description  Creates a new movie entry (authorized users only)
// @Tags         Movie
// @Accept       json
// @Produce      json
// @Param        movie body movie.CreateMovieRecordParam true "Movie details"
// @Success      201  {object}  global.CreatedOrUpdatedResponse  "Movie id"
// @Failure      400  {object}  global.MessageResponse "Invalid request parameters"
// @Failure      401  {object}  global.MessageResponse "Unauthorized"
// @Failure      500  {object}  global.MessageResponse "Internal server error"
// @Router       /movie [post]
func (e *MovieExternal) CreateMovie(c *gin.Context) {
	param := movie.CreateMovieRecordParam{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	if err := validator.ValidateStruct(param); err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": err.Error()})
		return
	}

	id, err := e.movieUsecase.CreateNewMovieRecord(param)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
