package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

func TestUpdTimeResponse(t *testing.T) {
	prometheus.DefaultRegisterer = prometheus.NewRegistry()
	Start()

	tests := []struct {
		name       string
		method     string
		endpoint   string
		statusCode string
		spentTime  float64
	}{
		{
			name:       "GET request with 200 status",
			method:     "GET",
			endpoint:   "/api/item/123",
			statusCode: "200",
			spentTime:  0.05,
		},
		{
			name:       "GET request with 404 status",
			method:     "GET",
			endpoint:   "/api/item/121",
			statusCode: "404",
			spentTime:  0.02,
		},
		{
			name:       "Very fast response",
			method:     "GET",
			endpoint:   "/api/item/122",
			statusCode: "200",
			spentTime:  0.001,
		},
		{
			name:       "Slow response",
			method:     "GET",
			endpoint:   "/api/reports",
			statusCode: "200",
			spentTime:  2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdTimeResponse(tt.method, tt.endpoint, tt.statusCode, tt.spentTime)

			metric := &dto.Metric{}
			histogram := timeResponses.WithLabelValues(tt.method, tt.endpoint, tt.statusCode)
			err := histogram.(prometheus.Metric).Write(metric)
			if err != nil {
				t.Fatalf("Failed to write metric: %v", err)
			}

			if metric.Histogram.GetSampleCount() == 0 {
				t.Error("Expected sample count to be greater than 0")
			}

			if metric.Histogram.GetSampleSum() < tt.spentTime {
				t.Errorf("Expected sample sum to be at least %f, got %f", tt.spentTime, metric.Histogram.GetSampleSum())
			}
		})
	}
}

func TestIncHttpRequestsTotal(t *testing.T) {
	prometheus.DefaultRegisterer = prometheus.NewRegistry()
	Start()

	tests := []struct {
		name       string
		method     string
		endpoint   string
		statusCode string
		increments int
	}{
		{
			name:       "Single GET request",
			method:     "GET",
			endpoint:   "/api/item/120",
			statusCode: "200",
			increments: 1,
		},
		{
			name:       "Error requests",
			method:     "GET",
			endpoint:   "/api/item/123",
			statusCode: "404",
			increments: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.increments; i++ {
				IncHttpRequestsTotal(tt.method, tt.endpoint, tt.statusCode)
			}

			metric := &dto.Metric{}
			counter := httpRequestsTotal.WithLabelValues(tt.method, tt.endpoint, tt.statusCode)
			err := counter.(prometheus.Metric).Write(metric)
			if err != nil {
				t.Fatalf("Failed to write metric: %v", err)
			}

			expectedValue := float64(tt.increments)
			actualValue := metric.Counter.GetValue()

			if actualValue != expectedValue {
				t.Errorf("Expected counter value %f, got %f", expectedValue, actualValue)
			}
		})
	}
}

func TestMetricsIntegration(t *testing.T) {
	prometheus.DefaultRegisterer = prometheus.NewRegistry()
	Start()

	requests := []struct {
		method     string
		endpoint   string
		statusCode string
		duration   float64
	}{
		{"GET", "/api/users", "200", 0.05},
		{"GET", "/api/users", "200", 0.03},
		{"POST", "/api/users", "201", 0.15},
		{"GET", "/api/users/999", "404", 0.02},
	}

	for _, req := range requests {
		UpdTimeResponse(req.method, req.endpoint, req.statusCode, req.duration)
		IncHttpRequestsTotal(req.method, req.endpoint, req.statusCode)
	}

	counterMetric := &dto.Metric{}
	counter := httpRequestsTotal.WithLabelValues("GET", "/api/users", "200")
	err := counter.(prometheus.Metric).Write(counterMetric)
	if err != nil {
		t.Fatalf("Failed to write counter metric: %v", err)
	}

	if counterMetric.Counter.GetValue() != 2 {
		t.Errorf("Expected 2 GET /api/users 200 requests, got %f", counterMetric.Counter.GetValue())
	}

	histogramMetric := &dto.Metric{}
	histogram := timeResponses.WithLabelValues("GET", "/api/users", "200")
	err = histogram.(prometheus.Metric).Write(histogramMetric)
	if err != nil {
		t.Fatalf("Failed to write histogram metric: %v", err)
	}

	if histogramMetric.Histogram.GetSampleCount() != 2 {
		t.Errorf("Expected 2 histogram samples for GET /api/users 200, got %d", histogramMetric.Histogram.GetSampleCount())
	}

	expectedSum := 0.08
	actualSum := histogramMetric.Histogram.GetSampleSum()
	if actualSum < expectedSum-0.001 || actualSum > expectedSum+0.001 {
		t.Errorf("Expected histogram sum around %f, got %f", expectedSum, actualSum)
	}
}
