using Microsoft.AspNetCore.Mvc;

namespace CheckoutSample.Controllers
{
    public class HomeController : Controller
    {
        public IActionResult Index()
        {
            return View();
        }
    }
}
