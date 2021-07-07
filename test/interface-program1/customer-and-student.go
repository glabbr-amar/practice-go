package main

var customerMap map[int]Customer
var studentMap map[int]Student

type Customer struct {
	name         string
	id           int
	address      string
	mobileNumber int64
}
type Student struct {
	name         string
	id           int
	address      string
	mobileNumber int64
}

type StudentAndCustomer interface {
	Create(name string, id int, address string, mobileNumber int64)
}

func (s Student) Create(name string, id int, address string, mobileNumber int64) {
	storeStudent := Student{
		name:         name,
		id:           id,
		address:      address,
		mobileNumber: mobileNumber,
	}
	studentMap[id] = storeStudent
}

func (c Customer) Create(name string, id int, address string, mobileNumber int64) {
	storeCustomer := Customer{
		name:         name,
		id:           id,
		address:      address,
		mobileNumber: mobileNumber,
	}
	customerMap[id] = storeCustomer
}

func GetCustomerById(id int) Customer {
	return customerMap[id]
}
