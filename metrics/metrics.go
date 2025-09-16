package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var metricsList = []prometheus.Collector{
	timeResponses,
	httpRequestsTotal,
}

var (
	timeResponses = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Time taken to respond to requests in seconds.",
			Buckets: []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"method", "endpoint", "status_code"})

	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed.",
		},
		[]string{"method", "endpoint", "status_code"},
	)
)

func Start() {
	prometheus.MustRegister(metricsList...)
}

func UpdTimeResponse(method, endpoint, status_code string, spentTime float64) {
	timeResponses.WithLabelValues(method, endpoint, status_code).Observe(spentTime)
}

func IncHttpRequestsTotal(method, endpoint, status_code string) {
	httpRequestsTotal.WithLabelValues(method, endpoint, status_code).Inc()
}
