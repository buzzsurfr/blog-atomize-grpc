package main

import (
	"context"
	"fmt"
)

func PlaceOrder(ctx context.Context, customerID int, cents int) error {
	// Process the order itself
	fmt.Printf("Orders Backend: Order placed by customer %d for %d cents.\n", customerID, cents)

	// Add loyalty points based on the order's amount.
	// 1 LP per cent (USD)
	loyaltyPoints := cents
	if err := AddPoints(context.Background(), customerID, loyaltyPoints); err != nil {
		fmt.Printf("Orders Backend: Something happened while adding %d loyalty points for customer ID %d.\n", loyaltyPoints, customerID)
	} else {
		fmt.Printf("Orders Backend: %d loyalty points added to customer with ID %d\n", loyaltyPoints, customerID)
	}

	// Assuming the order went through even if the points weren't assigned
	return nil
}
