package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/VeneLooool/fields-api/internal/app/api/v1/fields"
	"github.com/VeneLooool/fields-api/internal/config"
	pb "github.com/VeneLooool/fields-api/internal/pb/api/v1/fields"
	"github.com/VeneLooool/fields-api/internal/pkg/db"
	fields_repo "github.com/VeneLooool/fields-api/internal/repository/fields"
	fields_uc "github.com/VeneLooool/fields-api/internal/usecase/fields"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.New(ctx)
	if err != nil {
		log.Fatalf("failed to create new config: %s", err.Error())
	}

	go func() {
		if err := runGRPC(ctx, cfg); err != nil {
			log.Fatal(err)
		}
	}()

	if err := runHTTPGateway(ctx, cfg); err != nil {
		log.Fatal(err)
	}
}

func runGRPC(ctx context.Context, cfg *config.Config) error {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	dbAdapter, err := db.New(ctx)
	if err != nil {
		return err
	}
	defer dbAdapter.Close(ctx)

	fieldsServer, err := newServices(ctx, dbAdapter)
	if err != nil {
		return err
	}
	pb.RegisterFieldsServer(grpcServer, fieldsServer)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}
	
	log.Printf("gRPC server listening on :%s\n", cfg.GrpcPort)
	if err = grpcServer.Serve(grpcListener); err != nil {
		return err
	}
	return nil
}

func runHTTPGateway(ctx context.Context, cfg *config.Config) error {
	mux := runtime.NewServeMux()
	err := pb.RegisterFieldsHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%s", cfg.GrpcPort), []grpc.DialOption{
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

	withCORS := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Для preflight-запросов
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			h.ServeHTTP(w, r)
		})
	}

	// gRPC → REST mux
	http.Handle("/", withCORS(mux))

	log.Printf("HTTP gateway listening on :%s\n", cfg.HttpPort)
	if err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpPort), nil); err != nil {
		return err
	}

	return nil
}

func newServices(ctx context.Context, dbAdapter db.DataBase) (*fields.Implementation, error) {
	fieldsRepo := fields_repo.New(dbAdapter)
	fieldsUC := fields_uc.New(fieldsRepo)

	return fields.NewService(fieldsUC), nil
}
