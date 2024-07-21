## Taking User Input

This is one of the way to take User input. We will be using bufio package that go provides  
Link: https://pkg.go.dev/bufio#example-Scanner-Words

1. To Define bufio ,Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
2. We have a new method called New Reader method that takes input from os.Stdin

We would be storing it in a variable 'reader'  
 reader := bufio.NewReader(os.Stdin)

Here we use comma ok syntax as using reader we will be using a ReadString method that reads till the delimiter  
It return the data read before the error and error it self if it occurs. Now if do not require error we replace it with \_ , this will only give me the input and  
ignore the other.
