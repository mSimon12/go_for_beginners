# Go application
In this repository I present my own steps learning Go, while developing a first application.

## Installing Go Compiler

The first step to develop an application in Go is to install the required compiler. The installer or instructions fr installing it in you OS can be found at [Go Installation](https://go.dev/doc/install).

If you are running it on **Linux**, proceed with the following steps:
1. Remove previous Go installations:
`` rm -rf /usr/local/go`` and maybe ``sudo apt-get remove golang-go``
2. Download the installer from the link above, than extract it on */usr/local*
``sudo tar -C /usr/local -xzf go1.24.0.linux-amd64.tar.gz``
3. Add */usr/local/go/bin* to the PATH environment variable.
``export PATH=$PATH:/usr/local/go/bin``
4. Reboot your computer or run ``source $HOME/.profile``
5. Verify the installation with ``go version`` command.

**Tipp:** Install Go extension (from go.dev) if you are using VS Code for development. 

## Creating a Go Project
Go requires the initialization of a module to be able to compile the project, for this, use this command: ``go mod init go_app`` inside of your project folder. It will generate a go.mod file that contains information about your module.

Go organizes the system into packages, so the files must be associated to packages and this is made by using the keyword ``package``. It is usual to use the package main for the main applicaion. It also requires an entry point for the application, which is recognized by delaring a **main** function. The basic file for a Hello World app would look likt below:


```go
package main

import "fmt" 

func main() {
	fmt.Println("Hello World")
}
```
*Each application must have only one main function.
** Differently from python, Go required the built-in packages to be excplicitly imported into the application. Therefore we need to import the **format** package to have Print funcionallity available, ``import "fmt"``.

- Running the go application is very simple and only requires calling the command ``go run file_to_execute``. For our case it would be ``go run src/main.go``.

## Variables & Constants
 - Constants are declared with the **const** keyword, for example, ``cont myConstant = value``.
 - Variables are declared with the **var** keyword, and follow the Camel case syntax. For example, ``var myVariable = value``. It can also be done by the following sentence ``myVariable := value``.
 *When we assign a value for the variable at the creation, it recognize automatically the type.
 - When we just instanciate the variable without initializing it, it required a type definition, that can be done with the type after the var name ``var myVar string``
 - It is possible to print the variable type by using **%T** placeholder in a Print statement.
 - Go also have pointers, and like in C, the variable memory position can be passed with the **&** key ``var pointer = &myVar``
 
 ## Arrays
Arrays in Go are like in C++ and have fixed sizes. To define a new array we use square brackets [ ] with the size of the Array inside of it, followed by the type, defining the amount of values that can be stored. E.g. ``var myArray [50]string{}`` declares an Array with 50 possible values. 

It is possible to initialize the array with some values by adding such values inside the curly brackets -> ``var myArray = [50]string{"Maik", "Tom", "Bob"}``

*In Go, arrays accept only a single type. So it is not possible to use multiple types in the same array as in Python.

- Array elements are assigned, changed or read by using square brackets, like in Python and C++ ``fmt.Print(myArray[10])``. The first element is represented by 0.

## Slices
Slices are an abstraction of arrays that allows creating a list without a predefined size. The creation of a Slice is similar to Arrays, but without defining the size ``var mySlice []string{}``. And to add elements to the Slice we can use the built-in function **append** ``mySlice = append(mySlice, newValue)``.

## Loops
As well as in any language, Go provides a loops for executing code multiple times. Although, instead of providing multiple options like **while**, **do-while** and **for-each**, Go have only the **for** loop, which is made in a way that it can be implemented for all representing all the loop types.
- **Infinite loop:** can be accomplished by declaring the **for** loop without any parameter.
```go
for {
    fmt.Print("This will run forever")
}
```
- **Iterable loop**: use the **for** loop followed by variables that will receive the elements at each loop, keyword **range** and the iterable from where the items will be extracted, with a **:=** between them names for without any parameter.
***range** iterates over elements, and provides index and value for arrays and slices.
```go
for index, item := range itesList {
    fmt.Printf("Item %v is: %v\n", index, item)
}
```
*Tipp: an underscore can be used for not needed variables ``for _, item := range itesList``

## Built-in fuctions

- **fmt.Print(string), fmt.Println(string), fmt.Printf(string, vars)** = print string on prompt.
- **fmt.Scan(varPointer)** = Get input from user.
- **len(variable)** = length of the variable
- **append(slices, new_item)** = add new value to the slices variable
- **strings.Field(stringVar)** = split the string with white space separator.