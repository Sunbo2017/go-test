package main

import (
    "log"
    "net"
    "strings"

    "golang.org/x/net/context"
    "google.golang.org/grpc"

    pb "go-test/rpc-protobuf/customer"
)

const (
    port = ":50051"
)

// server is used to implement customer.CustomerServer.
type server struct {
    savedCustomers []*pb.CustomerRequest
}

// CreateCustomer creates a new Customer
func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
    log.Println("run in server.CreateCustomer")
    s.savedCustomers = append(s.savedCustomers, in)
    log.Printf("server.customers: %v", s.savedCustomers)
    return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
    log.Println("run in server.GetCustomers")
    for _, customer := range s.savedCustomers {
        if filter.Keyword != "" {
            log.Printf("filter.key: %v", filter.Keyword)
            if !strings.Contains(customer.Name, filter.Keyword) {
                continue
            }
        }
        if err := stream.Send(customer); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    //Create a new grpc server
    s := grpc.NewServer()
    pb.RegisterCustomerServer(s, &server{})
    s.Serve(lis)
}