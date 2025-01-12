package userController

/*import (
	"github.com/gin-gonic/gin"
	"manulatorre98/trading/middlewares"
	"manulatorre98/trading/model"
	"net/http"
)

type UsuarioController struct {
	userRepository model.UserRepository
}

func NewUsuarioController(usuarioRepository model.UserRepository) *UsuarioController {
	return &UsuarioController{userRepository: usuarioRepository}
}

func (c *UsuarioController) InsertUser(ctx *gin.Context) {
	var newUser model.User
	if err := ctx.BindJSON(&newUser); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, middlewares.ErrBadRequest)
		return
	}
	//Validate not existance of email
	user, _ := c.userRepository.GetUserByEmail(newUser.Email)
	if user != nil {
		ctx.AbortWithError(http.StatusConflict, middlewares.ErrConflict)
		return
	}
	//Validate not existance of username
	user, _ = c.userRepository.GetUserByUserName(newUser.Username)
	if user != nil {
		ctx.AbortWithError(http.StatusConflict, middlewares.ErrConflict)
		return
	}
	//Insert user
	r, err := c.userRepository.InsertUser(newUser)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, middlewares.ErrInternalServerError)
		return
	}
	newUser.ID = r
	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func (c *UsuarioController) GetUser(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.AbortWithError(http.StatusBadRequest, middlewares.ErrBadRequest)
		return
	}
	user, err := c.userRepository.GetUserByEmail(email)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, middlewares.ErrNotFound)
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}*/
