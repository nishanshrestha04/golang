# GO Lang Notes:

## Installation for linux:

Visit: [GO Downlaod](https://go.dev/dl/)
- Download the `tar.gz` go version
  
Go to the flolder where you have installed it and open the folder in termianl.

In my case. Its in my `/home/user/Download` folder and run the following command to extract and install the GO.

```bash
# Extract the file to '/usr/local`
sudo tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz

# Add FO to PATH
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
source ~/.bashrc # relaod the bash

# verify installation
go version

# You can run the following command to verify GO installation
mkdir ~/go-test && cd ~/go-test
echo '
package main
import "fmt"
func main() {
    fmt.Println("Go is working!")
}' \
> main.go

go run main.go

```
****
Now you are ready to use GO Lang 


## Data Types in GO

- int `1`
- string `"Hello World"`
- boolean `true/false`
- float `0.5`

## Variables
```go
var name string = "golang" // we have to use the variable is it being declared.


// simpler version
var name = "golang" // here go infer the type of the variable. i.e. "string"

// Short form to declare the variable
name := "golang"
```

## Constants
```go
// we use const keyword to define a constant value
const name = "golang" // we cannot change the value

// we can group the constant value in go 
const (
    port = 5000
    host = "localhost"
)
```

## Loops
-> For Loop 
For is only options for looping in GO lang
```go
// while loop format using for
i := 1
for i <= {
    fmt.Println(i)
    i = i + 1
}

// infinite loop
for {
    fmt.Println(1)
}

// traditional for loop
for i := 0; i < 3; i++{
    // break # breaks the code
    // continue skips a iteration
    fmt.Println(i)
}
for i := range 3{
    fmt.Println(i) // prints all the number < 3
}

```

## If Else
```go 
package main
func main(){
    // If else 
    age := 18

    if age >= 18{
        fmt.Println("Person is an Adult")
    } else{
        fmt.Println("Person is not an Adult")
    }

    // If else elif
    age := 16
    if age >= 18 {
        fmt.Println("Person is a Adult")
    } else if age >= 12 {
        fmt.Println("Person is a teenager")
    } else {
        fmt.Println("Person is a child")
    }

    // Examples
    var role = "admin"
    var hasPermissiono = true

    if role == "admin" || hasPermissions{
        fmt.Println("Yes")
    }

    // Directly declear varialbe in if condition
    if age :=15; age >= 18{
        fmt.Println("Person is an Adult", age)
    } else if age >=12 {
        fmt.Println ("Person is teenager", age)
    }
}
```
`Note: GO does not have ternary operator, we have to use noraml if else in that case`


## Switch statement

```go 
package main 

import "time"
func main(){
    // simple switch 
    i := 5
    switch i {
        case 1: 
            fmt.Println("One")
        case 2:
            fmt.Println("Three")
        case 3:
            fmt.Println("Three")
        default::
            fmt.Println("Other")
    }

    // Multiple condition switch
    switch time.Now().Weekday(){
        case time.Saturday, time.Sunday:
            fmt.Println("It's weekend")
        default:
            fmt.Println("It's working day")
    }

    // Type switch
    whoAmI := func(i interface{}){
        switch t := i.(type) {
            case int:
                fmt.Println("Its an Integer")
            case string:
                fmt.Println("Its a string")
            case bool:
                fmt.Println("Its a boolean")
            default:
                fmt.Println("other", t)
        }
    }
whoAmI(12) // Retruns its an integer
}

```