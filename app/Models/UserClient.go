package Models

import "log"

type UserClient struct {
	Uid            int    `json:"uid"`
	Nickname       string `json:"nickname"`
	IntegralAmount int    `json:"integral_amount"`
	OpenidMp       string `json:"openid_mp"`
}

func SelectAll() []UserClient {
	var clients []UserClient
	result := DbConn.Table("user_client").Find(&clients)
	log.Println(result.RowsAffected, `RowsAffected`)
	if result.RowsAffected > 1 {
		return clients
	}
	return clients
}
