package api

import (
	"encoding/json"
	"fmt"
	"github.com/marcuzy/pimonit/core"
	"github.com/marcuzy/pimonit/core/actions"
	"log"
	"net/http"
	"time"
)

//
//
func StartSever(port int, ts core.TimeSeries, ch core.ChartsGenerator) error {
	http.HandleFunc("/chart/download", func(writer http.ResponseWriter, request *http.Request) {
		to := time.Now()
		from := to.Add(-time.Minute * 15)
		buf, err := actions.GenerateChartPNG(ts, ch, from, to)
		if err != nil {
			log.Print(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Add("Content-Type", "image/png")
		writer.Write(buf.Bytes())
	})

	http.HandleFunc("/chart", func(writer http.ResponseWriter, request *http.Request) {
		to := time.Now()
		from := to.Add(-time.Minute * 15)
		items, err := actions.GetRange(ts, from, to)
		if err != nil {
			log.Print(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		js := make([]map[string]interface{}, len(items))
		for i, item := range items {
			js[i] = map[string]interface{}{
				"time":  item.Date.Format(time.RFC3339),
				"value": item.Value,
			}
		}
		b, err := json.Marshal(&js)
		if err != nil {
			log.Print(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Add("Content-Type", "application/json")
		writer.Write(b)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
