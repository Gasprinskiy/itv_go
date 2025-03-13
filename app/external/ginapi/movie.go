package external

import (
	"itv_go/config"
	"itv_go/external/middleware"
	"itv_go/internal/entity/global"
	"itv_go/internal/entity/movie"
	"itv_go/internal/usecase"
	"itv_go/tools/validator"
	"net/http"
	"strconv"

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

	movieGroup := ext.gin.Group("/movies")
	{
		movieGroup.POST(
			"",
			ext.authMiddleware.CheckAccesToken(),
			ext.CreateMovie,
		)
		movieGroup.GET(
			"",
			ext.authMiddleware.CheckAccesToken(),
			ext.GetMovieList,
		)
		movieGroup.GET(
			"/:id",
			ext.authMiddleware.CheckAccesToken(),
			ext.GetMovieByID,
		)
		movieGroup.PUT(
			"/:id",
			ext.authMiddleware.CheckAccesToken(),
			ext.UpdateMovie,
		)
		movieGroup.DELETE(
			"/:id",
			ext.authMiddleware.CheckAccesToken(),
			ext.DeleteMovie,
		)
	}

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
// @Router       /movies [post]
func (e *MovieExternal) CreateMovie(c *gin.Context) {
	param := movie.CreateMovieRecordParam{}

	if err := c.BindJSON(&param); err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": err.Error()})
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

// @Summary      Get movie list
// @Description  Returns a list of all movies (authorized users only)
// @Tags         Movie
// @Produce      json
// @Success      200  {array}   movie.Movie  "List of movies"
// @Failure      401  {object}  global.MessageResponse "Unauthorized"
// @Failure      500  {object}  global.MessageResponse "Internal server error"
// @Router       /movies [get]
// @Router /movies [get]
func (e *MovieExternal) GetMovieList(c *gin.Context) {
	movies, err := e.movieUsecase.GetMovieList()
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// @Summary      Get a movie by ID
// @Description  Returns movie details by ID (authorized users only)
// @Tags         Movie
// @Produce      json
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  movie.Movie  "Movie details"
// @Failure      400  {object}  global.MessageResponse "Invalid movie ID"
// @Failure      401  {object}  global.MessageResponse "Unauthorized"
// @Failure      404  {object}  global.MessageResponse "No movie found"
// @Failure      500  {object}  global.MessageResponse "Internal server error"
// @Router       /movies/{id} [get]
func (e *MovieExternal) GetMovieByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	movie, err := e.movieUsecase.GetMovieByID(id)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// @Summary      Update a movie
// @Description  Updates movie information by ID (authorized users only)
// @Tags         Movie
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "Movie ID"
// @Param        movie body      movie.CreateMovieRecordParam  true  "Movie details"
// @Success      200   {object}  global.CreatedOrUpdatedResponse  "Updated movie ID"
// @Failure      401   {object}  global.MessageResponse "Unauthorized"
// @Failure      400   {object}  global.MessageResponse "Invalid request payload"
// @Failure      500   {object}  global.MessageResponse "Internal server error"
// @Router       /movies/{id} [put]
func (e *MovieExternal) UpdateMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	param := movie.CreateMovieRecordParam{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": err.Error()})
		return
	}

	if err := validator.ValidateStruct(param); err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": err.Error()})
		return
	}

	movie := movie.Movie{
		ID:       id,
		Title:    param.Title,
		Director: param.Director,
		Plot:     param.Plot,
		Year:     param.Year,
	}

	updatedID, err := e.movieUsecase.UpdateMovie(movie)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": updatedID})
}

// @Summary      Delete a movie
// @Description  Deletes a movie by ID (authorized users only)
// @Tags         Movie
// @Produce      json
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  global.CreatedOrUpdatedResponse  "Deleted movie ID"
// @Failure      400  {object}  global.MessageResponse "Invalid movie ID"
// @Failure      401  {object}  global.MessageResponse "Unauthorized"
// @Failure      500  {object}  global.MessageResponse "Internal server error"
// @Router       /movies/{id} [delete]
func (e *MovieExternal) DeleteMovie(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(global.ErrStatusCodes[global.ErrInvalidParam], gin.H{"message": global.ErrInvalidParam.Error()})
		return
	}

	deletedID, err := e.movieUsecase.DeleteMovie(id)
	if err != nil {
		c.JSON(global.ErrStatusCodes[err], gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": deletedID})
}
