package main

import (
	"fmt"
	"time"
	"strings"
)

type PlantType string
const (
	hydro PlantType ="Hydro"
	wind PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string
const (
	active PlantStatus = "Active"
	inactive PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity float64
	plantStatus PlantStatus
}

type PowerGrid struct {
	load float64
	plants []PowerPlant
}

func GeneratePowerPlantReport(capacitiess... float64)  {
	for idx, cap := range capacitiess {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}

func GeneratePowerGridReport(plants []int, plantCapacitiess []float64, gridLoad float64, utilization *float64, capacity *float64)  {
	totalCapacity := 0.
	for _, cap := range plantCapacitiess {
		totalCapacity += cap
	}

	for _, plantID := range plants {
		*capacity += plantCapacitiess[plantID]
	}
	*utilization = *capacity/totalCapacity
	fmt.Printf("%-20s%.0f\n", "capacity", *capacity)
	fmt.Printf("%-20s%.0f\n", "load", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "utilization", *utilization*100)
}

func main() {

	plants := []PowerPlant{
		PowerPlant{hydro, 300,active},
		PowerPlant{wind, 30,active},
		PowerPlant{wind, 25,inactive},
		PowerPlant{wind, 35,active},
		PowerPlant{solar, 45,unavailable},
		PowerPlant{solar, 40,inactive},
	}

	grid :=PowerGrid{300, plants}
/*
	plantCapacitiess := [] float64 {30,30,30,60,60,100}
	activePlants := []int{0,1,}
	gridLoad := 75.
	utilization := 1.
	capacity := 0.
*/
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	for {
		fmt.Print("Please choose an option or press 'x' to exit: ")
		var option string
		fmt.Scanln(&option)
		if option == "x" {
			break
		}
		time.Sleep(1000 * time.Millisecond)

		switch option {
		case "1":
			//GeneratePowerPlantReport(plantCapacitiess...)
			grid.generatePlantReport()
		case "2":
			//GeneratePowerGridReport(activePlants, plantCapacitiess, gridLoad, &utilization, &capacity)
		default:
			fmt.Println("Invalid input")
		}
	}
}




func (pg *PowerGrid) generatePlantReport (){
	for idx, p := range pg.plants{
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))

		fmt.Printf("%-20s%.0f\n", "capacity", p.capacity)
		fmt.Printf("%-20s\n", "Status", p.plantStatus)
		fmt.Printf("%-20s\n", "Type", p.plantType)
		println()
	}
}



func main3()  {
	mp := messagePrinter{"hi test"}
	mp.printMessage()
	mp2 := enhancedMsgPrinter{}
	mp2.message = "mp2 enhenced"
	mp2.printMessage()

	mp3 := enhancedMsgPrinter{messagePrinter{"mp3"}}
	mp3.printMessage()

}

type messagePrinter struct {

	message string
}

func (mp *messagePrinter) printMessage(){
fmt.Println(mp.message)
}

type enhancedMsgPrinter struct {
	messagePrinter
}