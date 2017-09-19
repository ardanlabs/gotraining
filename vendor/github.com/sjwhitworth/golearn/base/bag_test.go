package base

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestBAGSimple(t *testing.T) {
	Convey("Given certain bit data", t, func() {
		// Generate said bits
		bVals := [][]byte{
			[]byte{1, 0, 0},
			[]byte{0, 1, 0},
			[]byte{0, 0, 1},
		}

		// Create a new DenseInstances
		inst := NewDenseInstances()
		for i := 0; i < 3; i++ {
			inst.AddAttribute(NewBinaryAttribute(fmt.Sprintf("%d", i)))
		}

		// Get and re-order the attributes
		attrSpecsUnordered := ResolveAllAttributes(inst)
		attrSpecs := make([]AttributeSpec, 3)
		for _, a := range attrSpecsUnordered {
			name := a.GetAttribute().GetName()
			So(name, ShouldBeIn, []string{"0", "1", "2"})

			if name == "0" {
				attrSpecs[0] = a
			} else if name == "1" {
				attrSpecs[1] = a
			} else if name == "2" {
				attrSpecs[2] = a
			}
		}

		inst.Extend(3)

		for row, b := range bVals {
			for col, c := range b {
				inst.Set(attrSpecs[col], row, []byte{c})
			}
		}

		Convey("All the row values should be the right length...", func() {
			inst.MapOverRows(attrSpecs, func(row [][]byte, i int) (bool, error) {
				for i := range attrSpecs {
					So(len(row[i]), ShouldEqual, 1)
				}
				return true, nil
			})
		})

		Convey("All the values should be the same...", func() {
			inst.MapOverRows(attrSpecs, func(row [][]byte, i int) (bool, error) {
				for j := range attrSpecs {
					So(row[j][0], ShouldEqual, bVals[i][j])
				}
				return true, nil
			})
		})

	})
}

func TestBAG(t *testing.T) {
	Convey("Given randomly generated bit data", t, func() {
		// Generate said bits
		bVals := make([][]byte, 0)
		for i := 0; i < 50; i++ {
			b := make([]byte, 3)
			for j := 0; j < 3; j++ {
				if rand.NormFloat64() >= 0 {
					b[j] = byte(1)
				} else {
					b[j] = byte(0)
				}
			}
			bVals = append(bVals, b)
		}

		// Create a new DenseInstances
		inst := NewDenseInstances()
		for i := 0; i < 3; i++ {
			inst.AddAttribute(NewBinaryAttribute(fmt.Sprintf("%d", i)))
		}

		// Get and re-order the attributes
		attrSpecsUnordered := ResolveAllAttributes(inst)
		attrSpecs := make([]AttributeSpec, 3)
		for _, a := range attrSpecsUnordered {
			name := a.GetAttribute().GetName()
			So(name, ShouldBeIn, []string{"0", "1", "2"})

			if name == "0" {
				attrSpecs[0] = a
			} else if name == "1" {
				attrSpecs[1] = a
			} else if name == "2" {
				attrSpecs[2] = a
			}
		}

		inst.Extend(50)

		for row, b := range bVals {
			for col, c := range b {
				inst.Set(attrSpecs[col], row, []byte{c})
			}
		}

		Convey("All the row values should be the right length...", func() {
			inst.MapOverRows(attrSpecs, func(row [][]byte, i int) (bool, error) {
				for i := range attrSpecs {
					So(len(row[i]), ShouldEqual, 1)
				}
				return true, nil
			})
		})

		Convey("All the values should be the same...", func() {
			inst.MapOverRows(attrSpecs, func(row [][]byte, i int) (bool, error) {
				for j := range attrSpecs {
					So(row[j][0], ShouldEqual, bVals[i][j])
				}
				return true, nil
			})
		})

	})
}
