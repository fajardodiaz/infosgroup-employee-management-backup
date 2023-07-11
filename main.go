package main

import "github.com/fajardodiaz/infosgroup-employee-management/initializer"

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDatabase()
}

func main() {

}
