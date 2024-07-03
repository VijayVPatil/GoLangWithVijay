### Let's GO

Go is a procedural programming language. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language.
Go lang is strong and statically typed, meaning variables can only have a single type. It is also a concurrent, and garbage-collected programming language.
It is designed to be simple, efficient, and easy to learn, making it a popular choice for building scalable network services, web applications, and command-line tools.
Go Lang, is compiled language which means a go code (main.go) is first compiled into an executable file (main.exe) using a compiler. Go has a fast compiler, which makes it easy to iterate quickly during development.

### Why Go?

The infrastructure has evolved over the last decade. Multi-core processors were becoming common and the use of cloud servers with thousands of servers and multiple processors was becoming a norm to deploy applications.
Here most of the applications failed to take benefit from this infrastructure advancement. Applications were only able to execute one task at a time but infrastructure could enable it.
Now many languages like Java, and C++ had features like concurrency and multithreading, but code could get pretty complicated. This is where Go comes into the picture as it is built while keeping in mind the above needs.
Go was designed to run on multiple cores and built to support concurrency.

### Installation

Get you latest version of Go Lang according to your operating system

https://go.dev/doc/install

Click on the installed package . Go with the default settings and install.
After this run the below command in command prompt to check if go is installed or not.

```
~ go version
```

Above should give the version of go installed with the Operating System installed for

### We need an IDE

I would prefer VS Code as my IDE

Install based upon your operating system 
https://code.visualstudio.com/download

### Let's Write a basic Hello World Program

1. Create a Directory Called First_Program(You can give any name)
2. Create a file called main.go
3. Open the integrated Terminal from VS Code
4. Install Go Extensions
   Below link show step by step process to setup go in vs code
   Link: https://learn.microsoft.com/en-us/azure/developer/go/configure-visual-studio-code

   Or you can follow below steps:
   Now we need a extension just to make our lives little easier  
   Go to Extensions -> Go(By Go team at google). Click on install.  
   After the installation is complete, open the command palette by pressing Ctrl + Shift + p  
   Run the Go: Install/Update Tools command  
   Select all the provided tools and click OK

5. Write the command: go mod init (Do not worry we will be learning this in future)
6. Now lets go write the simple go hello world program.

##### Checkout the Hello Folder in Day 01-Installation for the go code

Program:

```
package main

import "fmt"

func main() {
	fmt.Println("Hello World!!!!")
}
```

### Code walkthrough
1. The first line is where we defined the package of which the code is part.
2. Next, we import the built-in package "fmt" or format .fmt is a Go package that is used to format basic strings, values, inputs, and outputs. It can also be used to print and write from the terminal.
3. Next, we need to give a start point to go to start code execution . Hence the execution in go starts from the main function. One go program has only one main function.
4. Next, we print the Hello World!!!! This is done by fmt package using a Println function to output the string.

###### But what is the go mod init command that we ran?

In golang , as we seen above we imported the Println function from the fmt package. Now comes from the standard library, but when we want to import the packages from other module such as ex: Gorilla Mux(Routing Package used for backend API) we need to manage these dependencies using our own module and we do that by go.mod file . It helps us in tracking the dependencies.
You and me will lear more about this in future sections.

Now to run this go program, right click on the main.go file that you made and choose open in integrated terminal, just run the command  
go run main.go and should get the output like below

```
Hello World!!!!
```
