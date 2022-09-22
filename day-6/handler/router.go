package handler

//type Router struct {
//	Controller *Controller
//	Router     *echo.Echo
//	SecretKey  []byte
//}

//func New(c *Controller) *Router {
//	e := echo.New()

//jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
//	Claims:     &handler.Claims{},
//	SigningKey: secretKey,
//})

//e.GET("/", func(c echo.Context) error {
//	return c.String(http.StatusOK, "Hello, World!")
//})

// v1
//	v1 := e.Group("/v1")
//	// auth router
//	//v1.POST("/auth/login", r.Controller.LoginController(secretKey))
//	bookRouter(v1)
//	return router
//}
