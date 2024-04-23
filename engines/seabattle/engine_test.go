package seabattle

import (
	"reflect"
	"testing"
)

// This test is a *bit* overloaded, but it does it's job
func TestEngine_Shot(t *testing.T) {
	type fields struct {
		Game  *Game
		Ships Ships
	}
	type args struct {
		targetCells []Cell
	}
	type want struct {
		shots       []map[Cell]bool
		filledCells [][]Cell
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Simple test",
			fields: fields{
				Game: NewGame(),
				Ships: Ships{
					{
						kind:      ShipOne,
						cell:      Cell{1, 1},
						direction: UP,
					},
				},
			},
			args: args{targetCells: []Cell{{1, 1}}},
			want: want{
				shots: []map[Cell]bool{
					{
						{0, 0}: true, {1, 0}: true, {2, 0}: true,
						{0, 1}: true, {1, 1}: true, {2, 1}: true,
						{0, 2}: true, {1, 2}: true, {2, 2}: true,
					},
				},
				filledCells: [][]Cell{{
					{0, 0}, {0, 1}, {0, 2},
					{1, 0}, {1, 1}, {1, 2},
					{2, 0}, {2, 1}, {2, 2},
				}},
			},
			wantErr: false,
		},
		{
			name: "Simple test with larger ship",
			fields: fields{
				Game: NewGame(),
				Ships: Ships{
					{
						kind:      ShipFour,
						cell:      Cell{5, 5},
						direction: UP,
					},
				},
			},
			args: args{targetCells: []Cell{{5, 5}}},
			want: want{
				shots: []map[Cell]bool{{
					{5, 5}: true,
				}},
				filledCells: [][]Cell{{
					{5, 5},
				}},
			},
			wantErr: false,
		},
		{
			name: "Kill the bigger ship test",
			fields: fields{
				Game: NewGame(),
				Ships: Ships{
					{
						kind:      ShipFour,
						cell:      Cell{5, 5},
						direction: UP,
					},
				},
			},
			args: args{
				targetCells: []Cell{
					{5, 2},
					{5, 3},
					{5, 4},
					{5, 5},
				},
			},
			want: want{
				shots: []map[Cell]bool{
					{
						{5, 2}: true,
					},
					{
						{5, 3}: true,
					},
					{
						{5, 4}: true,
					},
					{
						{4, 1}: true,
						{4, 2}: true,
						{4, 3}: true,
						{4, 4}: true,
						{4, 5}: true,
						{4, 6}: true,

						{5, 1}: true,
						{5, 5}: true,
						{5, 6}: true,

						{6, 1}: true,
						{6, 2}: true,
						{6, 3}: true,
						{6, 4}: true,
						{6, 5}: true,
						{6, 6}: true,
					},
				},
				filledCells: [][]Cell{
					{
						{5, 2},
					},
					{
						{5, 3},
					},
					{
						{5, 4},
					},
					{
						{4, 1},
						{4, 2},
						{4, 3},
						{4, 4},
						{4, 5},
						{4, 6},

						{5, 1},
						{5, 5},
						{5, 6},

						{6, 1},
						{6, 2},
						{6, 3},
						{6, 4},
						{6, 5},
						{6, 6},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// place ships on the field
			tt.fields.Game.Field2.Ships = tt.fields.Ships

			e := &Engine{
				Game: tt.fields.Game,
			}
			expectedShots := map[Cell]bool{}
			for i := range tt.args.targetCells {

				for k := range tt.want.shots[i] {
					expectedShots[k] = true
				}

				filledCells, err := e.Shot(tt.args.targetCells[i])
				if (err != nil) != tt.wantErr {
					t.Fatalf("Engine.Shot() error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(tt.want.filledCells[i], filledCells) {
					t.Fatalf("filledCells \nwant:%v\ngot: %v\n", tt.want.filledCells[i], filledCells)
				}
				if !reflect.DeepEqual(expectedShots, e.Game.Field2.Shots) {
					t.Fatalf("Engine.Game.Field2.Shots want: %v\ngot:  %v", tt.want.shots, e.Game.Field2.Shots)
				}
				if e.Game.State.Turn() != "1" {
					t.Fatalf("If player hit the ship his turn should continue")
				}
			}

		})
	}
}
