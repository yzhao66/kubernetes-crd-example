package controllers

import (
	"encoding/json"
	"github.com/martin-helmich/kubernetes-crd-example/api/types/v1alpha1"
	"github.com/martin-helmich/kubernetes-crd-example/utils"
	"k8s.io/client-go/rest"
	"log"
	"strconv"
	"time"
)

var kubeclientset *rest.RESTClient
var deviceID2 = "traffic-light-instance-01"

type UpdataController struct {
	BaseController
}
type DeviceStatus struct {
	Status v1alpha1.DeviceStatus `json:"status"`
}
type JsonData struct {
	Data []Light
}
type Light struct {
	Color  string `json:"name"`
	Status string `json:"value"`
}

var namespace = "default"
var deviceID = "traffic-light-instance-01"

func init() {


	kubeConfig, err := utils.KubeConfig()
	kubeclientset, err := utils.NewCRDClient(kubeConfig)

	if err != nil {
		log.Fatalf("Failed to create KubeConfig, error : %v", err)
	}
	log.Println(kubeclientset)
}
func (controller *UpdataController) Update() {

	track := controller.Ctx.Input.RequestBody
	Update(track)

}

func Update(params []byte) bool {
	status := build(params)
	deviceStatus := &DeviceStatus{Status: status}
	body, err := json.Marshal(deviceStatus)
	if err != nil {
		log.Printf("Failed to marshal device status %v", deviceStatus)
		return false
	}
	result := kubeclientset.Patch(utils.MergePatchType).Namespace(namespace).Resource(utils.ResourceTypeDevices).Name(deviceID2).Body(body).Do()
	if result.Error() != nil {
		log.Printf("Failed to patch device status %v of device %v in namespace %v \n error:%+v", deviceStatus, deviceID, namespace, result.Error())
		return false
	} else {
		/*log.Printf("Track [ %s ] will be played on speaker %s", track, deviceID)*/
	}
	return true
}
func build(params []byte) v1alpha1.DeviceStatus {
	var redValue string
	var greenValue string
	var yellowValue string
	metadata := map[string]string{"timestamp": strconv.FormatInt(time.Now().Unix()/1e6, 10),
		"type": "string",
	}
	jsonData := JsonData{}
	err := json.Unmarshal(params, &jsonData)
	log.Println(err)

	for i := 0; i < len(jsonData.Data); i++ {

		if jsonData.Data[i].Color == "red" {
			redValue = jsonData.Data[i].Status
		}
		if jsonData.Data[i].Color == "green" {
			greenValue = jsonData.Data[i].Status
		}
		if jsonData.Data[i].Color == "green" {
			yellowValue = jsonData.Data[i].Status
		}
	}
	twins := []v1alpha1.Twin{{PropertyName: "red", Desired: v1alpha1.TwinProperty{Value: redValue, Metadata: metadata}, Reported: v1alpha1.TwinProperty{Value: redValue, Metadata: metadata}}, {PropertyName: "green", Desired: v1alpha1.TwinProperty{Value: greenValue, Metadata: metadata}, Reported: v1alpha1.TwinProperty{Value: greenValue, Metadata: metadata}}, {PropertyName: "yellow", Desired: v1alpha1.TwinProperty{Value: yellowValue, Metadata: metadata}, Reported: v1alpha1.TwinProperty{Value: yellowValue, Metadata: metadata}}}
	devicestatus := v1alpha1.DeviceStatus{Twins: twins}
	return devicestatus
}
