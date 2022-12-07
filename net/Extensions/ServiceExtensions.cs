using CheckoutSDK.Extensions.Configuration;

namespace ServiceExtensions
{
    public static class ServiceExtensions
    {
        public static IServiceCollection ConfigureCheckout(this IServiceCollection services, IConfiguration configuration)
        {
            return CheckoutServiceCollection.AddCheckoutSdk(services, configuration);
        }
    }
}
