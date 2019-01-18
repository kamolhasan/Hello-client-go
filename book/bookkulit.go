package book

import (
	kappv1 "github.com/appscode/kutil/apps/v1"
	kcorev1 "github.com/appscode/kutil/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func CreateDeploymentKutil() {

	clientset := CreateClientSet()
	var replica int32 = 3
	deploy := &appsv1.Deployment{

		ObjectMeta: metav1.ObjectMeta{
			Name: "books",
			Labels: map[string]string{
				"app": "books",
			},
			Namespace: "default",
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

	_, _, err := kappv1.CreateOrPatchDeployment(
		clientset,
		deploy.ObjectMeta,
		func(d *appsv1.Deployment) *appsv1.Deployment {
			d = deploy
			return d
		},
	)

	if err != nil {
		panic(err)
	}

}

func CreateServiceKutil() {
	clientset := CreateClientSet()

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bookservice",
			Namespace:"default",
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

	_, _, err := kcorev1.CreateOrPatchService(
		clientset,
		service.ObjectMeta,

		func(s *corev1.Service) *corev1.Service {
			s = service
			return s
		},
	)

	if err != nil {
		panic(err)
	}
}
