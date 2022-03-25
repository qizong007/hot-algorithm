package string

// dp[state][char] = next state
type KMP struct {
	dp      [][]int
	pattern string
}

func (k *KMP) Init(pat string) {
	// init
	k.pattern = pat
	k.dp = make([][]int, len(pat))
	for i := range k.dp {
		k.dp[i] = make([]int, 256)
	}
	m := len(pat)
	// pre state
	X := 0
	// base case (state: 0 -> 1)
	k.dp[0][pat[0]] = 1
	// build state machine
	for i := 1; i < m; i++ {
		var b byte
		for b = 0; b <= 255; b++ {
			if pat[i] == b {
				k.dp[i][b] = i + 1
			} else {
				k.dp[i][b] = k.dp[X][b]
			}
			if b == 255 {
				break
			}
		}
		// update pre state
		X = k.dp[X][pat[i]]
	}
}

func (k *KMP) Search(txt string) int {
	m := len(k.pattern)
	n := len(txt)
	state := 0
	for i := 0; i < n; i++ {
		state = k.dp[state][txt[i]]
		if state == m {
			return i - m + 1
		}
	}
	return -1
}
