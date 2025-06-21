package Enum

type MirrorMode int8

const (
	MirrorMode_Ingress       MirrorMode = 0
	MirrorMode_Egress        MirrorMode = 1
	MirrorMode_IngressEgress MirrorMode = 2
)

func (mm MirrorMode) String() string {
	switch mm {
	case MirrorMode_Ingress:
		return "Ingress"
	case MirrorMode_Egress:
		return "Egress"
	case MirrorMode_IngressEgress:
		return "IngressEgress"
	default:
		return "invalid"
	}
}
