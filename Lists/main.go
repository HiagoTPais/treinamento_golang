package lists

fruits := []string{"apple", "banana", "orange"}

fruits = append(fruits, "grape")

fmt.Println(fruits[0])

for _, fruit := range fruits {
	fmt.Println(fruit)
}

subList := fruits[1:3]

fmt.Println(len(fruits))