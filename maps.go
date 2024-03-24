package main

import "fmt"

func maps() {
	// simple map
	StudentAge := make(map[string] int)

	StudentAge["Ana"] = 25
	StudentAge["Jess"] = 20
	StudentAge["Eric"] = 23

	fmt.Println(StudentAge)
	fmt.Println(StudentAge["Jess"])

	// maps inside of maps
	// superhero := map[string]map[string]string{
	// 	"Superman" : map[string]string{
	// 	"RealName" : "Clarke Kent",
	// 	"City" : "Metropolis",
	// 	},

	// 	"Batman" : map[string]string{
	// 	"RealName" : "Bruce Wayne",
	// 	"City" : "Gotham City",
	// 	},
	// }
	// fmt.Println(superhero)

	// if temp, hero := superhero["Superman"]; hero {
	// 	fmt.Println(temp["RealName"], temp["City"])
	// }
}
