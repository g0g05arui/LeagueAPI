package main

type passwordType string

/*
	@Note : stochez doar hashul parolei pe MD5
	@Note : hasul parolei nu intra la marshal , dar poate fi "unmarshel"ed dintr un
			text json
	----TODO----
	de adaugat toate campurile
	----END----

*/
type User struct {
	UID          int64        `json:"UID"`
	Mail         string       `json:"Mail"`
	Username     string       `json:"Username"`
	PasswordHash passwordType `json:"Password"`
}

type HTTPResponse struct {
	Response string `json:"Response"`
	Code     int64  `json:"Code"`
}
