package lengthoflongestsubstring

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {

	var s = "abababababc"
	i := lengthOfLongestSubstring(s)
	if i != 3 {
		t.Error("case is not pass")
	}

}
