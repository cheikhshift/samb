# samb

Go HTTP route management with Nginx like config language

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

- Provide better documentation.
- Sublime/VSCode text plugins.
- Write package tests.
- Write tutorials/ guides.


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
	    	# response
	    	handler fmt.Println("Hello");
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
