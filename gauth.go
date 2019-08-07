package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"
	"time"
)

func TimeStamp() (int64, int) {
	time := time.Now().Unix()
	return time / 30, int(time % 30)
}

func normalizeSecret(sec string) string {
	noPadding := strings.ToUpper(strings.Replace(sec, " ", "", -1))
	padLength := 8 - (len(noPadding) % 8)
	if padLength < 8 {
		return noPadding + strings.Repeat("=", padLength)
	}
	return noPadding
}

func AuthCode(sec string, ts int64, encodeType string) (string, error) {
	key, err := base32.StdEncoding.DecodeString(sec)
	if err != nil {
		return "", err
	}
	enc := hmac.New(sha1.New, key)
	msg := make([]byte, 8, 8)
	msg[0] = (byte)(ts >> (7 * 8) & 0xff)
	msg[1] = (byte)(ts >> (6 * 8) & 0xff)
	msg[2] = (byte)(ts >> (5 * 8) & 0xff)
	msg[3] = (byte)(ts >> (4 * 8) & 0xff)
	msg[4] = (byte)(ts >> (3 * 8) & 0xff)
	msg[5] = (byte)(ts >> (2 * 8) & 0xff)
	msg[6] = (byte)(ts >> (1 * 8) & 0xff)
	msg[7] = (byte)(ts >> (0 * 8) & 0xff)
	if _, err := enc.Write(msg); err != nil {
		return "", err
	}
	hash := enc.Sum(nil)
	offset := hash[19] & 0x0f
	trunc := hash[offset : offset+4]
	trunc[0] &= 0x7F
	if strings.TrimSpace(encodeType) == "Steam" {
		steamChars := "23456789BCDFGHJKMNPQRTVWXY"
		steamCharsLen := int32(len(steamChars))
		var res int32
		if err = binary.Read(bytes.NewReader(trunc), binary.BigEndian, &res); err != nil {
			return "", err
		}
		hotp := make([]byte, 5, 5)
		for i := 0; i < 5; i++ {
			idx := res % steamCharsLen
			hotp[i] = steamChars[int(idx)]
			res = res / steamCharsLen
		}
		return string(hotp), nil
	} else {
		res := new(big.Int).Mod(new(big.Int).SetBytes(trunc), big.NewInt(1000000))
		return fmt.Sprintf("%06d", res), nil
	}
}
