package controllers

import (
    "github.com/astaxie/beego"
    "prometheus_adapter/models"
    "time"
)

type MetricsController struct {
    beego.Controller
}
//serve api /apis/custom-metrics.metrics.k8s.io/v1alpha1/namespaces/{namespace}/services/{service}/{metrics_name}

func (m *MetricsController) GetServiceMetric(){
    namespace:=m.Ctx.Input.Param(":namespace")
    serviceName:=m.Ctx.Input.Param(":service")
    metricName:=m.Ctx.Input.Param(":metric_name")

    metric:=models.NewMetricValue()
    metric.DescribedObject["kind"]="Service"
    metric.DescribedObject["namespace"]=namespace
    metric.DescribedObject["name"]=serviceName
    metric.DescribedObject["apiVersion"]="/_internal"
    metric.MetricName=metricName
    metric.Value="1000"
    metric.Window=60
    metric.Timestamp=GetTimestamp()
    metricsList:=models.NewMetricValuesList()
    metricsList.AddMetricValue(metric)
    m.Data["json"]=metricsList

    m.ServeJSON()
}


func GetTimestamp()string{
    now:=time.Now().Unix()
    tm:=time.Unix(now,0)
    return tm.Format("2006-01-02T03:04:05Z")
}