package example_7

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")



//运行参数通过 -period 50ms指定
func Test_7_4() {

	flag.Parse()

	fmt.Printf(" Sleeping for %v...", *period)

	fmt.Println(period.Seconds())
	fmt.Println(*temp)

	//time.Sleep(*period)

	fmt.Println()

	var celsius celsiusFlag

	s := "37°C"
	celsius.Set(s)
	fmt.Printf("celsiusFlag.Set invalid is %q result is %f", s, celsius.Celsius)
}

//转换
type Celsius float64    //摄氏度
type Fahrenheit float64 // 华氏温度

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	_, err := fmt.Sscanf(s, "%f%s", &value, &unit)

	if err != nil {
		return fmt.Errorf("invalid temperature %q", err.Error())
	}
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil

	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CelsiusFlag(name string, value Celsius, usage string) *Celsius  {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp  =  CelsiusFlag("temp", 20, "the temperature")