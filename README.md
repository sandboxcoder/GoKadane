# About
Basic game world simulation written in Go. The math functions are using linear algebra and Geometry concepts (Plane, Vector3, Triangle). When I created this project, I wanted to see what these core math functions would look like in Go instead of C++/C# (where I've implemented something similar here for [example](https://github.com/sandboxcoder/CSharpRayTest)).

# How to run
Make sure [Golang](https://go.dev/doc/install) is installed. From command line just type:
```
cd client
npm install
npm run build
cd ..
go run .
```
This will start the server. Connect to the server (http://localhost:8080 by default), to view the web page.

# How to run the tests from command line
```
go test -v Kadane/core
```

# Differences from C#/C++

Go doesnt utilize *operator overloading* so I needed to get a bit creative in the *Vector3* class where I wanted to implement Add, Subtract, and Multiply operations.