// github.com/checkout/checkout-sdk-go v1.0.5

package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	payments "github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/payments/nas/sources"
	"github.com/checkout/checkout-sdk-go/tokens"
)

type (
	PaymentsController struct{}
)

func NewPaymentsController() *PaymentsController {
	return &PaymentsController{}
}

func (c *PaymentsController) RequestPayment(w http.ResponseWriter, req *http.Request) {
	sdk, err := checkout.Builder().
		StaticKeys().
		WithEnvironment(configuration.Sandbox()).
		WithSecretKey("sk_sbox_XXX").
		Build()
	if err != nil {
		log.Printf("An error occurred while initializing Checkout SDK - %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cardToken tokens.CardTokenResponse
	if err = json.NewDecoder(req.Body).Decode(&cardToken); err != nil {
		log.Printf("An error occurred while deserealizing card token - %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenSource := sources.NewRequestTokenSource()
	tokenSource.Token = cardToken.Token

	paymentRequest := payments.PaymentRequest{
		Source:              tokenSource,
		Amount:              2499,
		Currency:            common.GBP,
		ProcessingChannelId: "pc_XXX",
	}

	response, err := sdk.Payments.RequestPayment(paymentRequest, nil)
	if err != nil {
		ckoErr := err.(errors.CheckoutAPIError)
		log.Printf("An error occurred while requesting payment - %s", ckoErr.Status)
		http.Error(w, ckoErr.Status, ckoErr.StatusCode)
		return
	}

	handleResponse(w, response.HttpMetadata.StatusCode, response)
}

func handleResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
