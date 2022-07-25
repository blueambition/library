package hashid

import (
	"github.com/speps/go-hashids"
)

//Id 混淆
func GetHash(id int, salt string, len int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = len
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{id})
	return e
}

//还原混淆
func GetHashId(hashId, salt string, len int) int {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = len
	h, _ := hashids.NewWithData(hd)
	e, _ := h.DecodeWithError(hashId)
	return e[0]
}
