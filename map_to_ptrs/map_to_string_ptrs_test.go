package map_to_ptrs

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestMapStringToStringPtrsSpecs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in benchmark mode.")
		return
	}
	r := gospec.NewRunner()
	r.AddSpec(MapStringToStringPtrsSpecs)
	gospec.MainGoTest(r, t)
}

func MapStringToStringPtrsSpecs(c gospec.Context) {

	c.Specify("[MapStringToStringPtrs][Make] returns a reference", func() {
		value := MakeMapStringToStringPtrs(2)
		value.Set("Bob", nil)
		value.Set("George", nil)
		c.Expect(value.Len(), gospec.Equals, 2)

		value2 := func(m MapStringToStringPtrs) MapStringToStringPtrs {
			i := string("123")
			m["Gary"] = &i
			return m
		}(value)

		c.Expect(value2.Len(), gospec.Equals, 3)
		c.Expect(value.Len(), gospec.Equals, 3)
		c.Expect(value2.String(), gospec.Equals, value.String())
	})

	c.Specify("[PtrMapString][Len]", func() {
		value := MakeMapStringToStringPtrs(0)
		c.Expect(value.Len(), gospec.Equals, 0)

		value.Set("Bob", nil)
		c.Expect(value.Len(), gospec.Equals, 1)

		value.Set("George", nil)
		c.Expect(value.Len(), gospec.Equals, 2)

	})

	c.Specify("[PtrMapString][Value]", func() {
		count := string("123")
		value := MakeMapStringToStringPtrs(2)
		value.Set("Bob", &count)
		value.Set("George", nil)

		c.Expect(value.Value("Bob"), gospec.Equals, &count)
		c.Expect(*value.Value("Bob"), gospec.Equals, string("123"))
		c.Expect(value.Value("George"), gospec.Satisfies, nil == value.Value("George"))
		c.Expect(value.Value("Cache Miss"), gospec.Satisfies, nil == value.Value("Cache Miss"))
	})

	c.Specify("[PtrMapString][String]", func() {
		count := string("123")
		value := MakeMapStringToStringPtrs(2)
		value.Set("Bob", &count)
		value.Set("George", nil)
		c.Expect(value.String(), gospec.Equals, "Bob = 123, George = NaN")
	})

}
