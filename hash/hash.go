package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash/fnv"
)

// MD5 returns the MD5 checksum of the string.
func MD5(text string) string {
	sum := md5.Sum([]byte(text))
	return hex.EncodeToString(sum[:])
}

// SHA1 returns the SHA1 checksum of the string.
func SHA1(text string) string {
	sum := sha1.Sum([]byte(text))
	return hex.EncodeToString(sum[:])
}

// SHA1Short returns the first 6 characters of the SHA1 checksum of the string.
func SHA1Short(text string) string {
	sum := sha1.Sum([]byte(text))
	return hex.EncodeToString(sum[:3])
}

// SHA256 returns the SHA256 checksum of the string.
func SHA256(text string) string {
	sum := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sum[:])
}

// FNV32 returns a new 32-bit FNV-1a encoded string.
func FNV32(text string) string {
	h := fnv.New32a()
	h.Write([]byte(text)) // nolint
	return hex.EncodeToString(h.Sum(nil))
}

// FNV64 returns a new 64-bit FNV-1a encoded string.
func FNV64(text string) string {
	h := fnv.New64a()
	h.Write([]byte(text)) // nolint
	return hex.EncodeToString(h.Sum(nil))
}

// FNV128 returns a new 128-bit FNV-1a encoded string.
func FNV128(text string) string {
	h := fnv.New128a()
	h.Write([]byte(text)) // nolint
	return hex.EncodeToString(h.Sum(nil))
}
