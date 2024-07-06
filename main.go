package main
import "fmt"

func main()  {
	fmt.Println("Hello, World!")

	// Variables
	var name string = "John Doe"
	var age int = 25
	var isCool = true
	fmt.Println(name, age, isCool)

	// Shorthand
	name2 := "Jane Doe"
	age2 := 30
	fmt.Println(name2, age2)

	// Constants
	const isCool2 = true
	fmt.Println(isCool2)

	// Multiple variables
	var (
		name3 = "John Doe"
		age3 = 25
		isCool3 = true
	)

	// Print type
	fmt.Printf("%T\n", name3)
	fmt.Printf("%T\n", age3)
	fmt.Printf("%T\n", isCool3)

	// struct
	type Person struct {
		firstName string
		lastName string
		city string
	}
	person1 := Person{firstName: "John", lastName: "Doe", city: "New York"}
	fmt.Println(person1.firstName)

	// Array
	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fmt.Println(fruits, fruits[1])

	// Slice
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])

	// Append
	fruitSlice = append(fruitSlice, "Banana")
	fmt.Println(fruitSlice)

	// Map
	emails := make(map[string]string)
	emails["John"] = "12"
	emails["Jane"] = "13"
	fmt.Println(emails["John"])
	fmt.Println(len(emails))

	// Delete
	delete(emails, "John")
	fmt.Println(emails)

	// Range
	ids := []int{33, 76, 54, 23, 11}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for loop map
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// for loop person   slice形式
	persons := []Person{
		{firstName: "John", lastName: "Doe", city: "New York"},
		{firstName: "Jane", lastName: "Doe", city: "New York"},
    }

    for _, person := range persons {
		fmt.Println(person)
    }

	// if
	x := 10
	y := 10
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	// switch
	color := "red"
	switch color {
		case "red":
			fmt.Println("Color is red")
		case "blue":
			fmt.Println("Color is blue")
		default:
			fmt.Println("Color is not red or blue")
	}

	
 }