using Checkout;
using Checkout.Common;
using Checkout.Payments.Request;
using Checkout.Payments.Request.Source;
using CheckoutSample.Models;
using Microsoft.AspNetCore.Mvc;

namespace CheckoutSample.Controllers;

public class CheckoutController : Controller
{
    private readonly ICheckoutApi _checkoutApi;
    private readonly ILogger _log;

    public CheckoutController(ICheckoutApi checkoutApi, ILogger<CheckoutController> logger)
    {
        _checkoutApi = checkoutApi;
        _log = logger;
    }

    [HttpPost]
    public async Task<ActionResult<string>> ExecutePaymentRequest([FromBody] TokenRequest data)
    {
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
}
