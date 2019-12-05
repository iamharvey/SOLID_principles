# SOLID Principles

## SOLID And History
SOLID is a mnemonic acronym for five software design principles:
- **S** - Single Responsibility Principle
- **O** - Open Closed Principle
- **L** - Liskov Substitution Principle
- **I** - Interface Segregation Principle
- **D** - Dependency Inversion Principle

These principles were first introduced by Robert C.Martin(Uncle Bob), According to SOLID Wikipedia page, four of the 
SOLID principles (OLID) were first introduced together in one of the Uncle Bob's paper 
([Design Principles and Design Patterns](https://web.archive.org/web/20150906155800/http://www.objectmentor.com/resources/articles/Principles_and_Patterns.pdf) 
published on [objectmentor.com](www.objectmentor.com) in 2000. The S principle was introduced in his famous book 
'Agile Software Development, Principles, Patterns, and Practices' which was published few years later. Uncle Bob mentioned 
that the S principle was inspired by the idea of cohesion introduced in two books: 
'Structured Analysis and System Specification' (Tom DeMarco) and 
'The Practical Guide to Structured Systems Design' (Meilir Page-Jones).

Uncle Bob also introduced several other design principles, particularly for object oriented software design approach. 
But he is not the first one that promotes SOLID. The SOLID acronym was introduced by Michael Feathers.

Uncle Bob

![](https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fwww.tomordonez.com%2Fimages%2Frobert-uncle-bob-martin.jpg&f=1&nofb=1)

Michael Feathers

![](http://agilesingapore.org/img/mfeathers.jpg)

<br>

## Why
The core of SOLID is to **embrace changes and collaboration between software modules**, to address the following typical issues in software development:
- rigidity (difficult to make change), 
- fragility（change brings service and development interruption)
- immobility（poor re-usability）, 
- viscosity（highly design coupling, difficult to apply changes)

Although SOLID principles were first introduced for object oriented software development, it is also useful for agile 
approach. Because its core value is applicable universally. Nowadays, SOLID is still very popular and plays an important 
role in junior developer training and code review.

<br>

## SOLID In Details

### S - Single Responsibility Principle
> A class should have one and only one reason to change, meaning that a class should only one job.

**Not So Good**
```
type Drone interface {
	fly()
}

type DroneX struct {
	Drone
	name string
}

func (dr *DroneX) fly() {
  fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
	dr.checkBattery()
	dr.checkPropeller()
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
```

<br>

**Good**
```
type Drone interface {
	fly()
}
 
type DroneX struct {
	name string
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

```

<br>

### O - Open Closed Principle
> Objects or entities should be open for extension, but closed for modification.

**Not So Good**
```
...
func (dr *DroneX) fly() {
  fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
	dr.prepare()
	dr.takeOff()
	dr.healthCheck()
	if dr.model == "Y" {
		// do some spinning around
		...
	}
}

```

<br>

**Good**
```
...
// extend Drone
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
```

<br>

### L - Liskov Substitution Principle

> Let q(x) be a property provable about objects of x of type T. Then q(y) should be provable for objects y of type S where S is a subtype of T.

What it actually means: 
    keep your method behaviour definition (name and functioning) as same as your parent. 

**Problematic**
```
...
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

// extend DroneX
type DroneZ struct {
	DroneX
}

func (dr *DroneZ) flyWithPCW() {
	fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
	dr.prepare()
	dr.takeOff()
	dr.healthCheck()
	dr.pirouettingCW()
}

func (dr *DroneZ) pirouettingCW() {
	fmt.Println("[flying] I am pirouetting clockwise ... ")
}

func getDrones() []Drone {
	return []Drone {&DroneY{}, &DroneZ{}, &DroneX{}}
}

func main() {
	for _, dr := range getDrones() {
		dr.fly()
	}
}
```

<br>

**Correction**
```
...
func (dr *DroneZ) fly() {
	fmt.Printf(">>> flying drone[%s] ... \n", dr.name)
	dr.prepare()
	dr.takeOff()
	dr.healthCheck()
	dr.pirouettingCW()
}
...
```

<br>

### I - Interface Segregation Principle
> A client should never be forced to implement an interface that it doesn't use or clients shouldn't be forced to depend on methods they do not use.

What it actually means: 
    do not modify your interface, make a new interface.

E.g., we wanna enable team-up for drones.

**Not So Good**
```
type Drone interface {
	fly()
	add()
}
```

<br>

**Good**
```
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
...
```

<br>

### D - Dependency Inversion Principle
> Entities must depend on abstractions not on concretions. It states that the high level module must not depend on the low level module, but they should depend on abstractions.

What it actually means: 
    high-level module should not dependent on low-level ones.
    
what we can do about:
    make dependency inversion (use dependency injection)

E.g., 
```
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
...
func main() {
	c1 := &Controller{cType: "NBCFB X-20B"}
	dr1 := NewDroneX("Hummingbird#1", c1)
	dr1.fly()
	c2 := &Controller{cType: "NBCFB W-93S"}
	dr2 := NewDroneX("Hummingbird#2", c2)
	dr2.fly()
}
```

<br>

## Summary
Why we need SOLID?
- embracing changes
- embracing changes
- embracing changes

There is a new saying: important thing should repeated three times.

What key moves you need to recall:
- inheritance
- override
- interfacing

How to check if my SOLID practice is on the right track? Check if your program is:
- easy to read  and understand **by others**
- easy to change **by others**
- easy to extend **by others**
- easy to collaborate software modules developed **by others**
- easy to test

## Reference:
1. [Design Principles and Design Patterns](https://web.archive.org/web/20150906155800/http://www.objectmentor.com/resources/articles/Principles_and_Patterns.pdf)
2. [SOLID, wikipedia](https://en.wikipedia.org/wiki/SOLID).

