package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

// CreateStack combines multiple middlewares into a single Middleware
// Order matters: The first one in the list is the "outermost" (runs first)
func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}
