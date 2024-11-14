package main

type testCase struct {
	springs        []byte
	damagedSprings []int64
	want           int64
}

var cases = []testCase{
	{
		springs:        []byte("???"),
		damagedSprings: []int64{1},
		want:           3,
	},
	{
		springs:        []byte("???"),
		damagedSprings: []int64{1, 1},
		want:           1,
	},
	{
		springs:        []byte("#"),
		damagedSprings: []int64{1},
		want:           1,
	},
	{
		springs:        []byte("##"),
		damagedSprings: []int64{1},
		want:           0,
	},
	{
		springs:        []byte("#"),
		damagedSprings: []int64{1, 1},
		want:           0,
	},
	{
		springs:        []byte("#.#"),
		damagedSprings: []int64{1, 1},
		want:           1,
	},
	{
		springs:        []byte("..."),
		damagedSprings: []int64{1, 1},
		want:           0,
	},
	{
		springs:        []byte("???.###"),
		damagedSprings: []int64{1, 1, 3},
		want:           1,
	},
	{
		springs:        []byte(".??..??...?##."),
		damagedSprings: []int64{1, 1, 3},
		want:           4,
	},
	{
		springs:        []byte("?#?#?#?#?#?#?#?"),
		damagedSprings: []int64{1, 3, 1, 6},
		want:           1,
	},
	{
		springs:        []byte("????.#...#..."),
		damagedSprings: []int64{4, 1, 1},
		want:           1,
	},
	{
		springs:        []byte("????.######..#####."),
		damagedSprings: []int64{1, 6, 5},
		want:           4,
	},
	{
		springs:        []byte("?###????????"),
		damagedSprings: []int64{3, 2, 1},
		want:           10,
	},
	{
		springs:        []byte("???.###????.###????.###????.###????.###"),
		damagedSprings: []int64{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3},
		want:           1,
	},
	{
		springs:        []byte(".??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##.?.??..??...?##."),
		damagedSprings: []int64{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3},
		want:           16384,
	},
	{
		springs:        []byte("?###??????????###??????????###??????????###??????????###????????"),
		damagedSprings: []int64{3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1},
		want:           506250,
	},
}
