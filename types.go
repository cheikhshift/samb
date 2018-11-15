package samb

type Project struct {
	Server          Server
	Require, Import []string
	Provider        []Global
	Package 		string
	Author          string
	Packages        []string
	Routes          Routes
	Templates       Templates
	Global          []Global
}

type Global struct {
	Name, Type, Return string
}

type Server struct {
	Host, Key string
	Port      string
	Webroot   string
	Require   []string
	Start     Go
	Init      Go
	Shutdown  Go
}

type Routes struct {
	Route   []Route
	Provide []string
}

type Route struct {
	Method, Path string
	Provide      []string
	Route        []Route
	Handler      string
	Go           Go
}

type Templates struct {
	Template []Template
}

type Template struct {
	FilePath, Type, Name string
}

type Go struct {
	Do []string
}
