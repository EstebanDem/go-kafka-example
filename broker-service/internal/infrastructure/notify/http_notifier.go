package notify

import (
	"broker-service/internal/application/service"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type HttpNotifier struct {
	AppsToNotify []service.InterestedApp
}

func NewHttpNotifier(appsToNotify []service.InterestedApp) service.NotifyService {
	return &HttpNotifier{
		AppsToNotify: appsToNotify,
	}
}

func (h *HttpNotifier) Send(message any) {
	postBody, _ := json.Marshal(message)
	messageBuffer := bytes.NewBuffer(postBody)

	resultCh := make(chan NotificationResult)
	var wg sync.WaitGroup

	for _, appToNotify := range h.AppsToNotify {
		wg.Add(1)

		go func(app service.InterestedApp) {
			defer wg.Done()
			response, err := http.Post(app.Address, "application/json", messageBuffer)
			if response != nil {
				defer response.Body.Close()
			}
			if err != nil {
				resultCh <- NotificationResult{App: app, StatusCode: 0, Error: err}
				return
			}
			resultCh <- NotificationResult{App: app, StatusCode: response.StatusCode, Error: nil}

		}(appToNotify)

	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		if result.Error != nil {
			log.Printf("Error sending message to %s: %v\n", result.App.Name, result.Error)
		} else {
			log.Printf("Message sent to %s, status code: %d\n", result.App.Name, result.StatusCode)
		}
	}

}

func (h *HttpNotifier) GetInterestedApps() []service.InterestedApp {
	return h.AppsToNotify
}

type NotificationResult struct {
	App        service.InterestedApp
	StatusCode int
	Error      error
}
