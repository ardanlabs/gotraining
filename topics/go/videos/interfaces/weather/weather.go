package weather

import (
	"log"
)

type Celsius float64

func (c Celsius) Unit() string {
	return "°C"
}

func (c Celsius) Amount() float64 {
	return float64(c)
}

func Current() (Celsius, error) {
	c := Celsius(23.4) // FIXME: get actual weather
	u := c.Unit()
	log.Printf("INFO: the temperature is %.1f %s", c, u)
	return c, nil
}
