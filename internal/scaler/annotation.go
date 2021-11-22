package scaler

import v1 "k8s.io/api/core/v1"

const NamespaceAnnotation = "saas.stackstate.com/beacher"

type ScaleAnnotationValue string

var (
	ScaleSkip ScaleAnnotationValue = "skip"
)

func IsAnnotatedWithSkip(ns v1.Namespace) bool {
	return ns.Annotations[NamespaceAnnotation] == string(ScaleSkip)
}
