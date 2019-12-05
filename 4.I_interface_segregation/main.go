package main

import (
	"fmt"
)

type Drone interface {
	fly()
}

type Team interface {
	add()
}

type DroneX struct {
	Drone
	Team
	name string
}

func (dr *DroneX) add(mate *Drone) {
	fmt.Printf("[team up] forming a team, adding %v to my team ... \n", mate)
}

func (dr *DroneX) prepare() {
	dr.checkBattery()
	dr.checkPropeller()
}

func (dr *DroneX) fly() {
	fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
	dr.prepare()
	dr.takeOff()
	dr.healthCheck()
}

func (dr *DroneX) checkBattery() {
	fmt.Println("[preparing] checking battery's status ... ")
}

func (dr *DroneX) checkPropeller() {
	fmt.Println("[preparing] checking propellers' status ... ")
}

func (dr *DroneX) takeOff() {
	fmt.Println("[taking off] taking off now ... ")
}

func (dr *DroneX) healthCheck() {
	fmt.Println("[flying] on the air, everything is ok, auto balancing enabled ... ")
}

func main() {
	var dr1 Drone = &DroneX{ name: "Hummingbird"}
	dr1.fly()
	var dr2 = &DroneX{ name: "Hummingbird"}
	dr2.fly()
	dr2.add(&dr1)
}

