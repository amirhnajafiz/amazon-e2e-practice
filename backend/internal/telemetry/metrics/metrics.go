package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics are being used to monitor the service status.
type Metrics struct {
	requests    *prometheus.CounterVec
	statusCodes *prometheus.CounterVec
	latency     *prometheus.HistogramVec
}

// NewMetrics returns a metrics struct with registered promehteus metrics.
func NewMetrics() *Metrics {
	return &Metrics{
		requests: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "total_requests",
			Help: "total number of requests",
		}, []string{"endpoint", "method"}),
		statusCodes: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "status_codes",
			Help: "counts of status codes returned by the service",
		}, []string{"endpoint", "method", "status_code"}),
		latency: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "requests_latency",
			Help: "requests handling latency",
		}, []string{"endpoint", "method"}),
	}
}

func (m *Metrics) AddRequest(endpoint, method string) {
	m.requests.With(prometheus.Labels{"endpoint": endpoint, "method": method}).Add(1)
}

func (m *Metrics) AddStatusCode(endpoint, method, statusCode string) {
	m.statusCodes.With(prometheus.Labels{"endpoint": endpoint, "method": method, "status_code": statusCode}).Add(1)
}

func (m *Metrics) ObserveLatency(endpoint, method string, value float64) {
	m.latency.With(prometheus.Labels{"endpoint": endpoint, "method": method}).Observe(value)
}
