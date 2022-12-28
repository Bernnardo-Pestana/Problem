
package controllers

import "github.com/Bernnardo-Pestana/Problem/api/middleware"


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
	
	s.Router.HandleFunc("/product/{id}",middlewares.SetMiddlewareJSON(s.DeleteProduct) ).Methods("DELETE")

		//purchases routes
	
		s.Router.HandleFunc("/purchases", middlewares.SetMiddlewareJSON(s.CreatePurchase)).Methods("POST")
		s.Router.HandleFunc("/purchases", middlewares.SetMiddlewareJSON(s.GetPurchases)).Methods("GET")
		
		s.Router.HandleFunc("/purchases/{id}", middlewares.SetMiddlewareJSON(s.GetPurchaseById)).Methods("GET")
		
		s.Router.HandleFunc("/purchases/{id}", middlewares.SetMiddlewareJSON(s.UpdatePurchase)).Methods("PUT")
		
		s.Router.HandleFunc("/purchases/{id}",middlewares.SetMiddlewareJSON(s.DeletePurchase) ).Methods("DELETE")
		
	
}