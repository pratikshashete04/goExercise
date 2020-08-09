# goExercise
This repository consists of the basic programs of Go Language

## 1. Functions:
* A function is a set of statements that execute a process together. Every Go program has at least one function, which is main().A function declaration informs the compiler the name, form of return and parameters of a method. The internal body of the system is given by function specification. 

* In Go programming language the general form of a function description is as follows − 

func function_name(parameter_list) (return_types){
    body_of_the_function
}

============================================================================

## 2. Slice:
* Slice is more efficient, scalable, convenient than array and is a lightweight data structure in Go language.
* Slice is a variable-length list that stores identical categories of elements(it does not store any data, it just describes a section of an underlying array), so you are not permitted to store specific types of elements in the same slice. 
* Slice is much like an array and has an index value and length, except the scale of the slice is resized much as an array, they are not of fixed format.
* Inside slice and list are connected, a slice is a connection to an underlying sequence.
* The duplicate items can be contained in the slice 

============================================================================

## 3. Structs:
* A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type. 
* Every object in real world that has any collection of properties/fields can be described as a struct.
* It can be termed as a lightweight class that does not support inheritance but supports composition. 

============================================================================

## 4. Method:
* Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it.
* The method can access the properties of the receiver with the help of the receiver argument.
* Here, the receiver can be of struct type or non-struct type.

* In Go programming language the general form of a method is as follows − 

func (receiver_name Type) method_name (parameter list)(return type){
    body_of_the_method
}

### Difference between method and function:
- Method contain receiver; while function does not contain any receiver.
- Method can accept both pointer and value; while function cannot accept both pointer and value.
- Methods of the same name but different types can be defined in the program; while functions of the same name but different type are not allowed to define in the program.
=========================================================================

## 5. Map:
* Golang Maps is a collection of unordered pairs of key-value. 
* It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys.
* In maps the values are not unique like keys and can be of any type like int, float64, rune, string, pointer, reference type, map type, etc.

* In Go programming language we can create a map using the following syntax:

// An Empty map
map[Key_Type]Value_Type{}

// Map with key-value pair
map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

============================================================================

## 6. Pointer:
* Pointers in Go programming language or Golang is a variable which is used to store the memory address of another variable.
* Pointers in Golang is also termed as the special variables. 
* The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).
* The default value or zero-value of a pointer is always nil.

* We can declare a pointer as follows:
  var pointer_name *Data_Type

============================================================================

## 7. Variadic function:
* The function that called with the varying number of arguments is known as "variadic function".
* A user is allowed to pass zero or more arguments in the variadic function.
* fmt.Printf is the example of the variadic function, it required one fixed argument at the starting after that it can accept any number of arguments.

* In Go programming language we can use variadic function by using the following syntax:

func function_name(para1, para2...type)type{
// code...
}

============================================================================

## 8. Closures:
* Go language an anonymous function can form a closure. 
* A closure is a special type of anonymous function that references variables declared outside of the function itself. 
* It is similar to accessing global variables which are available before the declaration of the function.

============================================================================

## 9. Flow control statements:

For loop:
* Go has only one looping construct, the for loop.
* The basic for loop has three components separated by semicolon (initial/condition/post).
* "for" is Go's "while".

* The syntax of a simple for loop is:
for initialization; condition; post {
  // statements..
}
_____________________________________________________________________________

If, If-else:
* The if statement: executes some code if the condition is true.
syntax: 
if condition{
  // code to be executed if the condition is true
}

* The if...else statement: executes some code if the condition is true and another if that condition is false. 
syntax:
if condition{
  // code to be executed if condition is true
}else{
  // code to be executed if the condition is false
}

* The if...else if...else statement: executes different code for more than two conditions.
syntax:
if condition-1{
  // code to be executed if condition-1 is true
}else if condition-1{
  code to be executed if condition-2 is true
}else{
  // code to be executed if both condition-1 and condition-2 are false
}
_____________________________________________________________________________

Switch:
* In goLang the switch statement is a multi way branch statement.
* It provides an efficient way to transfer the execution to different parts of a code based on the value(case) of the expression.
* There are two types of switch statements:

1.Expression Switch
syntax:
switch optionalStatement; optionalExpression{
  case expression1: Statement..
  case expression2: Statement..
  ...
  default: Statement
}

2.Type Switch
syntax:
switch optionalStatement; typeSwitchExpression{
  case typeList 1: Statement..
  case typeList 2: Statement..
  ...
  default: Statement..
}
_____________________________________________________________________________

Defer:
* In Go language, defer statements delay the execution of the function or method or an anonymous function until the nearby functions returns.
syntax:
// function
defer func func_name(parameter_list Type) return_Type{
  // code
}

// method
defer func (receiver Type) method_name(parameter_list){
  // code
}

// anonymous function
defer func (parameter_list)(return_type){
  // code
}()

============================================================================

## 10. Bbolt:
* In Bolt we can store database as key/value pair.
* Bolt is Go's package i.e. "github.com/boltdb/bolt".
* Bolt is not a relational database, it is a NoSQL database.
* Storage of data in bolt is divided into buckets.
* A bucket is simply a named collection of key/value pairs, just like Go's map.

============================================================================

## 11. goroutine:
* The goroutine is nothing but a function or method that is running in background concurrently with other goroutines.
* Go provides a special keyword "go" to create goroutine.
* When we call a  function or method with go prefix, then that function or method executes in a goroutine.
* We can apss control to another goroutines manually by telling the scheduler to schedule another available goroutines.
* For that, we can use time.Sleep() call.

============================================================================

## 12. Interface:
* An interface is a collection of method signatures that a Type can implement.
* Interface defines the behavior of the object of the type Type.
* In Go language, it is necessary to implement all the methods declared in the interface for implementing an interface.
* The Go language interfaces are implemented implicitly and it does not contain any specific keyword to implement an interface.
* In Go language, we can create an interface by using following syntax:
type interface_name interface{
  // Method signatures
}

============================================================================

