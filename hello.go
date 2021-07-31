package main

import (
   "fmt"
   "math/rand"
   "time"
)

func main() {
   const hotelName = "Gopher Paris Inn"
   const totalRooms = 134
   const firstRoomNumber = 110

   rand.Seed(time.Now().UTC().UnixNano())
   roomsOccupied := rand.Intn(totalRooms)
   roomsAvailable := totalRooms - roomsOccupied

   occupancyRate := occupancyRate(roomsOccupied, totalRooms)
   occupancyLevel := occupancyLevel(occupancyRate)

   fmt.Println("Hotel:", hotelName)
   fmt.Println("Number of rooms", totalRooms)
   fmt.Println("Rooms available", roomsAvailable)
   fmt.Println("                  Occupancy Level:", occupancyLevel)
   fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate)

   if roomsAvailable > 0 {
      fmt.Println("Rooms:")
      for i := 0; roomsAvailable > i; i++ {
         roomNumber := firstRoomNumber + i
         size := rand.Intn(6) + 1
         nights := rand.Intn(10) + 1
         printRoomDetails(roomNumber, size, nights)
      }
   } else {
      fmt.Println("No rooms available for tonight")
   }

}


func printRoomDetails(roomNumber, size, nights int){
   fmt.Println(roomNumber, ":", size, nights, "nights" )
}

func occupancyLevel(occupancyRate float32) string{
   if occupancyRate > 70 {
      return "High"
   } else if occupancyRate > 20 {
      return "Medium"
   } else {
      return "Low"
   }
}

func occupancyRate(roomsOccupied int, totalRooms int) float32 {
   return (float32(roomsOccupied) / float32(totalRooms)) *100
}
   