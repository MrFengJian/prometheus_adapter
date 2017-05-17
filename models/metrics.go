package models

type MetricValue struct {
    DescribedObject map[string]string `json:"describedObject"`
    MetricName      string `json:"metricName"`
    Value           string `json:"value"`
    Timestamp       string `json:"timestamp"`
    Window          int `json:"window"`
}

type MetricValuesList struct {
    Kind       string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata   map[string]string `json:"metadata"`
    Items      []MetricValue `json:"items"`
}

func NewMetricValuesList() MetricValuesList {
    return MetricValuesList{
        Kind:"MetricValueList",
        ApiVersion:"custom-metrics.metrics.k8s.io/v1alpha1",
        Metadata:make(map[string]string),
        Items:make([]MetricValue, 0),
    }
}

func NewMetricValue() MetricValue {
    return MetricValue{
        DescribedObject:make(map[string]string),
    }
}

func (ml *MetricValuesList) AddMetricValue(value MetricValue) {
    ml.Items = append(ml.Items, value)
}
