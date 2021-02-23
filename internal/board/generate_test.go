package board

import (
	"testing"
)

const t_empty = Empty
const t_points = King
const t_point = Knight

func board_set(p Point, ps Points) *Board {
	b := &Board{}
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			b.data[x][y] = &Piece{T: t_empty}
		}
	}

	for _, v := range ps {
		b.data[v.X][v.Y] = &Piece{T: t_points}
	}
	b.data[p.X][p.Y] = &Piece{T: t_point}

	return b
}

func in_board(b *Board, want Points) bool {
	for _, v := range want {
		p := b.data[v.X][v.Y]
		if p == nil || p.T != t_points {
			return false
		}
	}

	return true
}

func out_board(b *Board, want Points) bool {
	for x, v := range b.data {
		for y, b := range v {
			if b != nil && b.T != Empty && b.T != t_point {
				found := false

				for _, p := range want {
					if p.X == x && p.Y == y {
						if b.T == t_points {
							found = true
							break
						} else {
							found = false
							break
						}
					}
				}

				if !found {
					//fmt.Printf("x: %d | y: %d | %v", x, y, b.String())
					return true
				}
			}
		}
	}

	return false
}

func generate_test(t *testing.T, name string, p Point, ps Points, want []Point) {
	b := board_set(p, ps)
	t.Logf("have: %v", ps)
	t.Logf("want: %v", want)

	t.Logf("\n%s", b.String())
	if !in_board(b, want) {
		t.Fatalf("%s is not predictable. points not inside selection", name)
	} else if out_board(b, want) {
		t.Fatalf("%s is not predictable. points outside selection", name)
	}
}

func TestPointsMerge(t *testing.T) {
	ps := Points{
		{4, 3},
	}
	ps.Merge(Points{{3, 4}})

	if !ps.In(Point{4, 3}) || ps.In(Point{3, 4}) {
		t.Fatalf("Merge does not merge...")
	}
}

func TestPointsClean(t *testing.T) {
	ps := Points{
		{1, -4},
		{-6, 0},
		{1, 1},
	}
	ps.Clean()

	if ps.In(Point{1, -4}) || ps.In(Point{-6, 0}) || !ps.In(Point{1, 1}) {
		t.Fatalf("Clean does not invalid points")
	}
}

func TestDiagonal(t *testing.T) {
	p := Point{4, 3}
	generate_test(t, "diagonal", p, p.Diagonal(), []Point{
		{7, 6},
		{6, 5},
		{5, 4},
		{3, 2},
		{2, 1},
		{1, 0},
		{0, 7},
		{1, 6},
		{2, 5},
		{3, 4},
		{5, 2},
		{6, 1},
		{7, 0},
	})

}

func TestHorizontal(t *testing.T) {
	p := Point{4, 3}
	generate_test(t, "horizontal", p, p.Horizontal(), []Point{
		{4, 0},
		{4, 1},
		{4, 2},
		{4, 4},
		{4, 5},
		{4, 6},
		{4, 7},
	})
}

func TestVertical(t *testing.T) {
	p := Point{4, 3}
	generate_test(t, "vertical", p, p.Vertical(), []Point{
		{0, 3},
		{1, 3},
		{2, 3},
		{3, 3},
		{5, 3},
		{6, 3},
		{7, 3},
	})
}

func TestSquare(t *testing.T) {
	p := Point{4, 3}
	generate_test(t, "square", p, p.Square(), []Point{
		{5, 4},
		{5, 3},
		{5, 2},
		{4, 4},
		{4, 2},
		{3, 4},
		{3, 3},
		{3, 2},
	})

	// out of bounds?
	p = Point{7, 7}
	generate_test(t, "square", p, p.Square(), []Point{
		{6, 7},
		{7, 6},
		{6, 6},
	})
}

func TestKnight(t *testing.T) {
	p := Point{4, 3}
	generate_test(t, "knight", p, p.Knight(), []Point{
		{6, 4},
		{6, 2},
		{2, 4},
		{2, 2},
		{5, 5},
		{5, 1},
		{3, 5},
		{3, 1},
	})

	p = Point{0, 1}
	t.Log(p.Knight())
	generate_test(t, "knight", p, p.Knight(), []Point{
		{1, 3},
		{2, 0},
		{2, 2},
	})
}

func TestCorner(t *testing.T) {
	p := Point{4, 3}
	generate_test(t, "corner", p, p.Corner(), []Point{
		{5, 4},
		{5, 2},
		{3, 4},
		{3, 2},
	})
}

func TestDirection(t *testing.T) {
	p := Point{4, 3}
	ps := Points{
		{3, 2},
		{3, 3},
		{3, 4},
		{4, 4},
		{5, 4},
		{5, 3},
		{5, 2},
		{4, 2},
	}
	dirs := []uint8{
		Set(DirUp, DirLeft),
		DirUp,
		Set(DirUp, DirRight),
		DirRight,
		Set(DirDown, DirRight),
		DirDown,
		Set(DirDown, DirLeft),
		DirLeft,
	}

	for i, v := range ps {
		d := p.Direction(v)
		if d != dirs[i] {
			t.Fatalf("%d failed. want: %d, have: %d", i, dirs[i], d)
		}
	}
}

/*
func TestIsDiagonal(t *testing.T) {
	p := Point{4, 3}
	ps := []Point{
		{5, 4},
		{5, 2},
		{3, 4},
		{3, 2},
	}
	fn := []func(dst Point) bool{
		p.IsDiagonalUpRight,
		p.IsDiagonalUpLeft,
		p.IsDiagonalDownRight,
		p.IsDiagonalDownLeft,
	}

	for i, v := range ps {
		f := fn[i]

		if !f(v) {
			t.Fatalf("%d fails.", i)
		}

		for j, v := range ps {
			if i == j {
				continue
			}

			if f(v) {
				t.Fatalf("%d succeeds. %d", j, i)
			}
		}
	}
}

func TestIsHorizontal(t *testing.T) {
	p := Point{4, 3}
	ps := []Point{
		{3, 3},
		{5, 3},
	}
	fn := []func(dst Point) bool{
		p.IsUp,
		p.IsDown,
	}

	for i, v := range ps {
		f := fn[i]

		if !f(v) {
			t.Fatalf("%d fails.", i)
		}

		for j, v := range ps {
			if i == j {
				continue
			}

			if f(v) {
				t.Fatalf("%d succeeds. %d", j, i)
			}
		}
	}

}

func TestIsVertical(t *testing.T) {
	p := Point{4, 3}
	ps := []Point{
		{4, 2},
		{4, 4},
	}
	fn := []func(dst Point) bool{
		p.IsLeft,
		p.IsRight,
	}

	for i, v := range ps {
		f := fn[i]

		if !f(v) {
			t.Fatalf("%d fails.", i)
		}

		for j, v := range ps {
			if i == j {
				continue
			}

			if f(v) {
				t.Fatalf("%d succeeds. %d", j, i)
			}
		}
	}

}
*/
