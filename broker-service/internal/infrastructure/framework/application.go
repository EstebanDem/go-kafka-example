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

func RunBroker() {
	// Notify new user
	httpNewUserNotifier := notify.NewHttpNotifier(appsToNotifyNewUserRegistered())
	chMsg := make(chan service.NewUserMessage)
	chErr := make(chan error)
	kafkaConsumer := consumer.NewKafkaConsumer([]string{"127.0.0.1:9092"}, "new-user")

	go func() {
		log.Println("listening messages in new-user topic")
		kafkaConsumer.Read(context.Background(), chMsg, chErr)
	}()

	// Notify Email
	httpEmailNotifier := notify.NewHttpNotifier(appsToNotifyEmailRegistered())
	chEmailMsg := make(chan service.EmailMessage)
	chEmailErr := make(chan error)
	kafkaEmailConsumer := consumer.NewKafkaEmailConsumer([]string{"127.0.0.1:9092"}, "email-validation")

	go func() {
		log.Println("listening messages in email-validation topic")
		kafkaEmailConsumer.Read(context.Background(), chEmailMsg, chEmailErr)
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
			httpNewUserNotifier.Send(m)
		case err := <-chErr:
			log.Println(err)
		case m := <-chEmailMsg:
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
