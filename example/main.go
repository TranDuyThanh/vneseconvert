package main

import (
	"fmt"
	"github.com/TranDuyThanh/vneseconvert"
)

var sample = "NhiÒu ba ba nhËp khÈu nhiÔm vi khuÈn. Tõ ®Çu n¨m ®Õn nay , Chi côc B¶o vÖ nguån lîi thñy s¶n TP HCM ®· ph¸t hiÖn , trong 71 l« ba ba gièng vµ 5 l« ba ba bè mÑ nhËp khÈu tõ §µi Loan , Th¸i Lan , nhiÒu con bÞ nhiÔm mét sè vi khuÈn g©y bÖnh."

func main() {
	unicodeStr := vneseconvert.TCVN3ToUnicode(sample)
	fmt.Println(unicodeStr)
	fmt.Println(vneseconvert.UnicodeToVNI(unicodeStr))
	fmt.Println(vneseconvert.UnicodeToVIQR(unicodeStr))
}
