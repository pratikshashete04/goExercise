# goExercise
This repository consists of the basic programs of Go Language

### Day-1 :
## 1. Functions
* A function is a set of statements that execute a process together. Every Go program has at least one function, which is main().A function declaration informs the compiler the name, form of return and parameters of a method. The internal body of the system is given by function specification. 

* In Go programming language the general form of a function description is as follows − 

func function_name(parameter_list) (return_types){
    body_of_the_function
}

============================================================================

## 2. Slice
* Slice is more efficient, scalable, convenient than array and is a lightweight data structure in Go language.
* Slice is a variable-length list that stores identical categories of elements(it does not store any data, it just describes a section of an underlying array), so you are not permitted to store specific types of elements in the same slice. 
* Slice is much like an array and has an index value and length, except the scale of the slice is resized much as an array, they are not of fixed format.
* Inside slice and list are connected, a slice is a connection to an underlying sequence.
* The duplicate items can be contained in the slice 

============================================================================

## 3. Structs
* A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type. 
* Every object in real world that has any collection of properties/fields can be described as a struct.
* It can be termed as a lightweight class that does not support inheritance but supports composition. 

============================================================================

### Day-2 :
## 1. Method
* Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it.
* The method can access the properties of the receiver with the help of the receiver argument.
* Here, the receiver can be of struct type or non-struct type.

* In Go programming language the general form of a method is as follows − 

func (receiver_name Type) method_name (parameter list)(return type){
    body_of_the_method
}

* Difference between method and function:
- Method contain receiver; while function does not contain any receiver.
- Method can accept both pointer and value; while function cannot accept both pointer and value.
- Methods of the same name but different types can be defined in the program; while functions of the same name but different type are not allowed to define in the program.
============================================================================

## 1. Map
* 


============================================================================

## 1. Pointer
*