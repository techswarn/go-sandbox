package main

import "fmt"

type iVehicle interface {
    setName(name string)
    setSeats(seats int)
    getName() string
    getSeats() int
}

type vehicle struct {
    name  string
    seats int
}

func (v *vehicle) setName(name string) {
    v.name = name
}

func (v *vehicle) getName() string {
    return v.name
}

func (v *vehicle) setSeats(seats int) {
    v.seats = seats
}

func (v *vehicle) getSeats() int {
    return v.seats
}

type ferrari struct {
    vehicle
}

func newFerrari() iVehicle {
    return &ferrari{
        vehicle: vehicle{
            name:  "Ferrari",
            seats: 2,
        },
    }
}

type renault struct {
    vehicle
}

func newRenault() iVehicle {
    return &renault{
        vehicle: vehicle{
            name:  "Renault",
            seats: 6,
        },
    }
}

func getVehicle(vehicleType string) (iVehicle, error) {
    if vehicleType == "ferrari" {
        return newFerrari(), nil
    }
    if vehicleType == "renault" {
        return newRenault(), nil
    }
    return nil, fmt.Errorf("Wrong vehicle type passed")
}

func main() {
    ferrari, _ := getVehicle("ferrari")
    renault, _ := getVehicle("renault")
    printDetails(ferrari)
    printDetails(renault)
}

func printDetails(v iVehicle) {
    fmt.Printf("Name: %s", v.getName())
    fmt.Println()
    fmt.Printf("Seats: %d", v.getSeats())
    fmt.Println()
}