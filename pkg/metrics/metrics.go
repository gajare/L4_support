package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "Total number of HTTP request",
	}, []string{"method", "path"})

	HttpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http_request_duration_seconds",
		Help:      "Duration of HTTP requests",
	}, []string{"method", "path"})

	DatabaseOperationsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "database_operations_total",
		Help: "total number of database operation",
	}, []string{"operation", "table"})
)
