package metrics

import (
	"t-mk-opentrace/pkg/mq/consumer"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// consumerDuration 消费时间
	consumerDuration *prometheus.HistogramVec
	// consumerSuccessCounter 成功数量
	consumerSuccessCounter *prometheus.CounterVec
	// consumerFailedCounter 失败量
	consumerFailedCounter *prometheus.CounterVec
	// consumerPayloadSize payload大小
	consumerPayloadSize *prometheus.SummaryVec
	// productDuration 生产时间
	productDuration *prometheus.HistogramVec
	// productSuccessCounter 成功数
	productSuccessCounter *prometheus.CounterVec
	// productFailedCounter 是失败数
	productFailedCounter *prometheus.CounterVec
	// productPayloadSize payload大小
	productPayloadSize *prometheus.SummaryVec
)
var labels = []string{"payload"}

// RegisterConsumerMetrics 注册指标
func RegisterConsumerMetrics(namespace string) {
	metrics := make([]prometheus.Collector, 0, 4)
	// 消费指标
	consumerDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "kafka_consumer_duration_seconds",
		Help:      "kafka consumer latencies in seconds.",
	}, labels)
	metrics = append(metrics, consumerDuration)
	consumerSuccessCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "kafka_consumer_success_total",
		Help:      "total number of kafka consume successful",
	}, labels)
	metrics = append(metrics, consumerSuccessCounter)
	consumerFailedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "kafka_consumer_failed_total",
		Help:      "total number of kafka consume failed",
	}, labels)
	metrics = append(metrics, consumerFailedCounter)
	consumerPayloadSize = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Name:      "kafka_consumer_payload_byte_size",
		Help:      "kafka consumer payload byte size once",
	}, labels)
	metrics = append(metrics, consumerPayloadSize)
	prometheus.MustRegister(metrics...)
}

// RegisterProductMetrics 注册生产指标
func RegisterProductMetrics(namespace string) {
	metrics := make([]prometheus.Collector, 0, 4)
	// 生产指标
	productDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "kafka_product_duration_seconds",
		Help:      "kafka product latencies in seconds.",
	}, labels)
	metrics = append(metrics, productDuration)

	productSuccessCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "kafka_product_success_total",
		Help:      "total number of kafka product successful",
	}, labels)
	metrics = append(metrics, productSuccessCounter)

	productFailedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "kafka_product_failed_total",
		Help:      "total number of kafka product failed",
	}, labels)
	metrics = append(metrics, productFailedCounter)
	productPayloadSize = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Name:      "kafka_product_payload_byte_size",
		Help:      "kafka product payload byte size once",
	}, labels)
	metrics = append(metrics, productPayloadSize)
	prometheus.MustRegister(metrics...)
}

// IncMetricsByConsumer 统计消费服务
func IncMetricsByConsumer(payload *consumer.DefaultHandlerParam, startTime time.Time, err error) func() {
	size := len(payload.Payload.Value)
	consumerPayloadSize.WithLabelValues(labels...).Observe(float64(size))
	return func() {
		if err != nil {
			consumerFailedCounter.WithLabelValues(labels...).Inc()
		} else {
			consumerSuccessCounter.WithLabelValues(labels...).Inc()
		}
		time.Now()
		consumerDuration.WithLabelValues(labels...).Observe(time.Since(startTime).Seconds())
	}
}

// IncMetricsByProduct 统计生产服务
func IncMetricsByProduct(payload *consumer.DefaultHandlerParam, startTime time.Time, err error) func() {
	size := len(payload.Payload.Value)
	productPayloadSize.WithLabelValues(labels...).Observe(float64(size))
	return func() {
		if err != nil {
			productFailedCounter.WithLabelValues(labels...).Inc()
		} else {
			productSuccessCounter.WithLabelValues(labels...).Inc()
		}
		time.Now()
		productDuration.WithLabelValues(labels...).Observe(time.Since(startTime).Seconds())
	}
}
