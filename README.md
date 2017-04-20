# Go RestClient

This is a simple RestClient built in Go.

# How install

     $ git clone https://github.com/rancavil/restclient.git
     $ cd restclient
     $ go build restclient.go

or

     $ go get github.com/rancavil/restclient

Remember configurate GOPATH environment variable, example:

     $ export GOPATH=/home/myhome/goLibs

You can config your PATH to execute the commands in your environment.

     $ export PATH=$PATH:/home/myhome/goLibs/bin

# How execute

     $ ./resclient <url> <content-type> <file-with-request>

If you config your PATH, you can use:

     $ resclient <url> <content-type> <file-with-request>

For example, if you needs send a person (json) to register into rest service (restservice.com/person) you must do:

     $ ./restclient http://restservice.com/person application/json json-request.json

Or

     $ resclient <url> <content-type> <file-with-request>

If you config your PATH.

Example:

     $ restclient http://fakeserver.com/register application/json json-request.json

With json-request.json file is.

     {
          username: "joe"
          name: "Joe Doe"
          email: "joe@fakemail.com"
     }

