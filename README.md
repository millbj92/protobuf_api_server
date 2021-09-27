# Go gRPC server with ORM using the Ent framework. 


While I'm still a little unsteady about code generation, I can't deny that this framework was able to accomplish quite a bit in just a few minutes. I love that they've integrated protobuf into their workflow so nicely.  
  
  
This small POC was accomplished by following a tutorial from the main Ent website.

While it is just a small POC, and copied directly from a website, I feel like this small bit of code speaks volumes about the future of back end, and even the front-end.


With event-driven architecture being more tangible than ever before, I can absolutely see the value in using an ORM to bootstrap a protofile, create your routes, and move on to the next service!

Check out the tutorial here:  
https://entgo.io/blog/2021/03/18/generating-a-grpc-server-with-ent/
  
  
I will definitely be playing with this a lot more in the future.

The schema is in [./ent/schema/user.go](./ent/schema/user.go) if you'd like to see the model. I haven't tested with any other dialects than MySQL.

To run the app, there are two commands:

  - ```go run ./cmd/server/```
  - ```go run ./cmd/client/```

Once the server is started, it should automatically migrade the users table for you. The client will automatically make a Create request, retrieve the results, and send them back to stdout. 

There is a small unit test available as well: [./pb_test.go](./pb_test.go)
