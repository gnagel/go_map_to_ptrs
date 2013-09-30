package map_to_ptrs

import "testing"
import "github.com/orfjackal/gospec/src/gospec"

func TestMapStringToPtrSpecs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in benchmark mode.")
		return
	}
	r := gospec.NewRunner()
	r.AddSpec(MapStringToPtrSpecs)
	gospec.MainGoTest(r, t)
}

func MapStringToPtrSpecs(c gospec.Context) {

	c.Specify("[MapStringToStringPtrs] Instance of MapStringToPtr", func() {
		value1 := MakeMapStringToStringPtrs(2)
		value2 := MapStringToPtr(value1)

		value1.Set("Bob", nil)
		value1.Set("George", nil)
		c.Expect(value1.String(), gospec.Equals, "Bob = NaN, George = NaN")
		c.Expect(value2.String(), gospec.Equals, "Bob = NaN, George = NaN")
		c.Expect(value2.String(), gospec.Equals, value1.String())
	})

	c.Specify("[MapStringToInt64Ptrs] Instance of MapStringToPtr", func() {
		value1 := MakeMapStringToInt64Ptrs(2)
		value2 := MapStringToPtr(value1)

		value1.Set("Bob", nil)
		value1.Set("George", nil)
		c.Expect(value1.String(), gospec.Equals, "Bob = NaN, George = NaN")
		c.Expect(value2.String(), gospec.Equals, "Bob = NaN, George = NaN")
		c.Expect(value2.String(), gospec.Equals, value1.String())
	})

	c.Specify("[MapStringToFloat64Ptrs] Instance of MapStringToPtr", func() {
		value1 := MakeMapStringToFloat64Ptrs(2)
		value2 := MapStringToPtr(value1)

		value1.Set("Bob", nil)
		value1.Set("George", nil)
		c.Expect(value1.String(), gospec.Equals, "Bob = NaN, George = NaN")
		c.Expect(value2.String(), gospec.Equals, "Bob = NaN, George = NaN")
		c.Expect(value2.String(), gospec.Equals, value1.String())
	})

}
