package ant

import (
	"crypto/md5"
	"encoding/binary"
	"log"
    "strings"
    "ms/sun/helper"
)

func noErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Hash(str string) int {
	sh1 := md5.Sum([]byte(str))
	b := sh1[0]
	//fmt.Println(b)
	b = b >> 1
	//fmt.Println(b)
	bytes := []byte{b, sh1[1], sh1[2], sh1[3]}

	res := binary.BigEndian.Uint32(bytes)
	//fmt.Println(res, int32(res))
	return int(res)
}

func Hash32(string string) int32 {
    return int32(Hash(string)%10000)
}

func MyHash(string string) int {
	h := 15485862
	a := 7

	for i := 0; i < len(string); i++ {
		h = ((h * a) + int(int8(string[0]))) / 3
	}
	return h
}

func MyHash2(string string) int {
    s:=strings.Split(string,".")
    rs:= helper.IntToStr((Hash(s[0]) %1000)+1) + "00" + helper.IntToStr(Hash(s[1])%100000)
    //fmt.Println("+++++++"+rs)
    return helper.StrToInt(rs,-1)
}
