package main
import "fmt"

type vehicle struct {
  wheels         int
  seats          int
  hasLuggageRack bool
}

func newVehicleBuilder() *vehicle {
  return &vehicle{
    wheels: 0,
    seats: 0,
    hasLuggageRack: false,
  }
}

func(v *vehicle) WithWheels(wheels int) *vehicle {
  v.wheels = wheels
  v.seats = wheels
  return v
}

func(v *vehicle) WithSeats(seats int) *vehicle {
  if v.wheels > 2 {
    v.seats = seats
  }
  return v
}

func(v *vehicle) WithLuggageRack() *vehicle {
  if(v.wheels > 2) {
    v.hasLuggageRack = true
  }
  return v
}

func(v *vehicle) PrintVehicleDetails() {
  fmt.Printf(
    "Car Details -\n\tWheels: %d\n\tSeats: %d\n\tHas Luggage Rack: %v\n",
    v.wheels, 
    v.seats, 
    v.hasLuggageRack,
  )
}

func getSmallCar() *vehicle {
  return newVehicleBuilder().WithWheels(4)
}

func getLargeCar() *vehicle {
  return newVehicleBuilder().WithWheels(4).WithSeats(6).WithLuggageRack()
}

func main() {
    smallCar := getSmallCar()
    smallCar.PrintVehicleDetails()

    largeCar := getLargeCar()
    largeCar.PrintVehicleDetails()
}