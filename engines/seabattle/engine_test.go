package seabattle

import (
	"reflect"
	"testing"
)

func TestEngine_Shot(t *testing.T) {
	type fields struct {
		Game  *Game
		Ships Ships
	}
	type args struct {
		targetCells []Cell
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[Cell]bool
		wantErr bool
	}{
		{
			name: "Simple test",
			fields: fields{
				Game: NewGame(),
				Ships: Ships{
					{
						kind:      ShipOne,
						coord:     Cell{1, 1},
						direction: UP,
					},
				},
			},
			args: args{targetCells: []Cell{{1, 1}}},
			want: map[Cell]bool{
				{0, 0}: true, {1, 0}: true, {2, 0}: true,
				{0, 1}: true, {1, 1}: true, {2, 1}: true,
				{0, 2}: true, {1, 2}: true, {2, 2}: true,
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
						coord:     Cell{5, 5},
						direction: UP,
					},
				},
			},
			args: args{targetCells: []Cell{{5, 5}}},
			want: map[Cell]bool{
				{5, 5}: true,
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
						coord:     Cell{5, 5},
						direction: UP,
					},
				},
			},
			args: args{targetCells: []Cell{
				{5, 2},
				{5, 3},
				{5, 4},
				{5, 5},
			}},
			want: map[Cell]bool{
				{4, 1}: true,
				{4, 2}: true,
				{4, 3}: true,
				{4, 4}: true,
				{4, 5}: true,
				{4, 6}: true,

				{5, 1}: true,
				{5, 2}: true,
				{5, 3}: true,
				{5, 4}: true,
				{5, 5}: true,
				{5, 6}: true,

				{6, 1}: true,
				{6, 2}: true,
				{6, 3}: true,
				{6, 4}: true,
				{6, 5}: true,
				{6, 6}: true,
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
			for _, shotCell := range tt.args.targetCells {
				if err := e.Shot(shotCell); (err != nil) != tt.wantErr {
					t.Errorf("Engine.Shot() error = %v, wantErr %v", err, tt.wantErr)
				}
			}

			if !reflect.DeepEqual(tt.want, e.Game.Field2.Shots) {
				t.Errorf("Engine.Game.Field2.Shots want=%v, got %v", tt.want, e.Game.Field2.Shots)
			}
		})
	}
}
