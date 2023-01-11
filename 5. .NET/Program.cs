using Checkout;
using ServiceExtensions;

var builder = WebApplication.CreateBuilder(args);
builder.Logging.ClearProviders();
builder.Logging.AddConsole();
builder.Configuration.AddJsonFile("appsettings.json", optional: true, reloadOnChange: true);
builder.Services.AddControllersWithViews();
builder.Services.ConfigureCheckout(builder.Configuration);
var app = builder.Build();

if (!app.Environment.IsDevelopment())
{
    app.UseHttpsRedirection();
    app.UseHsts();
}

app.Services.GetService<CheckoutSdk>();

app.UseStaticFiles();
app.UseRouting();

app.MapControllerRoute(name: "payments",
    pattern: "payments",
    defaults: new { controller = "Checkout", action = "ExecutePaymentRequest" });

app.MapControllerRoute(
    name: "default",
    pattern: "{controller=Home}/{action=Index}/{id?}");

app.Run();
