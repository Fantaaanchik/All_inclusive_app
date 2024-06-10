package server

func (h Handler) AllRoutes() {
	h.Engine.POST("/login", h.Login)
	h.Engine.POST("/register", h.Register)
	h.Engine.GET("/get_all_users", h.GetAllUsers)
	h.Engine.GET("/get_user/:id", h.GetUser)
	h.Engine.PUT("/update_user/:id", h.UpdateUser)
	h.Engine.DELETE("/delete_user/:id", h.DeleteUser)

}
