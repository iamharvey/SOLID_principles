package main

import "fmt"

type Drone interface {
	fly()
}

type Controller struct {
	cType string
}

type DroneX struct {
	Drone
	name string
	c *Controller
}

func NewDroneX(name string, c *Controller) *DroneX {
	return &DroneX{
		Drone: nil,
		name:  name,
		c:	c,
	}
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
	c1 := &Controller{cType: "NBCFB X-20B"}
	dr1 := NewDroneX("Hummingbird#1", c1)
	dr1.fly()
	c2 := &Controller{cType: "NBCFB W-93S"}
	dr2 := NewDroneX("Hummingbird#2", c2)
	dr2.fly()
}



