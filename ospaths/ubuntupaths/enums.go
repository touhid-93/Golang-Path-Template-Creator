package ubuntupaths

import "gitlab.com/evatix-go/pathhelper/pathwrapper"

//goland:noinspection ALL
const (
	Netplan         pathwrapper.Wrapper = "/etc/netplan"
	NetplanConfig99 pathwrapper.Wrapper = "/etc/netplan/99_config.yaml"
	NetplanConfig00 pathwrapper.Wrapper = "/etc/netplan/00-installer-config.yaml"
)
