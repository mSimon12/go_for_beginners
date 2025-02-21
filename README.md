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

### Variables & Constants
 - Constants are declared with the **const** keyword, for example, ``cont myConstant = value``.
 - Variables are declared with the **var** keyword, and follow the Camel case syntax. For example, ``var myVariable = value``. It can also be done by the following sentence ``myVariable := value``.
 *When we assign a value for the variable at the creation, it recognize automatically the type.
 - When we just instanciate the variable without initializing it, it required a type definition, that can be done with the type after the var name ``var myVar string``
 - It is possible to print the variable type by using **%T** placeholder in a Print statement.
 - Go also have pointers, and like in C, the variable memory position can be passed with the **&** key ``var pointer = &myVar``


### Built-in fuctions

- **fmt.Print(string), fmt.Println(string), fmt.Printf(string, vars)** = print string on prompt.
- **fmt.Scan(varPointer)** = Get input from user.