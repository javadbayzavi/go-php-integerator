/*
*** @author: Javad Bayzavi
*** @version: 1.0.1
*** @email: javadbayzavi@gmail.com
*** @year: 2021
*/

package Entities

type Entity struct {
	Id int `json:"Id"`
}

func (this *Entity) getId() int {
	return this.Id
}
