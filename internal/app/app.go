package app

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"github.com/igdanit/sso/internal/grpc/auth"
	pb "github.com/igdanit/sso/protos/gen/go/auth"
)

type App struct {
	port   int
	logger *slog.Logger
	server *grpc.Server
}

// Run create and start GRCP server
func (app *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return err
	}
	pb.RegisterAuthServer(app.server, &auth.AuthService{})
	err = app.server.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func (app *App) Stop() error {
	app.server.GracefulStop()
	return nil
}

func New(port int, logger *slog.Logger, server *grpc.Server) *App {
	app := App{
		port:   port,
		logger: logger,
		server: server,
	}
	return &app
}
