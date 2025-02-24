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

**Tip:** Install Go extension (from go.dev) if you are using VS Code for development. 

## Creating a Go Project
Go requires the initialization of a module to be able to compile the project, for this, use this command: ``go mod init go_app`` inside of your project folder. It will generate a go.mod file that contains information about your module.

Go organizes the system into packages, so the files must be associated to packages and this is made by using the keyword ``package``. It is usual to use the package main for the main application. It also requires an entry point for the application, which is recognized by declaring a **main** function. The basic file for a Hello World app would look like below:


```go
package main

import "fmt" 

func main() {
	fmt.Println("Hello World")
}
```
*Each application must have only one main function.
** Differently from python, Go required the built-in packages to be explicitly imported into the application. Therefore we need to import the **format** package to have Print functionality available, ``import "fmt"``.

- Running the go application is very simple and only requires calling the command ``go run file_to_execute``. For our case it would be ``go run src/main.go``.

## Variables & Constants
 - Constants are declared with the **const** keyword, for example, ``cont myConstant = value``.
 - Variables are declared with the **var** keyword, and follow the Camel case syntax. For example, ``var myVariable = value``. It can also be done by the following sentence ``myVariable := value``.
 *When we assign a value for the variable at the creation, it recognize automatically the type.
 - When we just instantiate the variable without initializing it, it required a type definition, that can be done with the type after the var name ``var myVar string``
 - It is possible to print the variable type by using **%T** placeholder in a Print statement.
 - Go also have pointers, and like in C, the variable memory position can be passed with the **&** key ``var pointer = &myVar``
 - Variables can be made global (**package variables)** by creating them outside a function context. These can only be declared using the **var** keyword and not the **:=** fast assignment.
 
 ### Arrays
Arrays in Go are like in C++ and have fixed sizes. To define a new array we use square brackets [ ] with the size of the Array inside of it, followed by the type, defining the amount of values that can be stored. E.g. ``var myArray [50]string{}`` declares an Array with 50 possible values. 

It is possible to initialize the array with some values by adding such values inside the curly brackets -> ``var myArray = [50]string{"Maik", "Tom", "Bob"}``

*In Go, arrays accept only a single type. So it is not possible to use multiple types in the same array as in Python.

- Array elements are assigned, changed or read by using square brackets, like in Python and C++ ``fmt.Print(myArray[10])``. The first element is represented by 0.

### Slices
Slices are an abstraction of arrays that allows creating a list without a predefined size. The creation of a Slice is similar to Arrays, but without defining the size ``var mySlice []string{}``. And to add elements to the Slice we can use the built-in function **append** ``mySlice = append(mySlice, newValue)``.

### Maps
Go provide a data structure similar to a Dictionary, where we can save elements as a key, value pair. It is provided by the **map** type and a new variable of this type can be declared by using the key map, and specifying the key and value types. E. g. ``var myVar = map[string]int`` means the keys must be strings and the values integers.

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
for index, item := range itemsList {
    fmt.Printf("Item %v is: %v\n", index, item)
}
```
*Tip: an underscore can be used for not needed variables ``for _, item := range itemsList``

- **Conditional loop (while)**: use the **for** loop followed by the condition to be checked at every loop.
```go
runCounter := 0
for runCounter < 10 {
    fmt.Printf("COunter %v\n", runCounter)
    runCounter++
}
```

## Conditionals

### If - Else
Conditional **IF** statement is pretty similar to other languages. Use keyword **if** followed by the condition to be checked. Than use curly brackets to delimitate the code related to the condition.
Else-if follows the same principle and the remaining **else** just need no condition. The conditional checks that follow the first one must e declared in the same line that we closed the curly brackets.
```go
if myVar > 0 {
    fmt.Println("Var is positive!")
} else if myVar < 0 {
    fmt.Println("Var is negative!")
} else {
    fmt.Println("Var is NULL!")
}
```

### Switch
Switch statement allows to compare valid values from a variable in a simplest way, when compared to a bunch of if statements. It can be done using the **switch** keyword, followed by the checked variable and curly brackets. Inside the brackets we check all possible options using **case** keyword.
```go
switch cityName {
    case "Tokio":
        //code here
    case "Berlin":
        //code here
    case "Rome", "Paris":
        //code here
    default:
        fmt.Println("Invalid city!")
}
```
*Differently from other languages, Go will not check the next cases if one is valid, since it provides automatically an break when a case is executed. 


### Logic and Condition operators:
The operators in Go follow the same syntax from C++. So keywords like **or**, **and** and **not** from Python are not possible.
- AND - &&
- OR - ||
- NOT - !
- equal to - ==
- not equal - !=
- greater than - >
- smaller than - <

**Bitwise operators:**
- AND - &
- OR - |
- NOT - !
- XOR - ^
- left shift - <<
- right shift - >>

## Functions
Functions are very important elements which provide a better encapsulation of code, making possible the separation of the app in domains where each function process only one thing. This are fundamental for a organized and clean code structure. In Go a function is declared by using the **func** keyword followed by the function name and the code is delimited by curly brackets.

Arguments for the function need to be defined inside the parenthesis and have their types specified. Also for making possible returning something with the function, the return type must be defined outside the parenthesis and right before the opening of the brackets.
```go
func sumFunction(val1 int, val2 int) int {
    sum := val1 + val2
    return sum
}
```

In Go it is also possible to return more than one value. To specify the types of each return we need to use a second parenthesis declaring the types. Then in the **return** statement we separate values with a comma.
```go
func mathSubSum (val1 int, val2 int) (int,int) {
    sub := val1 - val2
    sum := val1 + val2
    return sub, sum
}
```

## Built-in functions

- **fmt.Print(string), fmt.Println(string), fmt.Printf(string, vars)** = print string on prompt.
- **fmt.Scan(varPointer)** = Get input from user.
- **len(variable)** = length of the variable
- **append(slices, new_item)** = add new value to the slices variable
- **strings.Field(stringVar)** = split the string with white space separator.
- **strings.Contains(stringVar, checkedOccurrence)** = check presence of substring in a string.
- **make(t Type)** = allocates and initialize an object of type slice, map or chan.
- **strconv.FormatUint(uint64(uint_value), base)** = format uint into a string according to desired base.


## Go Packages
In go we can also use multiple files to develop our application in a more organized way and better structured. To make one file recognize the declaration from another file we need to set the same **package** for both. For example, in the **main.go** file we have set `package main``, so we do teh same in another file that is shared with it.

Then we need to update the go command to make sure that the compiler includes all required files. This is accomplished by adding the extra files as following arguments in the run call. E.g ``go run src/main.go src/greetings.go`` If all the files are in the same folder it is possible to call ``go run .`` to run the complete folder.

Of course we are not going to have only one package in our whole project. When the project gets big we will have multiple packages, and for that we define different packages names to be added with the **package** keyword, and we aggregate all the files from the same package in a subdirectory. To include this a package in a file from our project we need to use the complete path to from the project source to make it valid. The project source is defined by the name declared as **module** in the **go.mod** file.

In go we also need to define if a function is visible outside the package, and for this we need to capitalize the first letter from the function name. For example, **getUserName() would be an private** function while **GetUserName() is public**.
