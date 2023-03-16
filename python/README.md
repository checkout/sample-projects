[![build-status](https://github.com/checkout/sdk-samples/actions/workflows/create_python_package.yml/badge.svg)](https://github.com/checkout/sdk-samples/actions/workflows/create_python_package.yml)
[![PyPI - latest](https://img.shields.io/pypi/v/checkout-sdk?label=latest&logo=pypi)](https://pypi.org/project/checkout-sdk)

This project is an example of how to start an integration with Python and [Checkout SDK](https://github.com/checkout/checkout-sdk-python)

This project uses `pip` as the package manager for all the requirements that a project needs, ensure that you have pip installed before running the project.

# :sparkles: Before Start

To remain comply with PCI regulations, you need to protect your card numbers, fortunately Checkout have [Frames Framework](https://www.checkout.com/docs/integrate/frames#Who_is_Frames_for?) that helps you to tokenize the payment card, this wrapper includes method `submitCard()`. In the below example, we call this when the "Pay Now" button is clicked.

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

```python
try:
    sdk = CheckoutSdk.builder() \
        .secret_key('sk_sbox_XXX') \
        .environment(environment=Environment.sandbox()) \
        .build()
except Exception as e:
    return json.dumps({'status': e.http_metadata.status_code, 'error-message': e.http_metadata.reason})
```

If you have any questions regarding SDK usage, please refer to SDK landing [page](https://github.com/checkout/checkout-sdk-python)

Then you need to build your request, in this case is a `payment request` with `token source` and then
just call the the SDK function to request a payment

```python
request_token = request.json

token_source = RequestTokenSource()
token_source.token = request_token['token']

payment_request = PaymentRequest()
payment_request.source = token_source
payment_request.amount = 2499
payment_request.currency = Currency.GBP
payment_request.processing_channel_id = 'pc_XXX'

response = sdk.payments.request_payment(payment_request)
return json.dumps(response.__dict__, default=lambda o: o.__dict__, indent=4)
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
python -m pip install --upgrade pip
pip install -r requirements-dev.txt
flask --app app run
```
