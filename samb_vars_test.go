package samb

// Used as test file
var testFileOne = []byte(`

routes {

	route {
		method *;
		path "/panic";
	}
}

server {
  
  host localhost;
  port 5555;

}`)

// Used as test file
var testFileTwo = []byte(`

routes {

	route {
		method *;
		path "/new";
	}
}

server {
  
  host localhost;
  port 5556;

}`)

// Used as test file
var testFileThree = []byte(`

routes {

	route {
		method *;
		path "/panic";
	}
}

server {
  
  host localhost;
  port 1556;

}`)

// Used as test file
var testFileFour = []byte(`



routes {

	route {
		method *;
		path "/panic";
	}

	route {
		method *;
		path "/two";
	}
}

server {
  host localhost;
  port 556;

}`)
