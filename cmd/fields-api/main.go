package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"

	"github.com/VeneLooool/fields-api/internal/app/api/v1/fields"
	pb "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := runGRPC(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := runHTTPGateway(ctx); err != nil {
		log.Fatal(err)
	}
}

func runGRPC(_ context.Context) error {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	fieldsServer := fields.NewService()
	pb.RegisterFieldsServer(grpcServer, fieldsServer)

	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	log.Println("gRPC server listening on :50051")
	if err = grpcServer.Serve(grpcListener); err != nil {
		return err
	}
	return nil
}

func runHTTPGateway(ctx context.Context) error {
	mux := runtime.NewServeMux()
	err := pb.RegisterFieldsHandlerFromEndpoint(ctx, mux, "localhost:50051", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalf("failed to register gateway: %s", err.Error())
	}

	// Serve Swagger JSON and Swagger UI
	fs := http.FileServer(http.Dir("./swagger-ui")) // директория со статикой UI
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))

	// Serve Swagger JSON файл
	http.HandleFunc("/swagger/fields.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/pb/api/v1/fields/fields.swagger.json")
	})

	// gRPC → REST mux
	http.Handle("/", mux)

	log.Println("HTTP gateway listening on :8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
