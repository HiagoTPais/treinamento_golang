package tuples

type Person struct {
	Name string
	Age int
	City string
}

person := Person{"Alice", 30, "New York"}

fmt.Println(person.Name, person.Age, person.City)

point := struct { X, Y int }{10, 20}

fmt.Println(point.X, point.Y)