# Go RestClient

This is a simple RestClient built in Golang.

# How install

     $ git clone 
     $ go build restclient.go

# How execute

     $ ./resclient <url> <content-type> <file-with-request>

For example is you needs send a person to register into rest service (restservice.com/person) you must do:

     $ ./restclient http://restservice.com/person application/json json-request.json

With the json-request.json file.

     {
          username: "joe"
          name: "Joe Doe"
          email: "joe@fakemail.com"
     }



