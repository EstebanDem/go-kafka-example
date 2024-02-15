package framework

import (
	"net/http"
	"registration-service/internal/application/usecase"
	"registration-service/internal/domain/user"
	"registration-service/internal/infrastructure/inputport/rest"
	"registration-service/internal/infrastructure/storage"
)

func NewApplication() *http.ServeMux {
	mux := http.NewServeMux()

	// Dependency Injection
	appRepositories := applicationRepositories{
		userRepository: storage.NewUserInMemoryRepository(),
	}

	appUseCases := applicationUseCases{
		addUser:     usecase.NewAddUserUseCase(appRepositories.userRepository),
		getAllUsers: usecase.NewGetAllUsersUseCase(appRepositories.userRepository),
	}

	mux.HandleFunc("POST /users", rest.AddUserHandler(appUseCases.addUser))
	mux.HandleFunc("GET /users", rest.GetAllUsersHandler(appUseCases.getAllUsers))

	return mux
}

// Dependencies

type applicationRepositories struct {
	userRepository user.UserRepository
}

type applicationUseCases struct {
	addUser     usecase.AddUserUseCase
	getAllUsers usecase.GetAllUsersUseCase
}
