package ginprometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler() gin.HandlerFunc {
	return HandlerFor(prometheus.DefaultGatherer)
}

func HandlerFor(g prometheus.Gatherer) gin.HandlerFunc {
	return gin.WrapH(promhttp.HandlerFor(g, promhttp.HandlerOpts{}))
}
