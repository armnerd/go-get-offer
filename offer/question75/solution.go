package question75

// 字符流中第一个不重复的字符
func Run() string {
	return "done!"
}

var (
	seq   []byte
	count map[byte]int
)

func init() {
	seq = make([]byte, 0)
	count = make(map[byte]int)
}

func Insert(char byte) {
	num, ok := count[char]
	if !ok {
		seq = append(seq, char)
	}
	if num == 1 {
		for k, v := range seq {
			if v == char {
				seq = append(seq[:k], seq[k+1:]...)
			}
		}
	}
	count[char] += 1
}

func FirstAppearingOnce() byte {
	if len(seq) == 0 {
		return '#'
	} else {
		return seq[0]
	}
}
