package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"otel-demo/otel"
)

func main() {
	otelEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if otelEndpoint == "" {
		otelEndpoint = "localhost:4317"
	}

	cleanup, err := otel.InitTelemetry(otelEndpoint)
	if err != nil {
		log.Fatalf("failed to init telemetry: %v", err)
	}
	defer cleanup()

	r := chi.NewRouter()

	// Standard middleware
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// ---------------- Routes ----------------
	// ================= USER =================
	r.Route("/user", userRoutes)

	// ================= ORDER =================
	r.Route("/order", orderRoutes)

	// ================= PRODUCT =================
	r.Route("/product", productRoutes)

	// ================= PAYMENT =================
	r.Route("/payment", paymentRoutes)

	// Root
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, OpenTelemetry!")
	})

	// otel wrap once
	handler := otelhttp.NewHandler(r, "chi-server")
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// ---------- Grouped logically ----------
func userRoutes(r chi.Router) {
	r.Get("/list", dummyHandler("User List"))
	r.Get("/profile", dummyHandler("User Profile"))
	r.Post("/create", dummyHandler("Create User"))
}

func orderRoutes(r chi.Router) {
	r.Get("/list", dummyHandler("Order List"))
	r.Get("/detail", dummyHandler("Order Detail"))
	r.Post("/create", dummyHandler("Create Order"))
}

func productRoutes(r chi.Router) {
	r.Get("/list", dummyHandler("Product List"))
	r.Get("/detail", dummyHandler("Product Detail"))
}

func paymentRoutes(r chi.Router) {

	r.Get("/list", dummyHandler("Payment List"))
	r.Get("/detail", dummyHandler("Payment Detail"))
	r.Post("/create", dummyHandler("Create Payment"))
}

func dummyHandler(message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span := trace.SpanFromContext(r.Context())
		traceID := span.SpanContext().TraceID().String()

		log.Printf(
			`{"msg":"api_called","trace_id":"%s","path":"%s"}`,
			traceID,
			r.URL.Path,
		)

		fmt.Fprintf(w, "%s endpoint called!\n", message)
	}
}
