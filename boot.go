package main

import (
	"fmt"
	"log"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	apierrors "k8s.io/client-go/1.4/pkg/api/errors"
	"k8s.io/client-go/1.4/pkg/fields"
	"k8s.io/client-go/1.4/pkg/labels"
	"k8s.io/client-go/1.4/rest"
)

func deletePods(kubeClient *kubernetes.Clientset, daemonSetName string, succChan chan<- string, errChan chan<- error) {
	daemonSet, err := kubeClient.DaemonSets("deis").Get(daemonSetName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			succChan <- fmt.Sprintf("%s daemonset not found", daemonSetName)
			return
		}
		errChan <- err
		return
	}
	daemonImage := daemonSet.Spec.Template.Spec.Containers[0].Image
	labelMap := labels.Set{"app": daemonSetName}
	loggerPods, err := kubeClient.Pods("deis").List(api.ListOptions{LabelSelector: labelMap.AsSelector(), FieldSelector: fields.Everything()})
	if err != nil {
		errChan <- fmt.Errorf("Error getting pods for %s: %s", daemonSetName, err.Error())
		return
	}
	for _, pod := range loggerPods.Items {
		podImage := pod.Spec.Containers[0].Image
		if podImage != daemonImage {
			err := kubeClient.Pods("deis").Delete(pod.GetName(), &api.DeleteOptions{})
			if err != nil {
				errChan <- fmt.Errorf("error deleting %s pod: %s", pod.GetName(), err.Error())
				return
			}
		}
	}
	succChan <- fmt.Sprintf("pods for %s daemonset deleted successfuly", daemonSetName)
}

func main() {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	succChan, errChan := make(chan string), make(chan error)
	daemonSets := []string{"deis-logger-fluentd", "deis-monitor-telegraf", "deis-registry-proxy"}
	for _, daemonset := range daemonSets {
		go deletePods(kubeClient, daemonset, succChan, errChan)
	}
	for i := 0; i < len(daemonSets); i++ {
		select {
		case successMsg := <-succChan:
			fmt.Println(successMsg)
		case err = <-errChan:
			fmt.Println(err)
		}
	}
	if err != nil {
		log.Fatal("error in deleting pods", err)
	}
}
