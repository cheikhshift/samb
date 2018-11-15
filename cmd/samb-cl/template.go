package main

var serverTemplate = `# Go package import path of your project.
package %s;


# import providers
require "./providers.se";


# Globals are exported via package 
# globals
global {
	name Prod;
	type bool;
    # Directive return specifies the 
    # value of the global variable
	return false;
}



server {
    host 127.0.0.1;
    port 8080;

    # Import web route definitions
    require "./endpoints.se";


    start {
    	do println("Hello");
    }

    shutdown {
    	do println("Bye");
    }  
}


`

var providerTemplate = `# Providers are used
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
}`

var routeTemplate = `# Routes' definition
# Import Go packages with directive import
# For example import "net/http";


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
	    }

	    route {
	    	method GET;
	    	path "Foo";

	    	go {
	    		# custom Go code to run
	    		do println("Hello");
	    	}

	    	# Handler can be any function.
	    	# Should be a function that handles the request
	    	# response
	    	handler fmt.Println("Hello");
	    }



	}
}`
