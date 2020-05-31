package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

//
//
type timeDataCollector struct {
	data []*timeData
}

type timeData struct {
	date time.Time
	value float64
}

func (t *timeDataCollector) add(value float64) {
	t.data = append(t.data, &timeData{
		date:  time.Now(),
		value: value,
	})
}

func (t *timeDataCollector) getRange(from, to time.Time) []*timeData {
	var res []*timeData
	for _, item := range t.data {
		if item.date.After(from) && item.date.Before(to) {
			res = append(res, item)
		}
	}

	return res
}

func (t *timeDataCollector) avg(from, to time.Time) float64 {
	rang := t.getRange(from, to)
	sum := 0.0
	for _, r := range rang {
		sum = sum + r.value
	}

	return sum / float64(len(rang))
}

//
//
func getCurrentTemp() (float64, error) {
	out, err := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp").Output()
	if err != nil {
		return 0, err
	}
	re := regexp.MustCompile(`\d+\.\d+`)
	res := string(re.Find(out))

	return strconv.ParseFloat(res, 64)
}

//
//
func startSever(port int, db *timeDataCollector) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		temp, err := getCurrentTemp()
		tempStr := ""

		if err != nil {
			tempStr = "n/a"
		} else {
			tempStr = fmt.Sprintf("%.2f", temp)
		}
		to := time.Now()
		from := to.Add(-time.Hour)
		rang := db.getRange(from, to)
		rangStr := ""
		for _, r := range rang {
			rangStr = rangStr + fmt.Sprintf("<b>%s</b>: %.2f<br>", r.date, r.value)
		}
		avgStr := fmt.Sprintf("%.2f", db.avg(from, to))

		writer.Write([]byte(fmt.Sprintf(`
				<html>
					<body>
				 	<p>Current temp: %s</p>
					<br>
					<p>Avg temp: %s</p>
					<br>
					%s
					</body>
				</html>
			`, tempStr, avgStr, rangStr)))
		writer.Header().Add("Content-Type", "text/html; charset=utf-8")
		writer.WriteHeader(200)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

//
//
func collectTempData(db *timeDataCollector, every time.Duration) {
	for {
		temp, err := getCurrentTemp()
		if err != nil {
			log.Print(err)
		} else {
			db.add(temp)
		}
		time.Sleep(every)
	}
}

//
//
func main() {
	timeDB := &timeDataCollector{}
	go collectTempData(timeDB, time.Second * 10)
	if err := startSever(8081, timeDB); err != nil {
		panic(err)
	}
}
