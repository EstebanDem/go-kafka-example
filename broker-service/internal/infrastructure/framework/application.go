package framework

import (
	"broker-service/internal/application/service"
	"broker-service/internal/infrastructure/consumer"
	"broker-service/internal/infrastructure/notify"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	kafkaBrokerAddress = "127.0.0.1:9092"
	emailTopic         = "email-validation"
	newUserTopic       = "new-user"
)

func RunBroker() {
	// Notify new user
	httpNewUserNotifier := notify.NewHttpNotifier(appsToNotifyNewUserRegistered())
	chMsg := make(chan service.ConsumerMessage)
	chErr := make(chan error)
	newUserMessageConstructor := func() service.ConsumerMessage {
		return &service.NewUserMessage{}
	}
	kafkaConsumer := consumer.NewKafkaConsumer([]string{kafkaBrokerAddress}, newUserTopic)

	go func() {
		log.Println("listening messages in 'new-user' topic")
		kafkaConsumer.Read(context.Background(), chMsg, chErr, newUserMessageConstructor)
	}()

	// Notify Email
	httpEmailNotifier := notify.NewHttpNotifier(appsToNotifyEmailRegistered())
	chEmailMsg := make(chan service.ConsumerMessage)
	chEmailErr := make(chan error)
	emailMessageConstructor := func() service.ConsumerMessage {
		return &service.EmailMessage{}
	}
	kafkaEmailConsumer := consumer.NewKafkaConsumer([]string{kafkaBrokerAddress}, emailTopic)

	go func() {
		log.Println("listening messages in 'email-validation' topic")
		kafkaEmailConsumer.Read(context.Background(), chEmailMsg, chEmailErr, emailMessageConstructor)
	}()

	// Setup signal handling
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer close(chMsg)
	defer close(chErr)
	defer close(chEmailMsg)
	defer close(chEmailErr)

	for {
		select {
		case <-quit:
			log.Println("broker-service going down")
			return
		case m := <-chMsg:
			log.Println("new message in 'new-user' topic")
			httpNewUserNotifier.Send(m)
		case err := <-chErr:
			log.Println(err)
		case m := <-chEmailMsg:
			log.Println("new message in 'email-validated' topic")
			httpEmailNotifier.Send(m)
		case err := <-chEmailErr:
			log.Println(err)
		}
	}
}

func appsToNotifyNewUserRegistered() []service.InterestedApp {
	return []service.InterestedApp{
		{
			Name:    "email-service",
			Address: "http://localhost:9977/validate",
		},
	}
}

func appsToNotifyEmailRegistered() []service.InterestedApp {
	return []service.InterestedApp{
		{
			Name:    "registration-service",
			Address: "http://localhost:9099/users/validate_email",
		},
	}
}
