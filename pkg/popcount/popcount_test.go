package popcount

import "testing"

var tests = map[uint64]int{
	0:                  0,
	1:                  1,
	2:                  1,
	3:                  2,
	222:                6,
	254:                7,
	255:                8,
	0xFFFFFFFFFFFFFFFF: 64,
}

func TestPopCount1(t *testing.T) {
	for given, want := range tests {
		got := PopCount1(given)
		if got != want {
			t.Errorf("PopCount1(%d) = %d, want %d", given, got, want)
		}
	}
}

func TestPopCount2(t *testing.T) {
	for given, want := range tests {
		got := PopCount2(given)
		if got != want {
			t.Errorf("PopCount2(%d) = %d, want %d", given, got, want)
		}
	}
}

func TestPopCount3(t *testing.T) {
	for given, want := range tests {
		got := PopCount3(given)
		if got != want {
			t.Errorf("PopCount3(%d) = %d, want %d", given, got, want)
		}
	}
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := range tests {
			_ = PopCount1(n)
		}
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := range tests {
			_ = PopCount2(n)
		}
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := range tests {
			_ = PopCount3(n)
		}
	}
}
