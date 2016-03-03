package main

type NotFoundError struct {}

func (e NotFoundError) Error() string {
    return "Element Not Found"
}


type JsonInvalidError struct {}

func (e JsonInvalidError) Error() string {
    return "Json invalid"
}
