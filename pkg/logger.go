package pkg

// Logger 是 SDK 内部使用的日志抽象。
//
// SDK 刻意不直接依赖 slog/logrus/zap 等具体库,避免把调用方绑死在某个 logger
// 生态上。业务侧可以在启动时提供一个实现,把 SDK 内部的细粒度诊断日志
// (网络错误、RDT 透传、cache 失效行为等)统一路由到自己的日志系统。
//
// 职责划分:
//   - OnTokenRefresh / OnAuthRetry 是**事件 hook**,强类型结构体,适合做
//     指标统计 / 告警路由等。
//   - Logger 是**诊断日志**,松散 kv,用来记录事件 hook 未覆盖的边角情况。
//
// 两者正交,业务侧各自按需注册。
//
// 方法签名参考 slog.Logger:fields 是变长 any,每两个一组 (key, value);
// 奇数个或 key 非字符串时,Adapter 实现应自行决定如何处理(通常是丢弃或
// 塞到 "?key" 里)。
type Logger interface {
	Debug(msg string, fields ...any)
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
}

// DefaultLogger 是 SDK 包级日志入口,默认 no-op。业务侧可在进程启动时替换:
//
//	pkg.DefaultLogger = myAdapter{...}
//
// 只应在程序启动阶段设置一次,运行期并发写没有锁保护。
var DefaultLogger Logger = noopLogger{}

type noopLogger struct{}

func (noopLogger) Debug(string, ...any) {}
func (noopLogger) Info(string, ...any)  {}
func (noopLogger) Warn(string, ...any)  {}
func (noopLogger) Error(string, ...any) {}
