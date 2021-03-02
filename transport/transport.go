package transport

type Server interface {
	// ListenAndServe listens on the TCP network address (given by
	// configuration) and then calls Serve to handle requests on
	// incoming connections.
	ListenAndServe() error

	// Shutdown gracefully shuts down the server without interrupting any
	// active connections
	Shutdown() error

	// GetAddress returns the server address for log usage
	GetAddress() string
}
