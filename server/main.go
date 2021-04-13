package main

func main() {

	r := SetupRouter()

	// listen and serve on 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		checkError(err)
	}
}
