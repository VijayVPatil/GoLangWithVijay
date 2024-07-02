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
![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/0baaa4d6-ff50-4551-8e93-4eac2f051234)


Click on the installed package . Go with the default settings and install.
After this run the below command in command prompt to check if go is installed or not.

### We need an IDE

I would prefer VS Code as my IDE

https://code.visualstudio.com/download
![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/9ebfff2c-a89b-4cf7-b5bb-338a16691454)


### Let's Write a basic Hello World Program
1) Create a Directory Called Main(You can give any name)
2) Create a file called main.go
3) Open the integrated Terminal from VS Code
4) Write the command:  go mod init  (Do not worry we will be learning this in future)
 ![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/7c32415d-fc1a-4155-addf-d97fc5014e03)

5) Now lets go write the simple go hello world program.
6) Install Go Extensions
![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/bb352bae-7042-4dcf-871d-8168309dcedd) 
7)Now we need a extension just to make our lives little easier   
Go to Extensions -> Go(By Go team at google). Click on install.    
After the installation is complete, open the command palette by pressing Ctrl + Shift + p     
Run the Go: Install/Update Tools command          
Select all the provided tools and click OK     


8) Make sure install all tools that you see at right bottom corner as a pop up

Program:

package main

import "fmt"

func main() {
	fmt.Println("Hello World!!!!")
}

1)The first line is where we defined the package of which the code is part.      
2)Next, we import the built-in package "fmt" or format .fmt is a Go package that is used to format basic strings, values, inputs, and outputs. It can also be used to print and write from the terminal.     
3)Next, we need to give a start point to go to start code execution . Hence the execution in go starts from the main function. One go program has only one main function.      
4)Next, we print the Hello Go! This is done by fmt package using a Println function to output the string.     

Now to run this go program , just run the command    
go run main.go     
![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/482727f3-2cd3-4cf5-8d9a-a447e08f4d2e)


