package main

import "log"

func main() {
	amar, ranjan, deepak, satyam := Customer{}, Customer{}, Customer{}, Customer{}
	govind, abhay, anuj, sunny := Student{}, Student{}, Student{}, Student{}
	/*customerList := []StudentAndCustomer{amar,ranjan,deepak,satyam}
	studentList := []StudentAndCustomer{govind,abhay,anuj,sunny}*/
	amar.Create("amar", 0, "manpur", 84946161944)
	ranjan.Create("ranjan", 1, "gaya", 79616494669)
	deepak.Create("deepak", 2, "bihar", 496164941619)
	satyam.Create("satyam", 3, "manpur", 4941694616)

	govind.Create("govind", 0, "manpur", 84946161944)
	abhay.Create("abhay", 1, "gaya", 79616494669)
	anuj.Create("anuj", 2, "bihar", 496164941619)
	sunny.Create("sunny", 3, "manpur", 4941694616)

	log.Println(GetCustomerById(0))

}
