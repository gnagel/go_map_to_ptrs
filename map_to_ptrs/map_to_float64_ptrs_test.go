package map_to_ptrs

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestMapStringToFloat64PtrsSpecs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in benchmark mode.")
		return
	}
	r := gospec.NewRunner()
	r.AddSpec(MapStringToFloat64PtrsSpecs)
	gospec.MainGoTest(r, t)
}

func MapStringToFloat64PtrsSpecs(c gospec.Context) {

	c.Specify("[MapStringToFloat64Ptrs][Make] returns a reference", func() {
		value := MakeMapStringToFloat64Ptrs(2)
		value.Set("Bob", nil)
		value.Set("George", nil)
		c.Expect(value.Len(), gospec.Equals, 2)

		value2 := func(m MapStringToFloat64Ptrs) MapStringToFloat64Ptrs {
			i := float64(123.456)
			m["Gary"] = &i
			return m
		}(value)

		c.Expect(value2.Len(), gospec.Equals, 3)
		c.Expect(value.Len(), gospec.Equals, 3)
		c.Expect(value2.String(), gospec.Equals, value.String())
	})

	c.Specify("[PtrMapFloat64][Len]", func() {
		value := MakeMapStringToFloat64Ptrs(0)
		c.Expect(value.Len(), gospec.Equals, 0)

		value.Set("Bob", nil)
		c.Expect(value.Len(), gospec.Equals, 1)

		value.Set("George", nil)
		c.Expect(value.Len(), gospec.Equals, 2)

	})

	c.Specify("[PtrMapFloat64][Value]", func() {
		count := float64(123.456)
		value := MakeMapStringToFloat64Ptrs(2)
		value.Set("Bob", &count)
		value.Set("George", nil)

		c.Expect(value.Value("Bob"), gospec.Equals, &count)
		c.Expect(*value.Value("Bob"), gospec.Equals, float64(123.456))
		c.Expect(value.Value("George"), gospec.Satisfies, nil == value.Value("George"))
		c.Expect(value.Value("Cache Miss"), gospec.Satisfies, nil == value.Value("Cache Miss"))
	})

	c.Specify("[PtrMapFloat64][String]", func() {
		count := float64(123.456)
		value := MakeMapStringToFloat64Ptrs(2)
		value.Set("Bob", &count)
		value.Set("George", nil)
		c.Expect(value.String(), gospec.Equals, "Bob = 123.456000, George = NaN")
	})

}
