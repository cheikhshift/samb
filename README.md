# samb

[![Build Status](https://travis-ci.org/cheikhshift/samb.svg?branch=master)](https://travis-ci.org/cheikhshift/samb) [![GoDoc](https://godoc.org/github.com/cheikhshift/samb?status.svg)](https://godoc.org/github.com/cheikhshift/samb) [![Go Report Card](https://goreportcard.com/badge/github.com/cheikhshift/samb)](https://goreportcard.com/report/github.com/cheikhshift/samb) [![Maintainability](https://api.codeclimate.com/v1/badges/062da952018e56ea46d5/maintainability)](https://codeclimate.com/github/cheikhshift/samb/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/062da952018e56ea46d5/test_coverage)](https://codeclimate.com/github/cheikhshift/samb/test_coverage)

Go HTTP route management with Nginx like config language. The aim of `samb` is to provide a web server implementation
abstraction, where users can specify which Go handler/function should process a request for a request path.

Here is a sample server definition :

```
server {
    host 127.0.0.1;
    port 8080;

    # Import web route definitions
    require "./endpoints.se";


    start {
    	do println("Hello");
    	do println("Hello again");
    }

    shutdown {
	# directive do will execute passed
	# golang code
    	do println("Bye");
    }  
}

```

## Install

`samb` requires Go +v1.8

```
go get github.com/cheikhshift/samb/cmd/samb-cl
```

## Starting a project
Create a new directory, then run the following command.


	samb-cl -new -project=<NEW DIR PATH> <PACKAGE GO IMPORT PATH>

The command will add files in the new folder. Have a look at the files generated to give you an idea on how `samb` directives work. Some directives have comments further explaining their functionality.

## Transpiling
Run the following command to convert your directives into Go code.

	
	samb-cl -file=server.se -project=<NEW DIR PATH>

This will convert your directives into a Go library to handle your HTTP routes. A command will also be generated to launch your server, you can find the source code at `<NEW DIR PATH>/cmd/server`. 

## About parser
The following package is used to parse this Nginx like configuration language : [github.com/recoye/config](https://github.com/recoye/config), Nginx configuration style parser with golang

## Additional packages

[Parth : Path parsing for segment unmarshaling and slicing.](https://github.com/codemodus/parth)


## Todo

- [ ] Provide better documentation.
- [ ] Sublime/VSCode text plugins.
- [ ] Write package tests.
- [ ] Write tutorials/ guides.
- [ ] Introduce conditional operators with server and route directives.
- [ ] A command line to help with adding new handler source.
- [ ] Onyx command line to write to be transpiled tests with `this Nginx like configuration language`.
- [ ] Introduce a page directive.


### More samples

##### Sample Route

```
# Routes' definition
# Import Go packages with directive import
# For example import "net/http";
# samb source format checking.


routes {
    provide r;

    route {
	    method *;
	    # Defines route path.
	    # all sub routes have this path
	    # prepended.
	    path "/hello/";

	    # Provider variables
	    # within scope of entire 
	    # route.
	    provide w;
	    provide r;


	    go {
	    	do println("Hello");
	    	# The following commented
	    	# do directive will stop a
	    	# request :
	    	# do return;
	    }

	    route {
	    	method GET;
	    	path "Foo";

	    	go {
	    		do println("Hello");
	    	}

	    	# Handler can be any function.
	    	# Should be a function that handles the request
	    	# response. using provided variable r
		# with handler. This code will fail,
		# the referenced handler is not defined
	    	handler virtualPackage.Handle(r);
	    }


	}
}
```

##### Sample Providers

```
# Providers are used
# within endpoint requests.
provider {
	name r;
	type *http.Request;
    # directive return is not used here.
    # return can be used to define how your
    # provider is initialized. For example,
    # providing a variable with value "Foo"
    # : return string("Foo") 
}

provider {
	name w;
	type *http.ResponseWriter;
}
```
