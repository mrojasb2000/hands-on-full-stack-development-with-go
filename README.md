# hands-on-full-stack-development-with-go
Hands-On Full Stack Development with Go

Preface

The Go programming language has been rapidy adopted by developers for building web applications. With its impressive performance and ease of development, 

Chapter 2, h3. Go's Building Blocks.

For you to write Go code on your computer, you need to set up a Go workspace. A Go workspace is a folder where you will write your Go code. Setting up a Go workspace is relatively simple.

1.- First, make sure that you have Go installed. As we mentioned earlier, you can download and install Go site http://golang.org

2.- After installing Go, create a new folder in your computer for Go's workspace. Mine is called GoProjects.

3.- Inside your Go workspace folder, you will have to create three main folders: src, pkg, bin. It is very important to create folders with these exact names inside your Go workspace folder.

  * The src folder will host all of your code files.
  * The pkg folder typically hosts the compiled package files of your code.
  * The bin folder typically hosts the binary files that are produced by your Go programs.

4.- You will need to set two environment variables:
  
  * The first enviroment variable is called GoRoot, and will include the path to your Go install. GoRoot should tipically be taken care of by the Go installer. Howwever, if it's missing, or you would like to move your Go installation to a different place, then you need to set GoRoot.

  * The second environment variable is called GoPath. GoPath includes the path to your Go workspace folder. By default, if not set, GoPath is assumed to either to be as $HOME/go on Unix systems or %USERPROFILE%\go Windows.  


Chapter 3, h3. Go Concurrency

Chapter 4, h3. Frontend with React.js

Chapter 5, h3. Building a Frontend for GoMusic

Chapter 6, h3. RESTful Web APIs in Go with the Gin

Chapter 7, h3. Advanced Web Go Applications with Gin and React

Chapter 8, h3. Testing and Benchmarking Your Web API

Chapter 9, h3. Introduction to Isomorphic Go with GopherJS

Chapter 10, h3. Where to Go from here?