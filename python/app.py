# checkout-sdk>=3.0.10

import json
import os

from flask import Flask, render_template, request

from checkout_sdk.checkout_sdk import CheckoutSdk
from checkout_sdk.environment import Environment
from checkout_sdk.common.enums import Currency
from checkout_sdk.payments.payments import PaymentRequest
from checkout_sdk.payments.payments_previous import RequestTokenSource


app = Flask(__name__)


@app.route('/')
def render():
    return render_template('index.html')


@app.route('/pay', methods=['POST'])
def request_payment():
    try:
        sdk = CheckoutSdk.builder() \
                         .secret_key('sk_sbox_XXX') \
                         .environment(environment=Environment.sandbox()) \
                         .build()
    except Exception as e:
        return json.dumps({'status': e.http_metadata.status_code, 'error-message': e.http_metadata.reason})

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


if __name__ == '__main__':
    app.run()

