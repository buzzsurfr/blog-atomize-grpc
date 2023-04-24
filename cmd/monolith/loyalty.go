package main

import (
	"context"
	"fmt"
)

func GetPointsForCustomer(ctx context.Context, customerID int) (int, error) {
	// Mock that any customer has 100 loyalty points.
	loyaltyPoints := 100
	fmt.Printf("Loyalty Backend: Customer ID %d has %d loyalty points.\n", customerID, loyaltyPoints)

	return loyaltyPoints, nil
}

func AddPoints(ctx context.Context, customerID int, loyaltyPoints int) error {
	fmt.Printf("Loyalty Backend: Added %d loyalty points to customer ID %d\n", loyaltyPoints, customerID)

	return nil
}
