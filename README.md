# Go Basics

## 1. Variable Declaration

Declare a variable with a type but no value:

```go
var name string
```

Declare a variable with a type and value:

```go
var name string = "Huda"
```

Declare multiple variables:

```go
var country, continent string = "UK", "Europe"
```

Block declaration of multiple variables:

```go
var (
    isEmployed bool   = true
    salary     int    = 5000
    position   string = "developer"
)

```

2. Printing
   Use fmt.Printf to format and print values:

   ```go
   fmt.Printf("This is my country %s and this is my continent %s\n", country, continent)

   ```

3. Constants `const pi = 3.14`(untype) or `const name string = "Huda"`
   Compiler doesn't care if `const` isn't used, unlike `var`

4. [Enum](https://go.dev/wiki/Iota) with iota
   Use iota to create a sequence of related constants:

   ```go
    const (
        Jan = iota + 1
        Feb
        Mar
        Apr
   )
   ```

5. Function
   Capitalised function name is public and exported, uncapitalised function is private function

```go
    func add(a int, b int) int { // every param must declare type, and specify the type of the return value
        return a + b
    }
```

6. Array
   - Must declare the size of the array, and can't have mixed type in one array
   - Can't append (add, delete) but change change value at the index
   - Array can be more useful for memory allocation

```go
numbers := [5]int{1,2,3,4,5}
```

7. Slice
   Like array but more dynamic, it can be append and change

```go
allNumbers := numbers[:] //copy all value in the array
```

Declare slice

```go
fruits := []string{"apple", "banana", "strawberry"} // don't declare size, indicating its a slice

//append new item to the slice, this will create a new slice
fruits = append(fruits, "kiwis")
```

8. Map (hashmap/ key value)

```go
capitalCities := map[string]string{
   "USA": "Washington D.C.",
   "India": "New Delhi",
   "UK": "London"
}
```

9. Struct (structure, i.e: data type, object)

```go
type Person struct {}
```

10. Pointer and references
    Passing a pointer to the parameter, when calling this function, the value inside the scope will persist,

```go
func modifyPersonName(person *Person) {
   person.Name = "Huda"
   fmt.Println("Inside scope: new name", person.Name)
}

// when calling the function must pass the address of the param (with the '&' ampersand)
modifyPersonName(&person)

```

This is a Go 'method'

```go
func (person *Person) modifyPersonName() {
   person.name = "Huda"
   fmt.Println("Inside scope: new name", person.Name)
}

// declare the variable..
person := Person struct {}...

// call this 'method' (like a class in OOP)
person.modifyPersonName()
```

# DOCKER

running up docker container
`docker compose up --build`

-- TO GO INTO THE DB SERVER
psql -U postgres -h localhost -p 5432
