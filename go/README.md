[![build-status](https://github.com/checkout/sample-projects/actions/workflows/create_go_package.yml/badge.svg)](https://github.com/checkout/sample-projects/actions/workflows/create_go_package.yml)
[![GitHub release](https://img.shields.io/github/release/checkout/checkout-sdk-go.svg)](https://github.com/checkout/checkout-sdk-go/releases/)

This project is an example of how to start an integration with GO and [Checkout SDK](https://github.com/checkout/checkout-sdk-go)

This project uses `Go Modules` as the dependencies manager for all the requirements that a project needs,
ensure that you have Go Modules installed before running the project.

# :sparkles: Before Start

To remain comply with PCI regulations, you need to protect your card numbers, fortunately Checkout have [Frames Framework](https://www.checkout.com/docs/integrate/frames#Who_is_Frames_for?)
that helps you to tokenize the payment card, this wrapper includes method `submitCard()`. In the below example, we call this when the "Pay Now" button is clicked.

````html
<form id="payment-form" method="POST">
    <div class="one-liner">
        <div class="card-frame"></div>
        <button id="pay-button" disabled>
            PAY GBP 24.99
        </button>
    </div>
    <p class="error-message"></p>
    <p class="payment-message"></p>
</form>

<script src="https://cdn.checkout.com/js/framesv2.min.js"></script>
````

Then we intercept the `event` and let Frames to tokenize the card

````javascript
form.addEventListener('submit', function (event) {
    event.preventDefault();
    Frames.submitCard();
});
````

Make sure to provide your correct `Public Key` in `script.js`, which is the key that authorizes to you the access to Checkout API's

````javascript
Frames.init('pk_sbox_XXX');
````

Once that Frames validates and returns the tokenized token, then you can add another event to handle the request
to the backed. If you have any questions regarding the Frames events you can visit the [documentation page](https://www.checkout.com/docs/integrate/frames/frames-reference)

On the backend you need to initiate the SDK with the proper `Secret Key` as follows:

```go
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
```

If you have any questions regarding SDK usage, please refer to SDK landing [page](https://github.com/checkout/checkout-sdk-go)

Then you need to build your request, in this case is a `payment request` with `token source` and then
just call the SDK function to request a payment

```go
tokenSource := sources.NewRequestTokenSource()
tokenSource.Token = cardToken.Token

paymentRequest := payments.PaymentRequest{
    Source:   tokenSource,
    Amount:   2499,
    Currency: common.GBP,
    ProcessingChannelId: "pc_XXX",
}

response, err := sdk.Payments.RequestPayment(paymentRequest, nil)
if err != nil {
    ckoErr := err.(errors.CheckoutAPIError)
    log.Printf("An error occurred while requesting payment - %s", ckoErr.Status)
    http.Error(w, ckoErr.Status, ckoErr.StatusCode)
    return
}
```

And that's it! Your payment has been processed.

### :book: Checkout our official documentation.

* [Official Docs (Default)](https://docs.checkout.com/)
* [Official Docs (Previous)](https://docs.checkout.com/previous)

### :books: Check out our official API documentation guide, where you can also find more usage examples.

* [API Reference (Default)](https://api-reference.checkout.com/)
* [API Reference (Previous)](https://api-reference.checkout.com/previous)


# :rocket: Run the project and test it out

1. Build the project
```shell
go mod tidy

go build
```
2. Run the project
```shell
go run server.go
```
3. Go to http://localhost:4242
