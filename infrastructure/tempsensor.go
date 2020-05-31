package infrastructure

import (
	"errors"
	"github.com/marcuzy/pimonit/core/interfaces"
	"os/exec"
	"regexp"
	"strconv"
)

type (
	piTempSensor struct {
	}
)

func NewPiTempSensor() interfaces.TempSensor {
	return &piTempSensor{}
}

func (_ piTempSensor) CurrentTemperature(units interfaces.TemperatureUnits) (float64, error) {
	out, err := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp").Output()
	if err != nil {
		return 0, err
	}
	re := regexp.MustCompile(`\d+\.\d+`)
	res := string(re.Find(out))

	cels, err := strconv.ParseFloat(res, 64)
	if err != nil {
		return 0, err
	}
	switch units {
	case interfaces.TemperatureUnitsCelsius:
		return cels, nil
	default:
		return 0, errors.New("unsupported units")
	}
}
