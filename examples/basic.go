package main

import (
	"net/http"

	mesquite "github.com/go-mesquite/Mesquite"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + mesquite.URLParam(r, "Message")))
}

/*
// --- ---- --- Middleware --- --- ---
// Logger is a middleware handler that does request logging

	type Logger struct {
		handler http.Handler
	}

// ServeHTTP handles the request by passing it to the real handler and logging the request details

	func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		// Start the timer to see how long the controller took to execute
		startTime := time.Now()

		// Run the handler/controller
		l.handler.ServeHTTP(w, r)

		// Stop the timer and get how long the controller took to run
		duration := time.Since(startTime)
		milliseconds := duration.Milliseconds()

		// Instead of printing "0ms", we should print <1ms if it executed quickly
		if milliseconds == 0 {
			log.Printf("\x1b[32m%s %s <%dms\x1b[0m\n", r.Method, r.URL.Path, 1)
		} else {
			log.Printf("\x1b[32m%s %s %dms\x1b[0m\n", r.Method, r.URL.Path, milliseconds)
		}
	}

// NewLogger constructs a new Logger middleware handler

	func NewLogger(handlerToWrap http.Handler) *Logger {
		return &Logger{handlerToWrap}
	}
*/

func main() {
	// Optionally use the Mesquite router
	router := mesquite.NewRouter()

	// Create a route/function for one or more pages
	// There is a lot going on in the next line but it's simple if we break it down
	// http.MethodGet is the method (like GET, POST, DELETE ect.)
	// "/" is the URL that this route is attached to. Regex can be used here
	// func... is just an inline function defining the route. This could be done separately
	// http.ResponseWriter is the interface for constructing an HTTP response
	// *http.Request is the pointer to data that represents an HTTP request received by the server
	router.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root root root"))
	})

	// Use a shortcut to reference a function (I can't decide if I like this way)
	router.GET(`/hello/(?P<Message>\w+)`, Hello)

	// Serve static files from a folder
	router.Static("staticfiles", "/static")

	// TODO Add later

	//router.UseMiddleware(Func) (Figure out naming for different types of middleware needed)
	//router.Templates("templates")
	//router.ControllerFor404(Handle404), 500

	http.ListenAndServe(":8000", router)
}
