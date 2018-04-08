package pfs

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
)

var (
	// ChunkSize is the size of file chunks when resumable upload is used
	ChunkSize = int64(16 * 1024 * 1024) // 16 MB
)

// FullID prints repoName/CommitID
func (c *Commit) FullID() string {
	return fmt.Sprintf("%s/%s", c.Repo.Name, c.ID)
}

// NewHash returns a hash that PFS uses internally to compute checksums.
func NewHash() hash.Hash {
	return sha512.New()
}

// EncodeHash encodes a hash into a readable format.
func EncodeHash(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// GetBlock encodes a hash into a readable format in the form of a Block.
func GetBlock(hash hash.Hash) *Block {
	return &Block{
		Hash: base64.URLEncoding.EncodeToString(hash.Sum(nil)),
	}
}
