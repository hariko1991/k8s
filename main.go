package main

import (
	"context"
	"strconv"

	"github.com/baetyl/baetyl-go/v2/log"
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientCmd"
)

func main() {
	viper.SetConfigFile("./conf/conf.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.L().Error("", log.Code(err), log.Error(err))
		return
	}

	kube_path := viper.GetString("kube")
	log.L().Info("", log.Any("kube_path", kube_path))
	k8s_config, err := clientcmd.BuildConfigFromFlags("", kube_path)

	if err != nil {
		log.L().Error("", log.Error(err))
		return
	}

	k8s_client, err2 := kubernetes.NewForConfig(k8s_config)
	if err2 != nil {
		log.L().Error("", log.Error(err))
		return
	}

	nl, err3 := k8s_client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err3 != nil {
		log.L().Error("", log.Error(err3))
		return
	}

	log.L().Info(strconv.Itoa(len(nl.Items)))

	nl2, err4 := k8s_client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err4 != nil {
		log.L().Error("", log.Error(err4))
	}

	for _, val := range nl2.Items {
		log.L().Info("", log.Any("ClusterName", val.ClusterName), log.Any("Name", val.Name))
	}

}
