package board

import "testing"

func TestGetKing(t *testing.T) {
	const ourking = 28
	const theirking = 4

	have := GetKing(true)
	if ourking != have {
		t.Fatalf("GetQueen: want: %d - have: %d", ourking, have)
	}
	have = GetKing(false)
	if theirking != have {
		t.Fatalf("GetQueen: want: %d - have: %d", theirking, have)
	}
}

func TestGetQueen(t *testing.T) {
	const ourqueen = 27
	const theirqueen = 3

	have := GetQueen(true)
	if ourqueen != have {
		t.Fatalf("GetQueen: want: %d - have: %d", ourqueen, have)
	}
	have = GetQueen(false)
	if theirqueen != have {
		t.Fatalf("GetQueen: want: %d - have: %d", theirqueen, have)
	}
}

func verify_by_2int(t *testing.T, p1 bool, w1, w2 int8, fn func(bool) [2]int8) {
	vals := fn(p1)
	h1, h2 := vals[0], vals[1]

	if w1 != h1 || w2 != h2 {
		t.Fatalf("want: %d / have: %d", w1, h1)
		t.Fatalf("want: %d / have: %d", w2, h2)
	}
}

func TestGetBishops(t *testing.T) {
	verify_by_2int(t, true, 26, 29, GetBishops)
	verify_by_2int(t, false, 2, 5, GetBishops)
}

func TestGetKnights(t *testing.T) {
	verify_by_2int(t, true, 25, 30, GetKnights)
	verify_by_2int(t, false, 1, 6, GetKnights)
}

func TestGetRooks(t *testing.T) {
	verify_by_2int(t, true, 24, 31, GetRooks)
	verify_by_2int(t, false, 0, 7, GetRooks)
}

func TestGetRange(t *testing.T) {
	ourarr := []int8{}
	for i := int8(16); i < 32; i++ {
		ourarr = append(ourarr, i)
	}

	theirarr := []int8{}
	for i := int8(0); i < 16; i++ {
		theirarr = append(theirarr, i)
	}

	ourwant := GetRange(true)
	theirwant := GetRange(false)

	for i := 0; i < 16; i++ {
		ourhave := ourarr[i]
		theirhave := theirarr[i]

		if ourhave != ourwant[i] {
			t.Fatalf("getRange does not match want. want: %d | have: %d", ourwant[i], ourhave)
		}
		if theirhave != theirwant[i] {
			t.Fatalf("getRange does not match want. want: %d | have: %d", theirwant[i], theirhave)
		}
	}
}

func TestGetStartRow(t *testing.T) {
	brd := NewBoard()
	_, rook, err := brd.Get(Point{0, GetStartRow(true)})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if rook.Kind != Rook || rook.P1 != true {
		t.Log(rook)
		t.Fatalf("Piece is not rook, or the player does not match")
	}

	_, rook, err = brd.Get(Point{0, GetStartRow(false)})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if rook.Kind != Rook || rook.P1 != false {
		t.Log(rook)
		t.Fatalf("Piece is not rook, or the player does not match")
	}
}

func TestGetPawnRow(t *testing.T) {
	brd := NewBoard()
	_, rook, err := brd.Get(Point{0, GetPawnRow(true)})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if rook.Kind != Pawn || rook.P1 != true {
		t.Log(rook)
		t.Fatalf("Piece is not rook, or the player does not match")
	}

	_, rook, err = brd.Get(Point{0, GetPawnRow(false)})
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	if rook.Kind != Pawn || rook.P1 != false {
		t.Log(rook)
		t.Fatalf("Piece is not rook, or the player does not match")
	}
}

func TestIsIDValid(t *testing.T) {
	if IsIDValid(32) {
		t.Fatalf("IsIDValid(32): want: false | have: true")
	}
	if IsIDValid(-1) {
		t.Fatalf("IsIDValid(0): want: false | have: true")
	}
}

func TestGetEighthRank(t *testing.T) {
	if GetEighthRank(true) != 0 {
		t.Fatalf("EighthRank(p1) != 0")
	}
	if GetEighthRank(false) != 7 {
		t.Fatalf("EighthRank(p2) != 7")
	}
}

func TestBelongsTo(t *testing.T) {
	ourvalues := [16]int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, v := range ourvalues {
		if !BelongsTo(v, false) {
			t.Fatalf("%d should belong to id: '%d'", v, 0)
		}
	}

	theirvalues := [16]int8{16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	for _, v := range theirvalues {
		if !BelongsTo(v, true) {
			t.Fatalf("%d should belong to id: '%d'", v, 1)
		}
	}
}
