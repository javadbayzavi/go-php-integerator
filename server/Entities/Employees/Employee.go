/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
*/

package Entities

import (
	"go-microservices/lib/Entities"
)

type Employee struct {
	*Entities.Entity
	Name   string `json:"name"` //Define encode / decode mapper in Json format
	Family string `json:"family"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}
