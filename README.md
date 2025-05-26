# About
Basic game world simulation written in Go/React/WebGL. The math functions are using linear algebra and Geometry concepts (Plane, Vector3, Triangle). When I created this project, I wanted to see what these core math functions would look like in Go instead of C++/C# (where I've implemented something similar here for [example](https://github.com/sandboxcoder/CSharpRayTest)).

# How to run
Make sure [Golang](https://go.dev/doc/install) is installed (minimum `1.23`). 

Easiest way to run the app is within Visual Studio Code via the "Run Go and React" option in the dropdown (if you already built the react client). If not, then at least build the react client like this:
```
cd client
npm install
npm run build
```


Alternate method to build & run both apps from within a bash terminal. The command below will build the frontend client code and run the Go server:
```
cd client
npm install
npm run build
cd ..
go run .
```
Now type this command to spin up the react client host:
```
npm run start
```
Connect to http://localhost:3000 to view the single page app (SPA).

# How to run the tests from command line
```
go test -v Kadane/core
```

# Differences from C#/C++

Go doesnt utilize *operator overloading* so I needed to get a bit creative in the *Vector3* class where I wanted to implement Add, Subtract, and Multiply operations.