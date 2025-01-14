package config

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"

	"github.com/wundergraph/cosmo/router/internal/logging"
	"github.com/wundergraph/cosmo/router/internal/otel/otelconfig"
)

const defaultConfigPath = "config.yaml"

type Base64Decoder []byte

func (ipd *Base64Decoder) Decode(value string) error {
	decoded, err := b64.StdEncoding.DecodeString(value)
	if err != nil {
		return fmt.Errorf("could not decode base64 string: %w", err)
	}

	*ipd = decoded

	return nil
}

type BytesString uint64

func (b *BytesString) Decode(value string) error {
	decoded, err := humanize.ParseBytes(value)
	if err != nil {
		return fmt.Errorf("could not parse bytes string: %w", err)
	}

	*b = BytesString(decoded)

	return nil
}

func (b *BytesString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	return b.Decode(s)
}

type Graph struct {
	Name  string `yaml:"name" envconfig:"FEDERATED_GRAPH_NAME" validate:"required"`
	Token string `yaml:"token" envconfig:"GRAPH_API_TOKEN" validate:"required"`
}

type TracingExporterConfig struct {
	BatchTimeout  time.Duration `yaml:"batch_timeout" default:"10s" validate:"required,min=5s,max=120s"`
	ExportTimeout time.Duration `yaml:"export_timeout" default:"30s" validate:"required,min=5s,max=120s"`
}

type TracingExporter struct {
	Exporter              otelconfig.Exporter `yaml:"exporter" validate:"oneof=http grpc"`
	Endpoint              string              `yaml:"endpoint" validate:"http_url"`
	HTTPPath              string              `yaml:"path"`
	Headers               map[string]string   `yaml:"headers"`
	TracingExporterConfig `yaml:",inline"`
}

type Tracing struct {
	Enabled      bool              `yaml:"enabled" default:"true" envconfig:"TRACING_ENABLED"`
	SamplingRate float64           `yaml:"sampling_rate" default:"1" validate:"required,min=0,max=1" envconfig:"TRACING_SAMPLING_RATE"`
	Exporters    []TracingExporter `yaml:"exporters"`
}

type Prometheus struct {
	Enabled    bool   `yaml:"enabled" default:"true" envconfig:"PROMETHEUS_ENABLED"`
	Path       string `yaml:"path" default:"/metrics" validate:"uri" envconfig:"PROMETHEUS_HTTP_PATH"`
	ListenAddr string `yaml:"listen_addr" default:"127.0.0.1:8088" validate:"hostname_port" envconfig:"PROMETHEUS_LISTEN_ADDR"`
}

type MetricsOTLPExporter struct {
	Exporter otelconfig.Exporter `yaml:"exporter" validate:"oneof=http grpc"`
	Endpoint string              `yaml:"endpoint" validate:"http_url"`
	HTTPPath string              `yaml:"path"`
	Headers  map[string]string   `yaml:"headers"`
}

type Metrics struct {
	OTLP       MetricsOTLP `yaml:"otlp"`
	Prometheus Prometheus  `yaml:"prometheus"`
}

type MetricsOTLP struct {
	Enabled   bool                  `yaml:"enabled" default:"true" envconfig:"METRICS_OTLP_ENABLED"`
	Exporters []MetricsOTLPExporter `yaml:"exporters"`
}

type Telemetry struct {
	ServiceName string  `yaml:"service_name" default:"cosmo-router" envconfig:"TELEMETRY_SERVICE_NAME" validate:"required"`
	Tracing     Tracing `yaml:"tracing"`
	Metrics     Metrics `yaml:"metrics"`
}

type CORS struct {
	AllowOrigins     []string      `yaml:"allow_origins" default:"*" envconfig:"CORS_ALLOW_ORIGINS"`
	AllowMethods     []string      `yaml:"allow_methods" default:"HEAD,GET,POST" envconfig:"CORS_ALLOW_METHODS"`
	AllowHeaders     []string      `yaml:"allow_headers" default:"Origin,Content-Length,Content-Type" envconfig:"CORS_ALLOW_HEADERS"`
	AllowCredentials bool          `yaml:"allow_credentials" default:"true" envconfig:"CORS_ALLOW_CREDENTIALS"`
	MaxAge           time.Duration `yaml:"max_age" default:"5m" validate:"required,min=5m" envconfig:"CORS_MAX_AGE"`
}

type TrafficShapingRules struct {
	// All is a set of rules that apply to all requests
	All GlobalSubgraphRequestRule `yaml:"all"`
	// Apply to requests from clients to the router
	Router RouterTrafficConfiguration `yaml:"router"`
}

type RouterTrafficConfiguration struct {
	// MaxRequestBodyBytes is the maximum size of the request body in bytes
	MaxRequestBodyBytes BytesString `yaml:"max_request_body_size" default:"5MB" validate:"min=1000000"`
}

type GlobalSubgraphRequestRule struct {
	BackoffJitterRetry BackoffJitterRetry `yaml:"retry"`
	// See https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	RequestTimeout         time.Duration `yaml:"request_timeout" default:"60s" validate:"required,min=1s"`
	DialTimeout            time.Duration `yaml:"dial_timeout" default:"30s"`
	ResponseHeaderTimeout  time.Duration `yaml:"response_header_timeout" default:"0s"`
	ExpectContinueTimeout  time.Duration `yaml:"expect_continue_timeout" default:"0s"`
	TLSHandshakeTimeout    time.Duration `yaml:"tls_handshake_timeout" default:"10s"`
	KeepAliveIdleTimeout   time.Duration `yaml:"keep_alive_idle_timeout" default:"0s"`
	KeepAliveProbeInterval time.Duration `yaml:"keep_alive_probe_interval" default:"30s"`
}

type BackoffJitterRetry struct {
	Enabled     bool          `yaml:"enabled" default:"true" envconfig:"RETRY_ENABLED"`
	Algorithm   string        `yaml:"algorithm" default:"backoff_jitter" validate:"oneof=backoff_jitter"`
	MaxAttempts int           `yaml:"max_attempts" default:"5" validate:"required,min=1,required_if=Algorithm backoff_jitter"`
	MaxDuration time.Duration `yaml:"max_duration" default:"10s" validate:"required,min=1s,required_if=Algorithm backoff_jitter"`
	Interval    time.Duration `yaml:"interval" default:"3s" validate:"required,min=100ms,required_if=Algorithm backoff_jitter"`
}

type HeaderRules struct {
	// All is a set of rules that apply to all requests
	All       GlobalHeaderRule            `yaml:"all"`
	Subgraphs map[string]GlobalHeaderRule `yaml:"subgraphs" validate:"dive"`
}

type GlobalHeaderRule struct {
	// Request is a set of rules that apply to requests
	Request []RequestHeaderRule `yaml:"request" validate:"dive"`
}

type RequestHeaderRule struct {
	// Operation describes the header operation to perform e.g. "propagate"
	Operation string `yaml:"op" validate:"oneof=propagate"`
	// Matching is the regex to match the header name against
	Matching string `yaml:"matching" validate:"excluded_with=Named"`
	// Named is the exact header name to match
	Named string `yaml:"named" validate:"excluded_with=Matching"`
	// Default is the default value to set if the header is not present
	Default string `yaml:"default"`
}

type EngineDebugConfiguration struct {
	PrintOperationTransformations bool `envconfig:"ENGINE_DEBUG_PRINT_OPERATION_TRANSFORMATIONS"`
	PrintOperationEnableASTRefs   bool `envconfig:"ENGINE_DEBUG_PRINT_OPERATION_ENABLE_AST_REFS"`
	PrintPlanningPaths            bool `envconfig:"ENGINE_DEBUG_PRINT_PLANNING_PATHS"`
	PrintQueryPlans               bool `envconfig:"ENGINE_DEBUG_PRINT_QUERY_PLANS"`
	PrintNodeSuggestions          bool `envconfig:"ENGINE_DEBUG_PRINT_NODE_SUGGESTIONS"`
	ConfigurationVisitor          bool `envconfig:"ENGINE_DEBUG_CONFIGURATION_VISITOR"`
	PlanningVisitor               bool `envconfig:"ENGINE_DEBUG_PLANNING_VISITOR"`
	DatasourceVisitor             bool `envconfig:"ENGINE_DEBUG_DATASOURCE_VISITOR"`
}

type EngineExecutionConfiguration struct {
	Debug              EngineDebugConfiguration
	EnableSingleFlight bool `default:"true" envconfig:"ENGINE_ENABLE_SINGLE_FLIGHT"`
}

type OverrideRoutingURLConfiguration struct {
	Subgraphs map[string]string `yaml:"subgraphs" validate:"dive,required,url"`
}

type Config struct {
	Version string `yaml:"version"`

	Graph     Graph     `yaml:"graph"`
	Telemetry Telemetry `yaml:"telemetry"`
	CORS      CORS      `yaml:"cors"`

	Modules        map[string]interface{} `yaml:"modules"`
	Headers        HeaderRules            `yaml:"headers"`
	TrafficShaping TrafficShapingRules    `yaml:"traffic_shaping"`

	ListenAddr           string        `yaml:"listen_addr" default:"localhost:3002" validate:"hostname_port" envconfig:"LISTEN_ADDR"`
	ControlplaneURL      string        `yaml:"controlplane_url" validate:"required" default:"https://cosmo-cp.wundergraph.com" envconfig:"CONTROLPLANE_URL" validate:"uri"`
	PlaygroundEnabled    bool          `yaml:"playground_enabled" default:"true" envconfig:"PLAYGROUND_ENABLED"`
	IntrospectionEnabled bool          `yaml:"introspection_enabled" default:"true" envconfig:"INTROSPECTION_ENABLED"`
	LogLevel             string        `yaml:"log_level" default:"info" envconfig:"LOG_LEVEL" validate:"oneof=debug info warning error fatal panic"`
	JSONLog              bool          `yaml:"json_log" default:"true" envconfig:"JSON_LOG"`
	ShutdownDelay        time.Duration `yaml:"shutdown_delay" default:"30s" validate:"required,min=5s" envconfig:"SHUTDOWN_DELAY"`
	GracePeriod          time.Duration `yaml:"grace_period" default:"20s" validate:"required" envconfig:"GRACE_PERIOD"`
	PollInterval         time.Duration `yaml:"poll_interval" default:"10s" validate:"required,min=5s" envconfig:"POLL_INTERVAL"`
	HealthCheckPath      string        `yaml:"health_check_path" default:"/health" envconfig:"HEALTH_CHECK_PATH" validate:"uri"`
	ReadinessCheckPath   string        `yaml:"readiness_check_path" default:"/health/ready" envconfig:"READINESS_CHECK_PATH" validate:"uri"`
	LivenessCheckPath    string        `yaml:"liveness_check_path" default:"/health/live" envconfig:"LIVENESS_CHECK_PATH" validate:"uri"`
	GraphQLPath          string        `yaml:"graphql_path" default:"/graphql" envconfig:"GRAPHQL_PATH"`

	ConfigPath       string `envconfig:"CONFIG_PATH" validate:"omitempty,filepath"`
	RouterConfigPath string `yaml:"router_config_path" envconfig:"ROUTER_CONFIG_PATH" validate:"omitempty,filepath"`

	OverrideRoutingURL OverrideRoutingURLConfiguration `yaml:"override_routing_url"`

	EngineExecutionConfiguration EngineExecutionConfiguration
}

func LoadConfig(envOverride string) (*Config, error) {
	godotenv.Load(".env.local")
	godotenv.Load()

	if envOverride != "" {
		godotenv.Overload(envOverride)
	}

	var c Config

	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	configPathOverride := false

	if c.ConfigPath != "" {
		configPathOverride = true
	} else {
		// Ensure default
		c.ConfigPath = defaultConfigPath
	}

	// Configuration from environment variables. We don't have the config here.
	logLevel, err := logging.ZapLogLevelFromString(c.LogLevel)
	logger := logging.New(!c.JSONLog, c.LogLevel == "debug", logLevel).
		With(zap.String("component", "@wundergraph/router"))

	// Custom config path can only be supported through environment variable
	configBytes, err := os.ReadFile(c.ConfigPath)

	if err != nil {
		if configPathOverride {
			return nil, fmt.Errorf("could not read custom config file %s: %w", c.ConfigPath, err)
		} else {
			logger.Info("Default config file is not loaded",
				zap.String("configPath", defaultConfigPath),
				zap.Error(err),
			)
		}
	}
	expandedConfigBytes := []byte(os.ExpandEnv(string(configBytes)))

	if err == nil {
		if err := yaml.Unmarshal(expandedConfigBytes, &c); err != nil {
			return nil, fmt.Errorf("failed to unmarshal router config: %w", err)
		}
	}

	err = validator.New().Struct(c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
