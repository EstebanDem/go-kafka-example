package framework

import (
	"email-service/internal/application/service"
	"email-service/internal/application/usecase"
	"email-service/internal/infrastructure/inputport/rest"
	"email-service/internal/infrastructure/producer"
	"net/http"
)

const (
	kafkaBrokerAddress = "127.0.0.1:9092"
	topic              = "email-validation"
)

func NewApplication() *http.ServeMux {
	mux := http.NewServeMux()

	appProducers := applicationProducers{producer: producer.NewKafkaProducer(
		[]string{kafkaBrokerAddress}, topic),
	}

	appUseCases := applicationUseCases{validateEmail: usecase.NewEmailValidator(appProducers.producer)}

	mux.HandleFunc("POST /validate", rest.ValidateEmailHandler(appUseCases.validateEmail))

	return mux
}

type applicationUseCases struct {
	validateEmail usecase.ValidateEmailUseCase
}

type applicationProducers struct {
	producer service.Publisher
}
