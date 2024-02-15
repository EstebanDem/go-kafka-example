package framework

import (
	"net/http"
	"registration-service/internal/application/service"
	"registration-service/internal/application/usecase"
	"registration-service/internal/domain/user"
	"registration-service/internal/infrastructure/inputport/rest"
	"registration-service/internal/infrastructure/producer"
	"registration-service/internal/infrastructure/storage"
)

func NewApplication() *http.ServeMux {
	mux := http.NewServeMux()

	// Dependency Injection
	appProducers := applicationProducers{
		producer: producer.NewKafkaProducer([]string{"127.0.0.1:9092"}, "new-user"),
	}

	appRepositories := applicationRepositories{
		userRepository: storage.NewUserInMemoryRepository(),
	}

	appUseCases := applicationUseCases{
		addUser:     usecase.NewAddUserUseCase(appRepositories.userRepository, appProducers.producer),
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

type applicationProducers struct {
	producer service.Publisher
}
