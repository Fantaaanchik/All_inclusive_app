package server

func (h Handler) AllRoutes() {

	h.Engine.POST("/login", h.Login)
	h.Engine.POST("/register", h.Register)
	h.Engine.Group("/withAuth")
	h.Engine.Use(AuthMiddleware())

	h.Engine.GET("/get_all_users", h.GetAllUsers)
	h.Engine.GET("/get_user", h.GetUser)
	h.Engine.PUT("/update_user", h.UpdateUser)
	h.Engine.DELETE("/delete_user", h.DeleteUser)

}
