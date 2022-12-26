
package controllers

import "../middleware"


func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	//s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")

	//products routes
	
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.GetProducts)).Methods("GET")
	
	s.Router.HandleFunc("/product/{id}", middlewares.SetMiddlewareJSON(s.GetProductById)).Methods("GET")
	
	s.Router.HandleFunc("/product/{id}", middlewares.SetMiddlewareJSON(s.UpdateProduct)).Methods("PUT")
	
	s.Router.HandleFunc("/product/{id}", s.DeleteProduct ).Methods("DELETE")
	
}