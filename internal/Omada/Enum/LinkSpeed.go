package Enum

type LinkSpeed int8

const (
	LinkSpeed_Disabled LinkSpeed = -1
	LinkSpeed_Auto     LinkSpeed = 0
	LinkSpeed_10M      LinkSpeed = 1
	LinkSpeed_100M     LinkSpeed = 2
	LinkSpeed_1G       LinkSpeed = 3
	LinkSpeed_2_5G     LinkSpeed = 4
	LinkSpeed_10G      LinkSpeed = 5
	LinkSpeed_5G       LinkSpeed = 6
)
