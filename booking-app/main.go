package main //		mUST belong to a package

import "fmt" //comtainers of various applications check the documentations for packages
//entry point where do i start excuting the code
//should be a function

func main() {
	// var conferenceName = "Go Conference" // when creating a variable with a certain value it must used same to importing package
	conferenceName := "Go Conference" //Alternative syntax cannot declare syntax cannot be used when explicitly defining a variable
	const conferenceTickets = 50
	var remainingTickets uint = 50

	/*fmt.Println("Welcome to our", conferenceName, "booking application")*/
	//Print formated data
	fmt.Printf("Welcome to our %v booking application\n", conferenceName) //fmt package link
	// fmt.Println("We have", conferenceTickets, "tickets and",  remainingTickets, "are still available")
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets uint
	//ask user for their name

	/*READING USER INPUT

	fmt.Scan()
	*/

	fmt.Println("Enter your first name")

	fmt.Scan(&userName) //POINTER  variable that points the memory address on a variable

	/*POINTER
	where the value gets stored in the memory
	*/

	userTickets = 2

	fmt.Println("Enter no of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("The remaining tickets %v\n", remainingTickets)

	/*
		Data types
		- creating a variable and assign it a value immediately the go imply the type
		Else a type must be defined explicitly
		learn types
		- %v place holder variable
		- %T shows the type of variable

		#------Numeric Datatypes
		type uint  will not accept negative variables

		#----- Float DataTypes

		#--Alternative syntax


	*/

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	fmt.Printf("User %T booked %T tickets.\n", userName, userTickets)
	fmt.Println(&userName, &userTickets) //pointer


	/*
	Arrays and Slices 
	-size  [50] 
	-datatypes



	slices
	an abstracton of array

	more flexible and powerful

	index-based and have a size and can be resized when needed

	var booking[] string

	booking.append(bookings, 'firstName')



	
	*/

	// var bookings = [50] string{} declare and assing

	// var booking [50] string

	// booking[0] = userName



	/*Loops
		for {  indefinate loop
		
		
		}


		for index booking:= bookins {
		
		}
	
	
	*/


	// _ blank indentifer in loods 




	/*if  commitition {

	break
	
	}

	else{
	}




	for loop conditional
*/
}
