package application

import (
	"net/http"

	"github.com/checkout/sdk-samples/go/src/api/controllers"
)

type (
	paymentsController interface {
		RequestPayment(http.ResponseWriter, *http.Request)
	}

	Application struct {
		PaymentsController paymentsController
	}
)

func BuildApplication() *Application {
	return &Application{
		PaymentsController: controllers.NewPaymentsController(),
	}
}
