package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	lpb "github.com/buzzsurfr/blog-atomize-grpc/loyalty"
	opb "github.com/buzzsurfr/blog-atomize-grpc/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	loyaltyClient lpb.PointsClient
	ordersClient  opb.OrderClient
)

func handleLoyaltyPoints(w http.ResponseWriter, r *http.Request) {
	// All a customer can do is get loyalty points, add isn't visible from the API layer
	// GET uses the next path parameter as the customer ID. To get the loyalty points for customer 1234, use
	// Example: /loyalty/1234

	// Get customer ID form path
	var customerID int
	idString := strings.Split(r.URL.Path, "/")[2]
	// Error if customer ID is not present
	if idString == "" {
		fmt.Fprint(w, "Missing customer ID\n")
	}
	if intVar, err := strconv.Atoi(idString); err != nil {
		fmt.Fprintf(w, "Atoi() err: %v\n", err)
	} else {
		customerID = intVar
	}

	if pointsReply, err := loyaltyClient.GetPointsForCustomer(context.Background(), &lpb.GetPointsRequest{CustomerID: int32(customerID)}); err != nil {
		fmt.Printf("Loyalty Points Handler: GetPointsForCustomer() err: %v\n", err)
		fmt.Fprintf(w, "GetPointsForCustomer() err: %v\n", err)
	} else {
		fmt.Printf("Loyalty Points Handler: Customer %d has %d loyalty points.\n", customerID, pointsReply.Points)
		fmt.Fprintf(w, "Customer %d has %d loyalty points.\n", customerID, pointsReply.Points)
	}

}

func handleOrders(w http.ResponseWriter, r *http.Request) {

	// Get customer ID and amount from form data
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Get customer ID and amount from form data. Convert amount to cents
	var customerID, cents int
	if intVar, err := strconv.Atoi(r.PostForm.Get("customerID")); err != nil {
		fmt.Fprintf(w, "Atoi() err: %v\n", err)
	} else {
		customerID = intVar
	}
	if floatVar, err := strconv.ParseFloat(r.PostForm.Get("amount"), 64); err != nil {
		fmt.Fprintf(w, "ParseFloat() err: %v\n", err)
	} else {
		// Convert amount (in USD/float) to cents (int)
		cents = int(math.Round(floatVar * 100.0))
	}

	// ...

	if _, err := ordersClient.PlaceOrder(context.Background(), &opb.PlaceOrderRequest{CustomerID: int32(customerID), Cents: int32(cents)}); err != nil {
		fmt.Printf("Orders Handler: Unable to place order: %v\n", err)
		fmt.Fprintf(w, "Unable to place order: %v\n", err)
	} else {
		fmt.Printf("Orders Handler: Order for customer with ID %d placed successfully!\n", customerID)
		fmt.Fprintf(w, "Order for customer with ID %d placed successfully!\n", customerID)
	}
}

func main() {
	// Connect to Loyalty service
	loyaltyConn, loyaltyErr := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if loyaltyErr != nil {
		log.Fatalf("did not connect to loyalty: %v", loyaltyErr)
	}
	defer loyaltyConn.Close()
	loyaltyClient = lpb.NewPointsClient(loyaltyConn)

	// Connect to Orders service
	ordersConn, ordersErr := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if ordersErr != nil {
		log.Fatalf("did not connect to orders: %v", ordersErr)
	}
	defer ordersConn.Close()
	ordersClient = opb.NewOrderClient(ordersConn)

	fmt.Println("Starting server...")

	http.HandleFunc("/loyalty/", handleLoyaltyPoints)
	http.HandleFunc("/orders/", handleOrders)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print(err)
	}
}
