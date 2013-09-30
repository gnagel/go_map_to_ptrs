package map_to_ptrs

import "fmt"
import "strings"

// Interface for a map[string]*(string, float64, string, etc)
type MapStringToStringPtrs map[string]*string

func MakeMapStringToStringPtrs(length int) MapStringToStringPtrs {
	return make(map[string]*string, length)
}

func MakeMapStringToStringPtrsWithKeys(keys ...string) MapStringToStringPtrs {
	p := MakeMapStringToStringPtrs(len(keys))
	for _, key := range keys {
		p.Set(key, nil)
	}
	return p
}

// Length of the map
func (p MapStringToStringPtrs) Len() int {
	return len(p)
}

// Set all the values to nil
func (p MapStringToStringPtrs) Reset() {
	for _, key := range p.Keys() {
		p[key] = nil
	}
}

// Keys of the map
func (p MapStringToStringPtrs) Keys() []string {
	keys := make([]string, p.Len())[0:0]
	for key := range p {
		keys = append(keys, key)
	}
	return keys
}

// Get the value of the key
func (p MapStringToStringPtrs) ValueOf(key string) interface{} {
	return p.Value(key)
}

// Get the value of the key
func (p MapStringToStringPtrs) Value(key string) *string {
	ptr, ok := p[key]
	switch ok {
	case false:
		return nil
	default:
		return ptr
	}
}

// Values of the map
func (p MapStringToStringPtrs) Values() []*string {
	output := make([]*string, p.Len())[0:0]
	for _, value := range p {
		output = append(output, value)
	}
	return output
}

// Value of the key
func (p MapStringToStringPtrs) Set(key string, ptr *string) {
	p[key] = ptr
}

// Is the value nil?
func (p MapStringToStringPtrs) IsNil(key string) bool {
	return nil == p.ValueOf(key)
}

// Is the value not nil?
func (p MapStringToStringPtrs) IsNotNil(key string) bool {
	return !p.IsNil(key)
}

// Format the map as a string
func (p MapStringToStringPtrs) String() string {
	lines := make([]string, p.Len())[0:0]

	for key, ptr := range p {
		switch {
		case nil == ptr:
			lines = append(lines, fmt.Sprintf("%s = NaN", key))
		default:
			lines = append(lines, fmt.Sprintf("%s = %s", key, *ptr))
		}
	}

	return strings.Join(lines, ", ")
}
