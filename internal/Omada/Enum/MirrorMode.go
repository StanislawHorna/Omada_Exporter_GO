package Enum

type MirrorMode int8

const (
	MirrorMode_Ingress       MirrorMode = 0
	MirrorMode_Egress        MirrorMode = 1
	MirrorMode_IngressEgress MirrorMode = 2
)
