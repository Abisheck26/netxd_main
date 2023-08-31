package main

import (
	"context"
	"fmt"

	"github.com/Abisheck26/netxd-grpc/config"
	controllers "github.com/Abisheck26/netxd-grpc/netxd_dal_controllers"
	netxd_constants "github.com/Abisheck26/netxd-grpc/constants"
	cus "github.com/Abisheck26/netxd-grpc/netxd_customer"
	//controllers "github.com/Abisheck26/netxd-grpc/netxd_dal/netxd_dal_controllers"
	services "github.com/Abisheck26/netxd-grpc/netxd_dal_services"

	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
)

func initDatabase(client *mongo.Client) {

	customerCollection := config.GetCollection(client, "bankdb", "customer")
	controllers.CustomerService = services.NewCustomerServiceInit(customerCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", netxd_constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	cus.RegisterCustomerServiceServer(s, &controllers.RPCServer{})

	fmt.Println("server running on port", netxd_constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}
