<?php

require '../vendor/autoload.php';

use Checkout\CheckoutApiException;
use Checkout\CheckoutException;
use Checkout\CheckoutSdk;
use Checkout\Common\Currency;
use Checkout\Environment;
use Checkout\Payments\Request\PaymentRequest;
use Checkout\Payments\Request\Source\RequestTokenSource;
use Monolog\Handler\StreamHandler;
use Monolog\Logger;

$log = new Logger("checkout-sdk-php-sample");
$log->pushHandler(new StreamHandler("php://stdout"));

try {
    $api = CheckoutSdk::builder()->staticKeys()
        ->environment(Environment::sandbox())
        ->secretKey("sk_sbox_XXX")
        ->build();
} catch (CheckoutException $e) {
    $log->error("An exception occurred while initializing Checkout SDK : {$e->getMessage()}");
    http_response_code(400);
}

$postData = file_get_contents("php://input");
$request = json_decode($postData);


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
