package main

import (
	"container/heap"
	"fmt"
	"log"
	"math"
)

func main() {
	var numberOfBox int

	fmt.Println("Enter the number of box: ")
	_, err := fmt.Scanf("%d", &numberOfBox)
	if err != nil {
		log.Fatal("Bad input")
	}
	boxWeights := make([]int, numberOfBox)
	for i :=range boxWeights {
		fmt.Printf("Enter weight of box %d: ", i + 1 )
		_, err = fmt.Scanf("%d", &boxWeights[i])
		if err != nil {
			log.Fatal("Bad input")
		}
	}
	var numberOfTrucks int
	fmt.Println("Enter number of Truck")
	_, err =fmt.Scanf("%d", &numberOfTrucks)
	var trucks []*Truck
	for i :=0 ;i < numberOfTrucks; i++ {
		trucks = append(trucks, &Truck{
			weight: 0,
			index: i,
			ID: i,
		})
	}
	pq := make(PriorityQueue, len(trucks))

	i :=0
	for _ , value:= range trucks {
		pq[i] = value
		i++
	}
	heap.Init(&pq)

	// Assign box to lightest truck
	for _ , weight := range boxWeights {
		truck := heap.Pop(&pq).(*Truck)
		truck.weight += weight
		heap.Push(&pq, truck )
	}

	maxWeight := math.MinInt64
	var truckID int
	for pq.Len() >0 {
		truck := heap.Pop(&pq).(*Truck)
		fmt.Printf("Truck id: %d, Truck weight: %d\n", truck.ID, truck.weight)
		if truck.weight > maxWeight {
			maxWeight = truck.weight
			truckID = truck.ID
		}

	}
	fmt.Printf("Max weight on the Truck %d: %d\n", truckID, maxWeight)



}