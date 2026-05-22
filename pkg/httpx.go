package pkg

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// SharedTransport 是一个全局可复用的 HTTP Transport，配置了连接池、超时和 TLS 等参数
var SharedTransport = &http.Transport{

	// DialContext 定制了底层 TCP 连接的超时和存活设置
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, // 建立 TCP 连接超时
		KeepAlive: 30 * time.Second, // 连接存活检测间隔
	}).DialContext,

	// 最大空闲连接数：全局与每主机
	MaxIdleConns:        100, // 全局最大空闲连接数
	MaxIdleConnsPerHost: 10,  // 单主机最大空闲连接数

	// 最大并发连接数（Go1.11+）
	MaxConnsPerHost: 2000, // 单主机总并发连接数

	// 连接空闲多久后关闭（低于大多数服务器的 idle timeout，避免复用已被对端关闭的连接）
	IdleConnTimeout: 10 * time.Second,

	// TLS 握手超时
	TLSHandshakeTimeout: 30 * time.Second,

	// 等待服务器返回 100-continue 响应的超时时间
	ExpectContinueTimeout: 30 * time.Second,

	// 读取响应头的超时时间
	ResponseHeaderTimeout: 30 * time.Second,

	// 限制响应头最大字节数，防止过大导致OOM
	MaxResponseHeaderBytes: 1 << 20, // 1 MiB

	// 启用连接复用（keep-alive）
	DisableKeepAlives: false,

	// 启用自动解压（gzip/deflate）
	DisableCompression: false,

	TLSClientConfig: &tls.Config{
		// 关闭证书校验
		InsecureSkipVerify: true,
	},

	// 尝试启用 HTTP/2
	ForceAttemptHTTP2: true,
}

// DefaultClient 是基于 SharedTransport 的全局 HTTP 客户端，附带整体超时、Cookie 管理和重定向策略
var DefaultClient = &http.Client{
	Transport: SharedTransport,
	// 整个请求（含 DNS、Dial、TLS、读头、读体）的最大时长
	Timeout: 180 * time.Second,
	Jar:     nil,
}

var LongTimeHttpClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	Timeout: 10 * time.Minute,
	Jar:     nil,
}
