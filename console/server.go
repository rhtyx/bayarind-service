package console

import (
	"errors"
	"os"
	"os/signal"

	"github.com/rhtyx/bayarind-service.git/config"
	"github.com/rhtyx/bayarind-service.git/controller"
	"github.com/rhtyx/bayarind-service.git/db"
	"github.com/rhtyx/bayarind-service.git/repository"
	"github.com/rhtyx/bayarind-service.git/service"
	"github.com/rhtyx/bayarind-service.git/token"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Run: run,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func run(cmd *cobra.Command, _ []string) {
	config.GetConfig()
	db.InitPostgresDB()
	token.InitJWT()
	token.InitHMAC()

	authorRepository := repository.NewAuthorRepository(db.PostgresDB)
	bookRepository := repository.NewBookRepository(db.PostgresDB)
	userRepository := repository.NewUserRepository(db.PostgresDB)
	sessionRepository := repository.NewSessionRepository(db.PostgresDB)

	authorService := service.NewAuthorService(authorRepository)
	bookService := service.NewBookService(bookRepository, authorRepository)
	userService := service.NewUserService(userRepository)
	sessionService := service.NewSessionService(sessionRepository, userRepository, token.Jwt)

	ctrl := controller.NewController()
	ctrl.RegisterAuthorService(authorService)
	ctrl.RegisterBookService(bookService)
	ctrl.RegisterUserService(userService)
	ctrl.RegisterSessionService(sessionService)

	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		<-sigCh
		errCh <- errors.New("Received an interrupt")
	}()

	go runHTTPServer(ctrl, errCh)
	log.Error(<-errCh)
}

func runHTTPServer(ctrl *controller.Controller, errCh chan<- error) {
	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	ctrl.InitRoutes(e)
	errCh <- e.Start(":" + config.Port())
}
