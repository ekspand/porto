package metricskey

import "github.com/effective-security/metrics"

// Descriptions of emited metrics keys
var (
	HTTPReqPerf = metrics.Describe{
		Name:         "http_requests_perf",
		Type:         metrics.TypeSample,
		RequiredTags: []string{"method", "status", "uri"},
		Help:         "http_requests_perf provides quantiles for HTTP request.",
	}
	HTTPReqByRole = metrics.Describe{
		Name:         "http_requests_role",
		Type:         metrics.TypeCounter,
		RequiredTags: []string{"method", "status", "uri", "role"},
		Help:         "http_requests_role provides counts for HTTP request by role.",
	}

	GRPCReqPerf = metrics.Describe{
		Name:         "rpc_requests_perf",
		Type:         metrics.TypeSample,
		RequiredTags: []string{"method", "status"},
		Help:         "rpc_requests_perf provides quantiles for gRPC request.",
	}
	GRPCReqByRole = metrics.Describe{
		Name:         "rpc_requests_role",
		Type:         metrics.TypeCounter,
		RequiredTags: []string{"method", "status", "role"},
		Help:         "rpc_requests_role provides counts for gRPC request by role.",
	}
)

// Metrics returns slice of metrics from this repo
var Metrics = []*metrics.Describe{
	&HTTPReqPerf,
	&HTTPReqByRole,
	&GRPCReqPerf,
	&GRPCReqPerf,
	&GRPCReqByRole,
}
