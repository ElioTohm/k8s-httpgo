package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namespace = "k8shttpgo"

var (
	IndexerSuccessfulWrite = promauto.NewCounterVec(prometheus.CounterOpts{
		Name:      "event_1",
		Help:      "random event",
		Namespace: namespace,
	}, []string{"hostname", "timestamp"})
)
