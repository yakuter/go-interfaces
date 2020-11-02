package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//Simple Employee JSON object which we will parse
	empJson := `{
        "id" : 11,
        "name" : "Irshad",
        "department" : "IT",
        "designation" : "Product Manager"
	}`

	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(empJson), &result)

	//Reading each value by its key
	fmt.Println("Id :", result["id"],
		"\nName :", result["name"],
		"\nDepartment :", result["department"],
		"\nDesignation :", result["designation"])
}
