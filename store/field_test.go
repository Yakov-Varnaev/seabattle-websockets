package store

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FieldTestSuite struct {
	suite.Suite
}

func (suite *FieldTestSuite) TestNewField() {
	field := NewField()

	suite.Equal(len(field.Data), 10, "The field must be 10x10 size")
	suite.Equal(len(field.Data[0]), 10, "The field must be 10x10 size")
	for i := range 10 {
		for j := range 10 {
			cell := field.Data[i][j]
			suite.False(cell.IsShot, "Cell.IsShot must be false by default.")
			suite.Equal(cell.ShipID, 0, "Cell.ShipId must be 0 by default.")
		}
	}
}

func TestFieldTestSuite(t *testing.T) {
	suite.Run(t, new(FieldTestSuite))
}
