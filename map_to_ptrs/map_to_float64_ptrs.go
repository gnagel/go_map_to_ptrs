package map_to_ptrs

import "fmt"
import "strings"

// Interface for a map[string]*(float64, float64, string, etc)
type MapStringToFloat64Ptrs map[string]*float64

func MakeMapStringToFloat64Ptrs(length int) MapStringToFloat64Ptrs {
	return make(map[string]*float64, length)
}

// Length of the map
func (p MapStringToFloat64Ptrs) Len() int {
	return len(p)
}

// Keys of the map
func (p MapStringToFloat64Ptrs) Keys() []string {
	keys := make([]string, p.Len())[0:0]
	for key := range p {
		keys = append(keys, key)
	}
	return keys
}

// Get the value of the key
func (p MapStringToFloat64Ptrs) ValueOf(key string) interface{} {
	return p.Value(key)
}

// Get the value of the key
func (p MapStringToFloat64Ptrs) Value(key string) *float64 {
	ptr, ok := p[key]
	switch ok {
	case false:
		return nil
	default:
		return ptr
	}
}

// Values of the map
func (p MapStringToFloat64Ptrs) Values() []*float64 {
	output := make([]*float64, p.Len())[0:0]
	for _, value := range p {
		output = append(output, value)
	}
	return output
}

// Value of the key
func (p MapStringToFloat64Ptrs) Set(key string, ptr *float64) {
	p[key] = ptr
}

// Is the value nil?
func (p MapStringToFloat64Ptrs) IsNil(key string) bool {
	return nil == p.ValueOf(key)
}

// Is the value not nil?
func (p MapStringToFloat64Ptrs) IsNotNil(key string) bool {
	return !p.IsNil(key)
}

// Format the map as a string
func (p MapStringToFloat64Ptrs) String() string {
	lines := make([]string, p.Len())[0:0]

	for key, ptr := range p {
		switch {
		case nil == ptr:
			lines = append(lines, fmt.Sprintf("%s = NaN", key))
		default:
			lines = append(lines, fmt.Sprintf("%s = %0.6f", key, *ptr))
		}
	}

	return strings.Join(lines, ", ")
}
