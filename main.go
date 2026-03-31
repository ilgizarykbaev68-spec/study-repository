package main

import (
	"fmt"
	"study/feature1"
	"study/feature2"
	simpleconnection "study/feature_postgress/simple_connection"
)

func main() {
	fmt.Println("Hello, Git!")
	feature1.Feature1()
	feature2.Feature2()
	simpleconnection.CheckConnection()
}
