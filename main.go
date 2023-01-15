package main

import (
	// "encoding/json"
	"fmt"
	"os"
	"time"

	// "os/user"
	"strings"

	color "github.com/AdityaSubrahmanyaBhat/golang/dashDB/Colors"
	errorHandler "github.com/AdityaSubrahmanyaBhat/golang/dashDB/Error"
	dbFunctions "github.com/AdityaSubrahmanyaBhat/golang/dashDB/Functions"
	address "github.com/AdityaSubrahmanyaBhat/golang/dashDB/models/Address"
	u "github.com/AdityaSubrahmanyaBhat/golang/dashDB/models/User"
)

func main() {
	dir := "./"
	driver, err := dbFunctions.CreateNewDB(dir, nil)
	errorHandler.HandleError(err)
	employees := []u.User{
		{Name: "Aditya", Age: "21", Company: "jpmc", Address: address.Address{City: "Mysuru", State: "Karnataka", Country: "India", Pincode: "570002"}},
		{Name: "pranav", Age: "21", Company: "siemens", Address: address.Address{City: "Mysuru", State: "Karnataka", Country: "India", Pincode: "570002"}},
		{Name: "prashanth", Age: "21", Company: "halodoc", Address: address.Address{City: "Bangalore", State: "Karnataka", Country: "India", Pincode: "560002"}},
		{Name: "Gajanana", Age: "21", Company: "halodoc", Address: address.Address{City: "Mysuru", State: "Karnataka", Country: "India", Pincode: "570002"}},
	}

	for _, record := range employees {
		dbFunctions.Write("users", strings.ToLower(record.Name), u.User{
			Name:    record.Name,
			Age:     record.Age,
			Company: record.Company,
			Address: record.Address,
		}, driver)
	}

	// records, err := dbFunctions.ReadAll("users", driver)
	// errorHandler.HandleError(err)
	// fmt.Println(records)

	// var allUsers []string

	// for _, record := range records {
	// 	employeeFound := u.User{}
	// 	err := json.Unmarshal([]byte(record), &employeeFound)
	// 	errorHandler.HandleError(err)
	// 	allUsers = append(allUsers, employeeFound.Name)
	// }
	// fmt.Println(allUsers)

	

	fmt.Println(color.GetColor("Blue"), "Starting dashDB...")
	time.Sleep(3 * time.Second)
	fmt.Println(color.GetColor("Green"), "------------dashDB is live now------------", color.GetColor("Reset"))
	time.Sleep(2 * time.Second)
	choice := "n"
	var operationChoice int
	for choice != "y" {
		fmt.Println("1. Read all records\n2. Read particular record\n3. Create a new record\n4. Delete particular record\n5. Delete all records\n6. Exit")
		fmt.Print("Enter your choice of operation : ")
		fmt.Scanln(&operationChoice)
		switch operationChoice {
		case 1:
			{
				records, err := dbFunctions.ReadAll("users", driver)
				errorHandler.HandleError(err)
				fmt.Println(records)
			}
		// case 2: dbFunctions.Read("users", "Prashanth", driver);break;
		case 4:
			{
				var recordName string
				fmt.Print("Enter the name of the record to be deleted : ")
				fmt.Scanln(&recordName)
				err := dbFunctions.Delete("users", recordName, driver)
				if err != nil {
					errorHandler.HandleError(err)
				} else {
					fmt.Print(color.GetColor("Green"), "")
					fmt.Printf("Record with the name \"%s\" has been deleted successfully\n", recordName)
					fmt.Print(color.GetColor("Reset"), "")
				}
			}
		case 5:
			{
				err := dbFunctions.Delete("users", "", driver)
				if err != nil {
					errorHandler.HandleError(err)
				} else {
					fmt.Println(color.GetColor("Green"), "All records deleted successfully", color.GetColor("Reset"))
				}
			}
		default:
			{
				fmt.Println(color.GetColor("Blue"), "Exiting...", color.GetColor("Reset"))
				time.Sleep(1 * time.Second)
				os.Exit(0)
			}
		}
		fmt.Print("\nDo you want to exit?(y or n) : ")
		fmt.Scanln(&choice)
	}
	os.Exit(0)
}
