package constants

const (
	// Knative Eventing Namespace
	KnativeEventingNamespace = "knative-eventing"

	// Knative Controller Naming
	KafkaChannelControllerAgentName = "kafka-channel-controller"

	// CRD Kinds
	KnativeChannelKind      = "Channel"
	KafkaChannelKind        = "KafkaChannel"
	KnativeSubscriptionKind = "Subscription"

	// HTTP Port
	HttpPortName = "http"
	// IMPORTANT: HttpServicePort is the inbound port of the service resource. It must be 80 because the
	// Channel resource's url doesn't currently have a port set. Therefore, any client using just the url
	// will send to port 80 by default.
	HttpServicePortNumber = 80
	// IMPORTANT: HttpContainerPortNumber must be 8080 due to dependency issues in the channel. This
	// variable is necessary in order to reconcile the channel resources (service, deployment, etc)
	// correctly.
	// Refer to: https://github.com/kyma-incubator/knative-kafka/blob/master/components/channel/cmd/main.go
	HttpContainerPortNumber = 8080
	
	// Kafka Secret Data Keys
	KafkaSecretDataKeyBrokers  = "brokers"
	KafkaSecretDataKeyUsername = "username"
	KafkaSecretDataKeyPassword = "password"

	// Prometheus MetricsPort
	MetricsPortName = "metrics"

	// Logging Configuration
	LoggingConfigVolumeName = "logging-config"
	LoggingConfigMountPath  = "/etc/knative-kafka"
	LoggingConfigMapName    = "knative-kafka-logging"

	// Health Configuration
	HealthPort                = 8082
	ChannelLivenessDelay      = 10
	ChannelLivenessPeriod     = 5
	ChannelReadinessDelay     = 10
	ChannelReadinessPeriod    = 5
	DispatcherLivenessDelay   = 10
	DispatcherLivenessPeriod  = 5
	DispatcherReadinessDelay  = 10
	DispatcherReadinessPeriod = 5
)
