package main

import (
	"context"
	"fmt"
	"log"

	cus "github.com/Abisheck26/netxd-grpc/netxd_customer"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := cus.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &cus.Customer{
		Customerid: 26,
		FirstName:  "Abisheck",
		LastName:   "G",
		BankId:     1,
		Balance:    1000.0,
		CreatedAt:  "",
		UpdatedAt:  "",
		IsActive:   true,
	})
	if err != nil {
		log.Fatalf("failed to call CreateCustomer: %v", err)
	}
	fmt.Printf("Response: %s\n", response)
}
