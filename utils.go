package main

const (
	urlBase      string = "http://127.0.0.1:8080/gorl/"
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

func base62Decode(shortURL string) (id uint) {
	id = 0
	for _, ch := range shortURL {
		if 'a' <= ch && ch <= 'z' {
			id = id*62 + uint(ch-'a')
		} else if 'A' <= ch && ch <= 'Z' {
			id = id*62 + uint(ch-'A') + 26
		} else {
			id = id*62 + uint(ch-'0') + 52
		}
	}
	return
}
