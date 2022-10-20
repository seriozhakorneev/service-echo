package v1

// NewRouter -.
// Swagger spec:
// @title       Service Math API
// @description Service doing math
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
//func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Basic) {
//	// Options
//	handler.Use(gin.Logger())
//	handler.Use(gin.Recovery())
//
//	// Swagger
//	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
//	handler.GET("/swagger/*any", swaggerHandler)
//
//	// K8s probe
//	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
//
//	// Routers
//	h := handler.Group("/v1")
//	{
//	}
//}
