This project is an example of how to start an integration with PHP and [Checkout SDK](https://github.com/checkout/checkout-sdk-php)

This project uses `composer` as the package manager for all the requirements that a project needs,
ensure that you have composer installed before running the project.

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

On the backed you need to initiate the SDK with the proper `Secret Key` as follows:

```php
try {
    $api = CheckoutSdk::builder()->staticKeys()
        ->environment(Environment::sandbox())
        ->secretKey("sk_sbox_XXX")
        ->build();
} catch (CheckoutException $e) {
    $log->error("An exception occurred while initializing Checkout SDK : {$e->getMessage()}");
    http_response_code(400);
}
```

If you have any questions regarding SDK usage, please refer to SDK landing [page](https://github.com/checkout/checkout-sdk-php)

Then you need to build your request, in this case is a `payment request` with `token source` and then
just call the the SDK function to request a payment

```php
$requestTokenSource = new RequestTokenSource();
$requestTokenSource->token = $request->token;

$request = new PaymentRequest();
$request->source = $requestTokenSource;
$request->currency = Currency::$GBP;
$request->amount = 2499;

try {
    echo json_encode($api->getPaymentsClient()->requestPayment($request));
} catch (CheckoutApiException $e) {
    $log->error("An exception occurred while processing payment request");
    http_response_code(400);
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

```shell
composer install
composer start
```
