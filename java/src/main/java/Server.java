import static spark.Spark.port;
import static spark.Spark.post;
import static spark.Spark.staticFiles;

import controller.PaymentsController;

public class Server {
    public static void main(String[] args) {
        port(4242);
        staticFiles.location("/public");
        post("/payments", PaymentsController::pay);
    }
}
