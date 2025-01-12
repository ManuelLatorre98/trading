package dependencyContainer

import (
	"database/sql"
	userController "manulatorre98/trading/controller"
	"manulatorre98/trading/model"
	"manulatorre98/trading/repository/userRepository"
)

var (
	UserRepo model.UserRepository
	UserCtrl *userController.UsuarioController
)

type DependencyContainer struct {
	/*The repos must be accessed only by the controller, so are stored private in the container*/
	userRepo model.UserRepository
	UserCtrl *userController.UsuarioController
	/*Here put new controllers and repos*/
}

func NewDependencyContainer(db *sql.DB) *DependencyContainer {
	userRepo := userRepository.NewUserLocalRepository(db)
	userCtrl := userController.NewUsuarioController(userRepo)
	/*Here initialize the new repos and ctrlrs*/
	return &DependencyContainer{
		userRepo: userRepo,
		UserCtrl: userCtrl,
	}
}
