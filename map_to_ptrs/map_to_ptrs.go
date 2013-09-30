package map_to_ptrs

// Interface for a map[string]*(int64, float64, string, etc)
type MapStringToPtr interface {
	// Clear the contents of the map
	Reset()

	// Length of the contents of the map
	Len() int

	// Keys in the map
	Keys() []string

	// Value of the key
	ValueOf(key string) interface{}

	// Is the value nil or not-present?
	IsNil(key string) bool

	// Is the value not nil?
	IsNotNil(key string) bool
}

// // Is the value nil?
// func IsNil(p *MapStringToPtr, key string) bool {
// 	return nil == p.ValueOf(key)
// }
// 
// // Is the value not nil?
// func IsNotNil(p *MapStringToPtr, key string) bool {
// 	return !IsNil(p, key)
// }
