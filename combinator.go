package combinator

import "errors"

// Combinator represents our key generator
type Combinator struct {
	charset    []byte
	currentKey []byte
}

// NewCombinator is the constructor for the combinator object
func NewCombinator(charset []byte, start []byte) (*Combinator, error) {
	if len(charset) < 1 {
		return nil, errors.New("charset must have at least one character")
	}
	if len(start) < 1 {
		return nil, errors.New("the starting point must be at least one character")
	}
	return &Combinator{
		charset:    charset,
		currentKey: start,
	}, nil
}

// SetKey sets the key to a predefined value
func (c *Combinator) SetKey(key []byte) {
	c.currentKey = key
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
		} else {
			c.currentKey[pos] = 'A'
		}
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
		} else {
			c.currentKey[pos] = 'A'
		}
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
