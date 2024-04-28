package main

func main() {
	// CREATING RESERVATION

	// reservation := models.Reservation{
	// 	ID:         1231231, // Provide a unique ID for the reservation
	// 	ReserverId: "John Doe",
	// 	Date:       "11/11/2011",
	// 	StartTime:  "01:00 AM",
	// 	EndTime:    "01:50 AM",
	// 	Recursive:  true,
	// }
	// reservation.CreateReservation()

	// SELECTING BY TIME

	// reservations := models.GetReservationsByTime("01:00 AM", "01:50 AM")
	// for _, reservation := range reservations {
	// 	fmt.Println(reservation)
	// }

	// SELECTING BY RESERVER_ID

	// reservations := models.GetReservationsByReserverID("John Doe")
	// for _, reservation := range reservations {
	// 	fmt.Println(reservation)
	// }

	// EDITING AN ENTRY

	// reservationID := 1231231
	// reservation, err := models.GetReservationByID(reservationID)
	// if err != nil {
	// 	panic(err)
	// }

	// // Print the original reservation
	// fmt.Println("Original Reservation:")
	// fmt.Println(reservation)

	// // Modify some fields of the reservation
	// reservation.ReserverId = "Alice"
	// reservation.Date = "12/12/2012"

	// // Call the EditReservation function to update the reservation
	// newReservation := models.Reservation{
	// 	ID:         reservation.ID,
	// 	ReserverId: reservation.ReserverId,
	// 	Date:       reservation.Date,
	// 	StartTime:  reservation.StartTime,
	// 	EndTime:    reservation.EndTime,
	// 	Recursive:  reservation.Recursive,
	// }
	// fmt.Println("Editing Reservation...")
	// reservation.EditReservation(newReservation)
	// fmt.Println("Reservation Edited.")

	// // Print the updated reservation
	// fmt.Println("Updated Reservation:")
	// updatedReservation, err := models.GetReservationByID(reservationID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(updatedReservation)

	// DELETING A RESERVATION
	// reservationID := 1231231
	// reservation, err := models.GetReservationByID(reservationID)
	// if err != nil {
	// 	panic(err)
	// }
	// reservation.DeleteReservation()

	// CREATING

	// t := models.Teacher{
	// 	Id:         5,
	// 	Name:       "John Doe",
	// 	Department: "Computer Science",
	// }
	// t.CreateTeacher()

	// // SELECTING by id
	// teacher := models.Teacher{Id: 4}
	// teacher.Search()
	// fmt.Println(teacher)

	// // Selecting by name

	// // Single
	// teacher := models.Teacher{Name: "John Doe"}
	// teacher.Search()
	// fmt.Println(teacher)

	// // Multiple
	// var teachers []models.Teacher
	// teacher = models.Teacher{Name: "John Doe"}
	// err := teacher.Search(&teachers)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(teachers)

	// router := gin.Default()
	// t_r := router.Group("/teacher")
	// {
	// 	t_r.POST("/login", handler.HandlerTeacherLogin)
	// }
	// router.LoadHTMLGlob("templates/*")

	// router.GET("/reserve", handler.GetreserveHandler)
	// router.GET("/login", handler.GetLoginHandler)
	// router.Run(":80")
}
