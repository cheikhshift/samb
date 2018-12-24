# samb

[![Build Status](https://travis-ci.org/cheikhshift/samb.svg?branch=master)](https://travis-ci.org/cheikhshift/samb) [![GoDoc](https://godoc.org/github.com/cheikhshift/samb?status.svg)](https://godoc.org/github.com/cheikhshift/samb) [![Go Report Card](https://goreportcard.com/badge/github.com/cheikhshift/samb)](https://goreportcard.com/report/github.com/cheikhshift/samb) [![Maintainability](https://api.codeclimate.com/v1/badges/062da952018e56ea46d5/maintainability)](https://codeclimate.com/github/cheikhshift/samb/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/062da952018e56ea46d5/test_coverage)](https://codeclimate.com/github/cheikhshift/samb/test_coverage)

`samb` offers a structured language to build RESTful HTTP APIs. `samb` provides syntax support for languages similar to those used to write infrastructure as code. It offers:

- A simple way to inject variables in request scope.
- Nest API routes without long path prefixes.
- Catch request runtime panics. 
- Generate Go code following guidelines.
- A tested library to use `samb` as you wish.

Once you finish writing your code, you may then, deploy your project to your cloud provider of choice. 

### Editor Plugins :

- [VScode](https://marketplace.visualstudio.com/items?itemName=GopherSauce.samb)
- [Sublime 3](https://packagecontrol.io/packages/SAMB)

### Documentation
Learn more about `samb` code generation : [here](https://github.com/cheikhshift/samb/wiki). Scroll down to find more samples.


## Install

#### Requirements

- `samb` requires Go +v1.8
- [dep](https://github.com/golang/dep) (Dependency management)
- `$GOPATH` environment variable set.

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
The following package is used to parse this Nginx like configuration language : [github.com/recoye/config](https://github.com/recoye/config), Nginx configuration style parser with golang.

The following package is used to parse YAML : [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2)

## Additional packages

[Parth : Path parsing for segment unmarshaling and slicing.](https://github.com/codemodus/parth)

-----

#### Tooling

Here is a list of tools that generate Go code for your SAMB projects.

samb-handler : Generate Go HTTP handlers with the specified parameters.

- Install : `go get github.com/cheikhshift/samb/cmd/samb-handler`
- Help : Run `samb-handler -h`

samb-provide: Generate and add a provider to your project. This will generate the Go source used with your provider as well.

- Install : `go get github.com/cheikhshift/samb/cmd/samb-provide`
- Help : Run `samb-provide -h`

samb-medic : Generate recover directive functions for your project.
	
- Install : `go get github.com/cheikhshift/samb/cmd/samb-medic`
- Help : Run `samb-medic -h`


## Todo

- [x] Provide better documentation. Checkout the [wiki](https://github.com/cheikhshift/samb/wiki).
- [x] Sublime/VSCode text plugins.
- [x] Write package tests.
- [ ] Write tutorials/ guides.
- [x] A command line to help with adding new handler source.
- [ ] Implement direct deployment to GCP/AWS/Azure and ~~any VM with SSH~~.


### Samples

Find more sample projects [here](https://github.com/cheikhshift/samb-examples).
Here is a sample server definition (in Nginx like language) :

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

In YAML:

```
# Go package import path of your project.
package: github.com/cheikhshift/samb-examples/yaml-example

# import providers
require: ["./providers.se"]


# Globals are exported via package 
# globals
global:
  - name: Foo
    type : bool
    # Adding comments to exported variable.
    comment: Foo decides if a process should run
    return : false
  - name: AnotherVariable
    type : bool
    return: true

server: 
  host: 127.0.0.1
  port: 8081
  # Import web route definitions
  require : ["./endpoints.yml"]

  start:
    do:
      - println("HelloWorld")
      - println("Starting...")
```

##### Sample Route (file named endpoints.se)

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
