package user

import(
	"encoding/json"
	address "github.com/AdityaSubrahmanyaBhat/dashDB/models/Address"
)

type User struct{
	Name string
	Age json.Number
	Company string
	Address address.Address
}