package controller;

import com.checkout.CheckoutApiException;
import com.checkout.CheckoutApiImpl;
import com.checkout.CheckoutAuthorizationException;
import com.checkout.CheckoutSdk;
import com.checkout.Environment;
import com.checkout.common.Currency;
import com.checkout.payments.request.PaymentRequest;
import com.checkout.payments.request.source.RequestTokenSource;
import com.checkout.payments.response.PaymentResponse;
import com.checkout.tokens.CardTokenResponse;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import errors.ErrorResponse;
import org.eclipse.jetty.http.HttpStatus;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import spark.Request;
import spark.Response;

import java.util.concurrent.ExecutionException;

import static java.util.Objects.requireNonNull;

public class PaymentsController {

    private PaymentsController() {
    }

    private static final Logger LOG = LoggerFactory.getLogger(PaymentsController.class);
    private static final Gson serializer = new GsonBuilder().create();

    public static String pay(Request req, Response res) {

        CheckoutApiImpl checkoutSdk;
        try {
            checkoutSdk = CheckoutSdk.builder()
                .staticKeys()
                .secretKey(requireNonNull("sk_sbox_XXX"))
                .environment(Environment.SANDBOX)
                .build();
        } catch (CheckoutAuthorizationException e) {
            LOG.error("An error occurred while initializing Checkout SDK - {}", e.getMessage());
            res.status(400);
            return serializer.toJson(new ErrorResponse(400, e.getMessage()));
        }

        CardTokenResponse cardToken = serializer.fromJson(req.body(), CardTokenResponse.class);

        RequestTokenSource tokenSource = RequestTokenSource.builder()
            .token(cardToken.getToken())
            .build();

        PaymentRequest request = PaymentRequest.builder()
            .source(tokenSource)
            .amount(2499L)
            .currency(Currency.GBP)
            .processingChannelId("pc_XXX")
            .build();

        PaymentResponse response;
        try {
            response = checkoutSdk.paymentsClient().requestPayment(request).get();
        } catch (InterruptedException | ExecutionException e) {
            LOG.error("An error occurred while processing payment request - {}", e.getMessage());
            CheckoutApiException chkException = parseException(e);
            res.status(chkException.getHttpStatusCode());
            return serializer.toJson(new ErrorResponse(chkException.getHttpStatusCode(), e.getMessage()));
        }

        return serializer.toJson(response);
    }

    private static CheckoutApiException parseException(Exception e) {
        if (e.getCause() instanceof CheckoutApiException) {
            return (CheckoutApiException) e.getCause();
        }
        return new CheckoutApiException(HttpStatus.BAD_REQUEST_400, null, null);
    }
}


