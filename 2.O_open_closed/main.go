package main

import "fmt"

type Drone interface {
	fly()
}

type DroneX struct {
	Drone
	name string
	model string
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

// extend DroneX
type DroneY struct {
	DroneX
}

func (dr *DroneY) fly() {
	fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
	dr.prepare()
	dr.takeOff()
	dr.healthCheck()
	dr.spinAround()
}

func (dr *DroneY) spinAround() {
	fmt.Println("[flying] I am spinning around ... ")
}

func main() {
	dr := DroneY{}
	dr.name = "Falcon"
	dr.model = "Y"
	dr.fly()
}



