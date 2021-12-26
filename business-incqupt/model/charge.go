package model

import "os/user"

type Charge struct {
	User user.User
	Balance int `json:"balance"`
	Recharge int `json:"recharge"`
	OutPut int `json:"out_put"`
}
