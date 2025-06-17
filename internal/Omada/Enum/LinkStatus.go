package Enum

type LinkStatus int8

const (
	LinkStatus_Unknown LinkStatus = -1
	LinkStatus_Down    LinkStatus = 0
	LinkStatus_Up      LinkStatus = 1
)
