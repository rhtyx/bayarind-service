package controller

import (
	"github.com/rhtyx/bayarind-service.git/model"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	authorService  model.AuthorService
	bookService    model.BookService
	userService    model.UserService
	sessionService model.SessionService
}

func NewController() *Controller {
	return new(Controller)
}

func (c *Controller) RegisterAuthorService(authorService model.AuthorService) {
	c.authorService = authorService
}

func (c *Controller) RegisterBookService(bookService model.BookService) {
	c.bookService = bookService
}

func (c *Controller) RegisterUserService(userService model.UserService) {
	c.userService = userService
}

func (c *Controller) RegisterSessionService(sessionService model.SessionService) {
	c.sessionService = sessionService
}

func (c Controller) InitRoutes(route *echo.Echo) {
	r := route.Group("/api/v1")
	r.Use(HmacMiddleware)

	user := r.Group("/users", JwtMiddleware)
	user.GET("/", c.FindUserByID)
	user.PUT("/", c.UpdateUser)
	user.DELETE("/", c.DeleteUser)

	book := r.Group("/books", JwtMiddleware)
	book.POST("/", c.CreateBook)
	book.GET("/:id/", c.FindBookByID)
	book.GET("/", c.FindAllBooks)
	book.PUT("/:id/", c.UpdateBook)
	book.DELETE("/:id/", c.DeleteBook)

	author := r.Group("/authors", JwtMiddleware)
	author.POST("/", c.CreateAuthor)
	author.GET("/:id/", c.FindAuthorByID)
	author.GET("/", c.FindAllAuthors)
	author.PUT("/:id/", c.UpdateAuthor)
	author.DELETE("/:id/", c.DeleteAuthor)

	auth := r.Group("/auth")
	auth.POST("/login/", c.Login)
	auth.POST("/logout/", c.Logout, JwtMiddleware)
	auth.POST("/signup/", c.CreateUser)
	auth.POST("/refresh/", c.RefreshAccessToken)
}
