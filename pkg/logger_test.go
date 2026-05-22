package pkg

import "testing"

// recordingLogger 把每次调用记成一条,供测试断言。
type recordingLogger struct {
	records []logRecord
}

type logRecord struct {
	Level  string
	Msg    string
	Fields []any
}

func (r *recordingLogger) Debug(msg string, fields ...any) {
	r.records = append(r.records, logRecord{"DEBUG", msg, fields})
}
func (r *recordingLogger) Info(msg string, fields ...any) {
	r.records = append(r.records, logRecord{"INFO", msg, fields})
}
func (r *recordingLogger) Warn(msg string, fields ...any) {
	r.records = append(r.records, logRecord{"WARN", msg, fields})
}
func (r *recordingLogger) Error(msg string, fields ...any) {
	r.records = append(r.records, logRecord{"ERROR", msg, fields})
}

// TestDefaultLogger_NoopByDefault 验证默认 Logger 是 noop 且不会 panic。
// 这保证了 SDK 可以被未注册 logger 的 caller 安全使用。
func TestDefaultLogger_NoopByDefault(t *testing.T) {
	// 不注册任何 logger,DefaultLogger 应为 noopLogger。
	DefaultLogger.Debug("debug msg")
	DefaultLogger.Info("info msg", "k", "v")
	DefaultLogger.Warn("warn msg")
	DefaultLogger.Error("error msg", "k1", 1, "k2", "v2")
	// 不崩溃即通过
}

// TestDefaultLogger_CanBeReplaced 验证业务侧替换 DefaultLogger 后 SDK
// 内部日志会被正确路由过去。
func TestDefaultLogger_CanBeReplaced(t *testing.T) {
	oldLogger := DefaultLogger
	defer func() { DefaultLogger = oldLogger }()

	rec := &recordingLogger{}
	DefaultLogger = rec

	DefaultLogger.Warn("net error", "method", "GET", "url", "https://example.com", "error", "timeout")

	if len(rec.records) != 1 {
		t.Fatalf("want 1 record, got %d", len(rec.records))
	}
	r := rec.records[0]
	if r.Level != "WARN" {
		t.Fatalf("level = %q, want WARN", r.Level)
	}
	if r.Msg != "net error" {
		t.Fatalf("msg = %q, want \"net error\"", r.Msg)
	}
	if len(r.Fields) != 6 {
		t.Fatalf("fields count = %d, want 6", len(r.Fields))
	}
}
