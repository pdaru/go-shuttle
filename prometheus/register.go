package prometheus

import (
	"github.com/pdaru/go-shuttle/prometheus/listener"
	"github.com/pdaru/go-shuttle/prometheus/publisher"
	"github.com/prometheus/client_golang/prometheus"
)

func Register(registerer prometheus.Registerer) {
	listener.Metrics.Init(registerer)
	publisher.Metrics.Init(registerer)
}
