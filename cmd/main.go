package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/config"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/endpoint"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/service/models"
	transportgrpc "github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/grpc"
	"github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/grpc/pb"
	transporthttp "github.com/mahdi-asadzadeh/go-kit-accounts/pkg/transport/http"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ************ config ************
	mode := flag.String("mode", "dev", "Mode config envs for run project.")
	flag.Parse()
	config.LoadSettings(*mode)

	logger := InitLogger(os.Getenv("PORT"))
	openDB := InitDB(os.Getenv("DB_URL"))
	Migrate(openDB)

	// ************ Service ************
	userSer := service.NewUserService(os.Getenv("JWT_SECRET"), openDB)

	// ************ Endpoint ************
	userEnd := endpoint.NewUserEndpoint(userSer)

	// ************ Transport ************
	errChn := make(chan error)
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errChn <- fmt.Errorf("%s", <-ch)
	}()

	// http
	go func() {
		httpServer := transporthttp.NewHttpUserServer(userEnd, logger)
		logger.Log("msg", "HTTP", "addr", os.Getenv("HTTP_PORT"))
		errChn <- http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), httpServer)
	}()

	// grpc
	go func() {
		grpcListener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
		defer grpcListener.Close()
		if err != nil {
			logger.Log("during", "Listen", "err", err)
			errChn <- err
		}
		grpcServer := transportgrpc.NewGrpcUserServer(userEnd, logger)
		grpcBaseServer := grpc.NewServer()
		pb.RegisterUSerServiceServer(grpcBaseServer, grpcServer)
		logger.Log("msg", "gRPC", "addr", os.Getenv("GRPC_PORT"))
		errChn <- grpcBaseServer.Serve(grpcListener)
	}()

	logger.Log("exit", <-errChn)

}

func InitLogger(port string) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)
	return logger
}

func InitDB(dbUrl string) *gorm.DB {
	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
		os.Exit(-1)
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&models.User{})
}
