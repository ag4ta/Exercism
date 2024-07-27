package elon

import "fmt"

// Drive updates car's battery and driven distance based on speed and battery drain
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance returns a string which indicates distance driven
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns a string which indicates the level of battery
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish returns bool value which indicates if a car can finish a race
func (c Car) CanFinish(trackDistance int) bool {
	return c.battery/c.batteryDrain*c.speed >= trackDistance
}
