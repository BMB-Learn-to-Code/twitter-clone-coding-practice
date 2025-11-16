package router

import "net/http"

// Router defines the interface for HTTP routing functionality.
// This abstraction allows us to decouple from specific router implementations
// like chi, gorilla/mux, or gin, making the code more testable and flexible.
type Router interface {
	// Use adds middleware to the router
	Use(middlewares ...func(http.Handler) http.Handler)

	// Route creates a new sub-router with the given pattern
	Route(pattern string, fn func(Router)) Router

	// Get registers a GET route with the given pattern and handler
	Get(pattern string, handlerFn http.HandlerFunc)

	// Post registers a POST route with the given pattern and handler
	Post(pattern string, handlerFn http.HandlerFunc)

	// Put registers a PUT route with the given pattern and handler
	Put(pattern string, handlerFn http.HandlerFunc)

	// Delete registers a DELETE route with the given pattern and handler
	Delete(pattern string, handlerFn http.HandlerFunc)

	// Patch registers a PATCH route with the given pattern and handler
	Patch(pattern string, handlerFn http.HandlerFunc)

	// ServeHTTP makes Router implement http.Handler
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// ChiRouterInterface defines the methods we need from chi.Router
type ChiRouterInterface interface {
	Use(middlewares ...func(http.Handler) http.Handler)
	Route(pattern string, fn func(ChiRouterInterface)) ChiRouterInterface
	Get(pattern string, handlerFn http.HandlerFunc)
	Post(pattern string, handlerFn http.HandlerFunc)
	Put(pattern string, handlerFn http.HandlerFunc)
	Delete(pattern string, handlerFn http.HandlerFunc)
	Patch(pattern string, handlerFn http.HandlerFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// ChiAdapter wraps chi.Router to implement our Router interface
type ChiAdapter struct {
	router ChiRouterInterface
}

// NewChiAdapter creates a new ChiAdapter
func NewChiAdapter(chiRouter ChiRouterInterface) Router {
	return &ChiAdapter{router: chiRouter}
}

func (c *ChiAdapter) Use(middlewares ...func(http.Handler) http.Handler) {
	c.router.Use(middlewares...)
}

func (c *ChiAdapter) Route(pattern string, fn func(Router)) Router {
	var subRouter Router
	c.router.Route(pattern, func(r ChiRouterInterface) {
		subRouter = &ChiAdapter{router: r}
		fn(subRouter)
	})
	return subRouter
}

func (c *ChiAdapter) Get(pattern string, handlerFn http.HandlerFunc) {
	c.router.Get(pattern, handlerFn)
}

func (c *ChiAdapter) Post(pattern string, handlerFn http.HandlerFunc) {
	c.router.Post(pattern, handlerFn)
}

func (c *ChiAdapter) Put(pattern string, handlerFn http.HandlerFunc) {
	c.router.Put(pattern, handlerFn)
}

func (c *ChiAdapter) Delete(pattern string, handlerFn http.HandlerFunc) {
	c.router.Delete(pattern, handlerFn)
}

func (c *ChiAdapter) Patch(pattern string, handlerFn http.HandlerFunc) {
	c.router.Patch(pattern, handlerFn)
}

func (c *ChiAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}

// Factory function to create a new chi router wrapped in our interface
func NewRouter() Router {
	// This would import chi here, but to avoid circular imports,
	// we'll let the caller pass the chi router to NewChiAdapter
	panic("Use NewChiAdapter with chi.NewRouter() instead")
}
