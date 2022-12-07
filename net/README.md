This project is an example of how to start an integration with .Net and [Checkout SDK](https://github.com/checkout/checkout-sdk-net)

This project uses `Microsoft.AspNetCore.App` and `net6.0` and `NuGetPackages` as the package manager for all the requirements that a project needs.

# :sparkles: Before Start

To remain comply with PCI regulations, you need to protect your card numbers, fortunately Checkout have [Frames Framework](https://www.checkout.com/docs/integrate/frames#Who_is_Frames_for?)
that helps you tokenize the payment card, this wrapper includes method `submitCard()`. In the below example, we call this when the "Pay Now" button is clicked.

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

Then we intercept the `event` and let Frames tokenize the card

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

Once that Frames validates and returns the tokenized token, then you can add another event to handle the request to the backed. If you have any questions regarding the Frames events you can visit the [documentation page](https://www.checkout.com/docs/integrate/frames/frames-reference)

On the backed you need to initiate the SDK with the proper `Secret Key` as follows in `appsettings.json`:

```json
"Checkout": {
    "SecretKey": "sk_sbox_XXX",
    "Environment": "Sandbox",
    "PlatformType": "Default"
  }
```

If you have any questions regarding SDK usage, please refer to SDK landing [page](https://github.com/checkout/checkout-sdk-net)

Then you need to build your request, in this case is a `payment request` with `token source` and then
just call the the SDK function to request a payment

```cs
try
{
    PaymentRequest request = new()
    {
        Source = new RequestTokenSource { Token = data.Token },
        Amount = 2499,
        Currency = Currency.GBP
    };

    var response = await _checkoutApi.PaymentsClient().RequestPayment(request);
    return response.Body;
}
catch (CheckoutApiException e)
{
    IDictionary<string, object> errorDetails = e.ErrorDetails;
    _log.LogError("Error while processing request = {error}", errorDetails["error_codes"]);
    return BadRequest();
}
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
1. Restore the project dependencies
```shell
dotnet restore
```
2. Build the project
```shell
dotnet build
```
3. Run the project
```shell
dotnet run
```
