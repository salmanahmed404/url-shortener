package main

const (
	urlBase      string = "gorl.ly/"
	characterSet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func base62Encode(id uint) (enc string) {
	enc = urlBase
	for id > 0 {
		enc += string(characterSet[id%62])
		id = id / 62
	}
	return
}
