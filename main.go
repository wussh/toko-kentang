package main

import (
	"github.com/wussh/tokokentang/config"
	cartData "github.com/wussh/tokokentang/features/cart/data"
	cartHandler "github.com/wussh/tokokentang/features/cart/handler"
	cartService "github.com/wussh/tokokentang/features/cart/services"
	productData "github.com/wussh/tokokentang/features/product/data"
	productHandler "github.com/wussh/tokokentang/features/product/handler"
	productService "github.com/wussh/tokokentang/features/product/services"

	// trxD "github.com/wussh/tokokentang/features/transaction/data"
	// trxH "github.com/wussh/tokokentang/features/transaction/handler"
	// trxS "github.com/wussh/tokokentang/features/transaction/services"

	// tdxD "github.com/wussh/tokokentang/features/transaction_detail/data"
	// tdxH "github.com/wussh/tokokentang/features/transaction_detail/handler"
	// tdxS "github.com/wussh/tokokentang/features/transaction_detail/services"
	"log"

	usrD "github.com/wussh/tokokentang/features/user/data"
	usrH "github.com/wussh/tokokentang/features/user/handler"
	usrS "github.com/wussh/tokokentang/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := usrD.New(db)
	userSrv := usrS.New(userData)
	userHdl := usrH.New(userSrv)

	productDt := productData.New(db)
	productSrv := productService.New(productDt)
	productHdl := productHandler.New(productSrv)

	cartDt := cartData.New(db)
	cartSrv := cartService.New(cartDt)
	cartHdl := cartHandler.New(cartSrv)

	// trxData := trxD.New(db)
	// trxSrv := trxS.New(trxData)
	// trxHdl := trxH.New(trxSrv)

	// trxDataa := tdxD.New(db)
	// trxService := tdxS.New(trxDataa)
	// trxHandler := tdxH.New(trxService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	user := e.Group("/users")

	user.GET("", userHdl.Profile(), middleware.JWT([]byte(config.JWTKey)))
	user.PUT("", userHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	user.DELETE("", userHdl.Deactivate(), middleware.JWT([]byte(config.JWTKey)))

	e.GET("/products/:id", productHdl.GetById())
	e.GET("/products", productHdl.GetAll())
	e.GET("/products/category/:category", productHdl.GetAllByCategory())
	e.POST("/products", productHdl.Add(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/products/:id", productHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/products/:id", productHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))

	e.POST("/carts/:idProduct", cartHdl.Add(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/carts/:idCart", cartHdl.GetByIdC(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/carts", cartHdl.GetByIdU(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/checkout/:idCart", cartHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/carts/:idCart", cartHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))

	// e.POST("/transactions/:id", trxHdl.Add(), middleware.JWT([]byte(config.JWTKey)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
