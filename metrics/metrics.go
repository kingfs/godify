package metrics

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Metrics 指标收集器
type Metrics struct {
	mu sync.RWMutex
	
	// 请求统计
	TotalRequests    int64
	SuccessfulRequests int64
	FailedRequests   int64
	
	// 响应时间统计
	ResponseTimes []time.Duration
	
	// 错误统计
	ErrorCounts map[string]int64
	
	// 连接统计
	ActiveConnections int64
	TotalConnections  int64
	
	// 配置
	enabled bool
	logger  *logrus.Logger
}

// NewMetrics 创建新的指标收集器
func NewMetrics(enabled bool) *Metrics {
	return &Metrics{
		ErrorCounts: make(map[string]int64),
		enabled:     enabled,
		logger:      logrus.New(),
	}
}

// RecordRequest 记录请求
func (m *Metrics) RecordRequest(success bool, duration time.Duration) {
	if !m.enabled {
		return
	}
	
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.TotalRequests++
	if success {
		m.SuccessfulRequests++
	} else {
		m.FailedRequests++
	}
	
	m.ResponseTimes = append(m.ResponseTimes, duration)
	
	// 保持响应时间数组在合理大小
	if len(m.ResponseTimes) > 1000 {
		m.ResponseTimes = m.ResponseTimes[1:]
	}
}

// RecordError 记录错误
func (m *Metrics) RecordError(errorType string) {
	if !m.enabled {
		return
	}
	
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.ErrorCounts[errorType]++
}

// RecordConnection 记录连接
func (m *Metrics) RecordConnection(active bool) {
	if !m.enabled {
		return
	}
	
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.TotalConnections++
	if active {
		m.ActiveConnections++
	} else {
		m.ActiveConnections--
	}
}

// GetStats 获取统计信息
func (m *Metrics) GetStats() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	stats := map[string]interface{}{
		"total_requests":      m.TotalRequests,
		"successful_requests": m.SuccessfulRequests,
		"failed_requests":     m.FailedRequests,
		"error_counts":        m.ErrorCounts,
		"active_connections":  m.ActiveConnections,
		"total_connections":   m.TotalConnections,
	}
	
	// 计算平均响应时间
	if len(m.ResponseTimes) > 0 {
		var total time.Duration
		for _, duration := range m.ResponseTimes {
			total += duration
		}
		stats["avg_response_time"] = total / time.Duration(len(m.ResponseTimes))
		stats["min_response_time"] = m.minResponseTime()
		stats["max_response_time"] = m.maxResponseTime()
	}
	
	// 计算成功率
	if m.TotalRequests > 0 {
		stats["success_rate"] = float64(m.SuccessfulRequests) / float64(m.TotalRequests) * 100
	}
	
	return stats
}

// minResponseTime 获取最小响应时间
func (m *Metrics) minResponseTime() time.Duration {
	if len(m.ResponseTimes) == 0 {
		return 0
	}
	
	min := m.ResponseTimes[0]
	for _, duration := range m.ResponseTimes {
		if duration < min {
			min = duration
		}
	}
	return min
}

// maxResponseTime 获取最大响应时间
func (m *Metrics) maxResponseTime() time.Duration {
	if len(m.ResponseTimes) == 0 {
		return 0
	}
	
	max := m.ResponseTimes[0]
	for _, duration := range m.ResponseTimes {
		if duration > max {
			max = duration
		}
	}
	return max
}

// StartMetricsServer 启动指标服务器
func (m *Metrics) StartMetricsServer(ctx context.Context, port int) error {
	if !m.enabled {
		return nil
	}
	
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", m.handleMetrics)
	mux.HandleFunc("/health", m.handleHealth)
	
	server := &http.Server{
		Addr:    ":" + fmt.Sprintf("%d", port),
		Handler: mux,
	}
	
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			m.logger.Errorf("metrics server error: %v", err)
		}
	}()
	
	// 等待上下文取消
	<-ctx.Done()
	return server.Shutdown(context.Background())
}

// handleMetrics 处理指标请求
func (m *Metrics) handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// 这里可以返回Prometheus格式的指标
	// 为了简单起见，返回JSON格式
	response := `{
		"total_requests": %d,
		"successful_requests": %d,
		"failed_requests": %d,
		"success_rate": %.2f,
		"avg_response_time_ms": %.2f
	}`
	
	avgTime := float64(0)
	if len(m.ResponseTimes) > 0 {
		var total time.Duration
		for _, duration := range m.ResponseTimes {
			total += duration
		}
		avgTime = float64(total) / float64(len(m.ResponseTimes)) / float64(time.Millisecond)
	}
	
	successRate := float64(0)
	if m.TotalRequests > 0 {
		successRate = float64(m.SuccessfulRequests) / float64(m.TotalRequests) * 100
	}
	
	w.Write([]byte(fmt.Sprintf(response, 
		m.TotalRequests, 
		m.SuccessfulRequests, 
		m.FailedRequests,
		successRate,
		avgTime)))
}

// handleHealth 处理健康检查
func (m *Metrics) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy"}`))
}