This project is an example of how to start an integration with Ruby and [Checkout SDK](https://github.com/checkout/checkout-sdk-ruby)

This project uses `Bundler` to track and install the gems and versions needed.

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

```ruby
begin
  sdk = CheckoutSdk.builder
                   .static_keys
                   .with_secret_key('sk_sbox_XXX')
                   .with_environment(CheckoutSdk::Environment.sandbox)
                   .build
rescue => e
  logger.error "An error occurred while initializing Checkout SDK - #{e.message}"
  status 400
  body "An error occurred while initializing Checkout SDK - #{e.message}"
end
```

If you have any questions regarding SDK usage, please refer to SDK landing [page](https://github.com/checkout/checkout-sdk-ruby)

Then you need to build your request, in this case is a `payment request` with `token source` and then
just call the SDK function to request a payment

```ruby
payment_request = {
  source: {
    type: 'token',
    token: card_token['token']
  },
  amount: 2499,
  currency: CheckoutSdk::Common::Currency::GBP,
  processing_channel_id: 'pc_XXX'
} # Also available as CheckoutSdk::Payments::PaymentRequest

begin
  response = sdk.payments.request_payment payment_request
rescue => e
  logger.error "An error occurred while initializing Checkout SDK - #{e.message}"
  status 400
  body "An error occurred while requesting payment - #{e.message}"
end
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
bundle install
```
2. Run the project
```shell
bundle exec ruby server.rb
```
3. Go to http://localhost:4242
