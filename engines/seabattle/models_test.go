package seabattle

import (
	"reflect"
	"testing"
)

func TestShip_GetCells(t *testing.T) {
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
		{
			name: "ShipFour in the center of the field direction left",
			fields: fields{
				kind:      ShipFour,
				coord:     Cell{5, 5},
				direction: LEFT,
			},
			want: []Cell{
				{2, 5},
				{3, 5},
				{4, 5},
				{5, 5},
			},
		},
		{
			name: "ShipFour in the center of the field direction up",
			fields: fields{
				kind:      ShipFour,
				coord:     Cell{5, 5},
				direction: UP,
			},
			want: []Cell{
				{5, 2},
				{5, 3},
				{5, 4},
				{5, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Ship{
				kind:      tt.fields.kind,
				cell:      tt.fields.coord,
				direction: tt.fields.direction,
			}
			if got := s.GetCells(); !reflect.DeepEqual(got, tt.want) {
				t.Logf("%+v", s)
				t.Errorf("Ship.GetCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShips_GetCells(t *testing.T) {
	shipOne := Ship{
		kind:      ShipOne,
		cell:      Cell{0, 0},
		direction: DOWN,
	}
	shipFour := Ship{
		kind:      ShipFour,
		cell:      Cell{2, 0},
		direction: DOWN,
	}
	tests := []struct {
		name string
		s    Ships
		want map[Cell]Ship
	}{
		{
			name: "Simple test",
			s:    Ships{shipOne, shipFour},
			want: map[Cell]Ship{
				{0, 0}: shipOne,
				{2, 0}: shipFour,
				{2, 1}: shipFour,
				{2, 2}: shipFour,
				{2, 3}: shipFour,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetCells(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ships.GetCells() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestField_PlaceShip(t *testing.T) {
	type fields struct {
		Ships Ships
	}
	type args struct {
		ship Ship
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Ships
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Place ShipOne in the 0, 0",
			fields: fields{
				Ships: Ships{},
			},
			args: args{
				ship: Ship{
					kind:      ShipOne,
					cell:      Cell{0, 0},
					direction: UP,
				},
			},
			want: Ships{
				Ship{
					kind:      ShipOne,
					cell:      Cell{0, 0},
					direction: UP,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				Shots: map[Cell]bool{},
				Ships: tt.fields.Ships,
			}
			if err := f.PlaceShip(tt.args.ship); (err != nil) != tt.wantErr {
				t.Errorf("Field.PlaceShip() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.want, f.Ships) {
				t.Errorf(
					"Field.Ships are not equal\nwant: %v\ngot:  %v",
					tt.want, f.Ships,
				)
			}
		})
	}
}

func TestShip_CellsTaken(t *testing.T) {
	type fields struct {
		kind      ShipKind
		cell      Cell
		direction Direction
	}
	tests := []struct {
		name   string
		fields fields
		want   []Cell
	}{
		{
			name: "ShipOne 0,0",
			fields: fields{
				kind:      ShipOne,
				cell:      Cell{0, 0},
				direction: UP,
			},
			want: []Cell{
				{0, 0}, {0, 1},
				{1, 0}, {1, 1},
			},
		},
		{
			name: "ShipOne 1,1",
			fields: fields{
				kind:      ShipOne,
				cell:      Cell{1, 1},
				direction: UP,
			},
			want: []Cell{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 1}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
		},
		{
			name: "ShipFour UP 5, 5",
			fields: fields{
				kind:      ShipFour,
				cell:      Cell{5, 5},
				direction: UP,
			},
			want: []Cell{
				{4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5}, {4, 6},
				{5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6},
				{6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5}, {6, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Ship{
				kind:      tt.fields.kind,
				cell:      tt.fields.cell,
				direction: tt.fields.direction,
			}
			if got := s.CellsTaken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"Ship.CellsTaken() \ngot:  %d | %v\nwant: %d | %v",
					len(got), got, len(tt.want), tt.want,
				)
			}
		})
	}
}
