package data_structures_algorithms

import "strings"

type Codec struct{}

// Encodes a list of strings to a single string.
func (c *Codec) Encode(strs []string) string {
	var result strings.Builder
	for _, str := range strs {
		result.WriteByte(byte(len(str)))
		result.WriteString(str)
	}
	return result.String()
}

// Decodes a single string to a list of strings.
func (c *Codec) Decode(strs string) []string {
	// var count
	return nil
}

// The Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs))
