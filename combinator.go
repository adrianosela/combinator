package combinator

import (
	"errors"
	"math/rand"
	"time"
)

// Combinator represents our key generator
type Combinator struct {
	charset    []byte
	currentKey []byte
}

// NewCombinator is the constructor for the combinator object
func NewCombinator(charset []byte) (*Combinator, error) {
	if len(charset) < 1 {
		return nil, errors.New("charset must have at least one character")
	}
	rand.Seed(time.Now().UnixNano())
	return &Combinator{
		charset:    charset,
		currentKey: []byte{charset[0]},
	}, nil
}

// GetKey returns the current key in the combinator
func (c *Combinator) GetKey() []byte {
	return c.currentKey
}

// SetKey sets the key to a predefined value
func (c *Combinator) SetKey(key []byte) []byte {
	c.currentKey = key
	return c.currentKey
}

// RandKey sets a random key of length 'size' and returns it
func (c *Combinator) RandKey(size int) []byte {
	c.currentKey = []byte{}
	for i := 0; i < size; i++ {
		c.currentKey = append(c.currentKey, c.charset[rand.Intn(len(c.charset))])
	}
	return c.currentKey
}

// RandPerm is a sad and inefficient way of getting a random permutation
func (c *Combinator) RandPerm(size int) []byte {
	perm := []byte{}

	used := make(map[byte]bool)

	elems := 0
	for elems < size {
		c := c.charset[rand.Intn(len(c.charset))]
		if _, ok := used[c]; !ok {
			perm = append(perm, c)
			used[c] = true
		}
		elems++
	}

	return perm
}

// NextRight advances the key by incrementing from the right of the current key
func (c *Combinator) NextRight() []byte {
	for pos := 0; pos < len(c.currentKey); pos++ {
		if c.currentKey[pos] != '9' {
			c.currentKey[pos] = c.charset[c.indexOf(c.currentKey[pos])+1]
			return c.currentKey
		}
		if pos == len(c.currentKey)-1 {
			c.currentKey = append([]byte{'A'}, c.currentKey...)
		}
		c.currentKey[pos] = 'A'
	}
	return c.currentKey
}

// NextLeft advances the key by incrementing from the right of the current key
func (c *Combinator) NextLeft() []byte {
	for pos := len(c.currentKey) - 1; pos >= 0; pos-- {
		if c.currentKey[pos] != '9' {
			c.currentKey[pos] = c.charset[c.indexOf(c.currentKey[pos])+1]
			return c.currentKey
		}
		if pos == 0 {
			c.currentKey = append(c.currentKey, 'A')
		}
		c.currentKey[pos] = 'A'
	}
	return c.currentKey
}

func (c *Combinator) indexOf(b byte) int {
	for i, c := range c.charset {
		if c == b {
			return i
		}
	}
	return -1
}
