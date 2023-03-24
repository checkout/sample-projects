# frozen_string_literal: true
# 'checkout_sdk', '~> 1.1', '>= 1.1.1'

require 'checkout_sdk'
require 'sinatra'

set :root, File.dirname(__FILE__)
set :public_folder, -> { File.join(root, 'resources') }
set :static, true
set :port, 4242

get '/' do
  redirect '/index.html'
end

post '/payments' do
  content_type 'application/json'

  begin
    sdk = CheckoutSdk.builder
                     .static_keys
                     .with_secret_key('sk_sbox_XXX')
                     .with_environment(CheckoutSdk::Environment.sandbox)
                     .build
  rescue StandardError => e
    logger.error "An error occurred while initializing Checkout SDK - #{e.message}"
    status 400
    body "An error occurred while initializing Checkout SDK - #{e.message}"
  end

  card_token = JSON.parse(request.body.read)

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
  rescue StandardError => e
    logger.error "An error occurred while initializing Checkout SDK - #{e.message}"
    status 400
    body "An error occurred while requesting payment - #{e.message}"
  end

  serialize(response).to_json
end

def serialize(object, hash = {})
  case object
  when OpenStruct then
    object.each_pair do |key, value|
      hash[key] = serialize(value)
    end
    hash
  when Array then
    object.map { |v| serialize(v) }
  else object
  end
end
