package seabattle

import (
	"reflect"
	"testing"
)

func TestShip_GetCoordinates(t *testing.T) {
	type fields struct {
		kind      ShipKind
		coord     Cell
		direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   []Cell
	}{
		{
			name: "ShipOne in the corner",
			fields: fields{
				kind:      ShipOne,
				coord:     Cell{0, 0},
				direction: DOWN,
			},
			want: []Cell{{0, 0}},
		},
		{
			name: "ShipTwo in the corner",
			fields: fields{
				kind:      ShipTwo,
				coord:     Cell{0, 0},
				direction: DOWN,
			},
			want: []Cell{{0, 0}, {0, 1}},
		},
		{
			name: "ShipThree in the corner",
			fields: fields{
				kind:      ShipThree,
				coord:     Cell{0, 0},
				direction: DOWN,
			},
			want: []Cell{{0, 0}, {0, 1}, {0, 2}},
		},
		{
			name: "ShipFour in the corner",
			fields: fields{
				kind:      ShipFour,
				coord:     Cell{0, 0},
				direction: DOWN,
			},
			want: []Cell{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		},
		{
			name: "ShipFour in the center of the field direction right",
			fields: fields{
				kind:      ShipFour,
				coord:     Cell{5, 5},
				direction: RIGHT,
			},
			want: []Cell{{5, 5}, {6, 5}, {7, 5}, {8, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Ship{
				kind:      tt.fields.kind,
				coord:     tt.fields.coord,
				direction: tt.fields.direction,
			}
			if got := s.GetCoordinates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ship.GetCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShips_GetCoordinates(t *testing.T) {
	shipOne := Ship{
		kind:      ShipOne,
		coord:     Cell{0, 0},
		direction: DOWN,
	}
	shipFour := Ship{
		kind:      ShipFour,
		coord:     Cell{2, 0},
		direction: DOWN,
	}
	tests := []struct {
		name string
		s    Ships
		want map[Cell]*Ship
	}{
		{
			name: "Simple test",
			s:    Ships{&shipOne, &shipFour},
			want: map[Cell]*Ship{
				{0, 0}: &shipOne,
				{2, 0}: &shipFour,
				{2, 1}: &shipFour,
				{2, 2}: &shipFour,
				{2, 3}: &shipFour,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetCoordinates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ships.GetCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
