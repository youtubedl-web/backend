package file

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// GenerateHash generates an hash and creates a folder for the file to stored in
func GenerateHash(base string) (string, int) {
	// get current time as a string
	nowStr := strconv.Itoa(int(time.Now().Unix()))

	// hash it
	hasher := sha256.New()
	hasher.Write([]byte(nowStr))

	hash := hex.EncodeToString(hasher.Sum(nil))

	// create a folder named after the hash
	err := os.MkdirAll(filepath.Join(base, hash), 0755)
	if err != nil {
		return "", 0
	}

	return hash, 1
}
