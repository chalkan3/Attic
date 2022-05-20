package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"rodrigues.igor.com/attic/internal/database"
	grpc_clients "rodrigues.igor.com/attic/internal/grpc"
	"rodrigues.igor.com/attic/internal/items"
	nats_cli "rodrigues.igor.com/attic/internal/nats"
	"rodrigues.igor.com/attic/pb"

	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		httpPort    = fs.String("http_port", "9005", "application http port default 9000")
		grpcPort    = fs.String("grpc_port", "50054", "application grpc port default 50051")
		serviceName = fs.String("service_name", "item_service", "service name")
	)

	flag.Usage = fs.Usage
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	logger := createLogger(*httpPort)

	_, err := nats_cli.NewNATS("", true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	gorm, err := database.NewPostgresGORM("host=localhost user=postgres password=postgres dbname=physical_environment port=5433 sslmode=disable").Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	items.NewMigrations(gorm).Migrate()

	grpcClients := grpc_clients.NewClientGRPC()

	svc := items.NewInstrumentingMiddleware(items.NewMetrics(), items.NewLogMW(logger, items.NewService(items.NewRepository(gorm), grpcClients)))
	routes := items.NewHTTPServer(svc, logger)
	grpcServer := items.NewGRPCServer(svc, logger)

	logger.Log(
		"service name", *serviceName,
		"msg", "HTTP",
		"addr", *httpPort,
		"prom-path", "/metrics")

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":"+*grpcPort)
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterItemServer(baseServer, grpcServer)
		baseServer.Serve(grpcListener)
	}()

	go func() {
		errs <- http.ListenAndServe(":"+*httpPort, routes)
	}()

	logger.Log("exit", <-errs)
}

func createLogger(port string) log.Logger {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)
	return logger
}
