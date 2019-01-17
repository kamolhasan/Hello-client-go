package book

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func CreateClientSet() *kubernetes.Clientset {

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {

		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	return clientset
}

func CreateDeployment() {

	clientset := CreateClientSet()

	var replica int32 = 3
	deploy := &appsv1.Deployment{

		ObjectMeta: metav1.ObjectMeta{
			Name: "books",
			Labels: map[string]string{
				"app": "books",
			},
		},
		Spec: appsv1.DeploymentSpec{

			Replicas: &replica,

			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "books",
				},
			},

			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "books",
					Labels: map[string]string{
						"app": "books",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "books",
							Image:           "kamolhasan/bookapi",
							ImagePullPolicy: "IfNotPresent",
						},
					},
					RestartPolicy: "Always",
				},
			},
		},
	}
	_, err := clientset.AppsV1().Deployments("default").Create(deploy)

	if err != nil {
		panic(err)
	}
}

func DeleteDeployment() {
	clientset := CreateClientSet()

	err := clientset.AppsV1().Deployments("default").Delete("books", nil)
	if err != nil {
		panic(err)
	}
}

func UpdateDeployment() {
	clientset := CreateClientSet()
	demployment, err := clientset.AppsV1().Deployments("default").Get("books",metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	log.Println("Scaling Down from 3 to 1")

	var num int32 = 1
	demployment.Spec.Replicas = &num
	_, err = clientset.AppsV1().Deployments("default").Update(demployment)
	if err != nil {
		panic(err)
	}
}

func CreateService() {
	clientset := CreateClientSet()

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bookservice",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port:       8080,
					TargetPort: intstr.IntOrString{IntVal: 8000},
				},
			},
			Selector: map[string]string{
				"app": "books",
			},
			Type: "NodePort",
		},
	}

	_, err := clientset.CoreV1().Services("default").Create(service)

	if err != nil {
		panic(err)
	}
}

func DeleteService() {
	clientset := CreateClientSet()

	err := clientset.CoreV1().Services("default").Delete("bookservice", nil)

	if err != nil {
		panic(err)

	}
}

func CreateIngress() {
	clientset := CreateClientSet()

	ingress := &extv1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bookingress",
		},
		Spec: extv1beta1.IngressSpec{
			Rules: []extv1beta1.IngressRule{
				{
					Host: "booklist.com",
					IngressRuleValue: extv1beta1.IngressRuleValue{
						HTTP: &extv1beta1.HTTPIngressRuleValue{
							Paths: []extv1beta1.HTTPIngressPath{
								{
									Path: "/",
									Backend: extv1beta1.IngressBackend{
										ServiceName: "bookservice",
										ServicePort: intstr.FromInt(8080),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := clientset.ExtensionsV1beta1().Ingresses("default").Create(ingress)
	if err != nil {
		panic(err)
	}
}

func DeleteIngress() {
	clientset := CreateClientSet()

	err := clientset.ExtensionsV1beta1().Ingresses("default").Delete("bookingress", nil)
	if err != nil {
		panic(err)
	}

}

