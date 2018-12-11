package samb

// Project has your samb
//  directives.
type Project struct {
	Server          Server
	Require, Import []string
	// Provider defines the providers
	// to be used by HTTP routes.
	Provider  []Global
	Package   string
	Author    string
	Packages  []string
	Routes    Routes
	Templates Templates
	Global    []Global
}

type Global struct {
	Name, Type, Return, Comment string
}

// Server specifies the generated
// web server properties.
type Server struct {
	Host, Key string
	Port      string
	Webroot   string
	Require   []string
	// Routes field will enable
	// route nesting.
	Routes Routes
	Start  Go
	// Functions to be invoked
	// on panic.
	Recover  Go
	Init     Go
	Shutdown Go
}

// Routes holds a group of
// HTTP routes.
type Routes struct {
	Route   []Route
	Provide []string
	Doc     Documentation
}

// Route specifies the contract
// needed to be met by a request,
// and the handler (Go code) to be executed.
type Route struct {
	Method, Path string
	// Provide is a list of
	// provider names to be used by request.
	Provide []string
	Route   []Route
	Handler string
	// Go specifies Go code
	// to be ran prior to invocation
	// of handler. This code must respect
	// the scope of a Go HTTP handler function,
	// with variables r and w in scope.
	Go  Go
	Doc Documentation
}

// Array of Go statements
// to be executed.
type Go struct {
	Do []string
}

// Documentation will be used
// to generate HTML documentation
// of your code.
type Documentation struct {
	// Comment is used
	// with documentation
	// generation.
	Comment string
	// Alias specifies how the resource
	// should be used.
	Alias string
}

type Templates struct {
	Template []Template
}

type Template struct {
	FilePath, Type, Name string
}
