package map_to_ptrs

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestMapStringToInt64PtrsSpecs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in benchmark mode.")
		return
	}
	r := gospec.NewRunner()
	r.AddSpec(MapStringToInt64PtrsSpecs)
	gospec.MainGoTest(r, t)
}

func MapStringToInt64PtrsSpecs(c gospec.Context) {

	c.Specify("[MapStringToInt64Ptrs][Make] returns a reference", func() {
		value := MakeMapStringToInt64Ptrs(2)
		value.Set("Bob", nil)
		value.Set("George", nil)
		c.Expect(value.Len(), gospec.Equals, 2)

		value2 := func(m MapStringToInt64Ptrs) MapStringToInt64Ptrs {
			i := int64(123)
			m["Gary"] = &i
			return m
		}(value)
		
		c.Expect(value2.Len(), gospec.Equals, 3)
		c.Expect(value.Len(), gospec.Equals, 3)
		c.Expect(value2.String(), gospec.Equals, value.String())
	})

	c.Specify("[PtrMapInt64][Len]", func() {
		value := MakeMapStringToInt64Ptrs(0)
		c.Expect(value.Len(), gospec.Equals, 0)

		value.Set("Bob", nil)
		c.Expect(value.Len(), gospec.Equals, 1)

		value.Set("George", nil)
		c.Expect(value.Len(), gospec.Equals, 2)
		
	})
	
	c.Specify("[PtrMapInt64][Value]", func() {
		count := int64(123)
		value := MakeMapStringToInt64Ptrs(2)
		value.Set("Bob", &count)
		value.Set("George", nil)
	
		c.Expect(value.Value("Bob"), gospec.Equals, &count)
		c.Expect(*value.Value("Bob"), gospec.Equals, int64(123))
		c.Expect(value.Value("George"), gospec.Satisfies, nil == value.Value("George"))
		c.Expect(value.Value("Cache Miss"), gospec.Satisfies, nil == value.Value("Cache Miss"))
	})
	
	c.Specify("[PtrMapInt64][String]", func() {
		count := int64(123)
		value := MakeMapStringToInt64Ptrs(2)
		value.Set("Bob", &count)
		value.Set("George", nil)
		c.Expect(value.String(), gospec.Equals, "Bob = 123, George = NaN")
	})

}
