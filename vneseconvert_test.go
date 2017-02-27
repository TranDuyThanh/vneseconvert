package vneseconvert

import (
	"testing"
)

var (
	tcvn3Str   = "NhiÒu ba ba nhËp khÈu nhiÔm vi khuÈn. Tõ ®Çu n¨m ®Õn nay , Chi côc B¶o vÖ nguån lîi thñy s¶n TP HCM ®· ph¸t hiÖn , trong 71 l« ba ba gièng vµ 5 l« ba ba bè mÑ nhËp khÈu tõ §µi Loan , Th¸i Lan , nhiÒu con bÞ nhiÔm mét sè vi khuÈn g©y bÖnh."
	vniStr     = "Nhieàu ba ba nhaäp khaåu nhieãm vi khuaån. Töø ñaàu naêm ñeán nay , Chi cuïc Baûo veä nguoàn lôïi thuûy saûn TP HCM ñaõ phaùt hieän , trong 71 loâ ba ba gioáng vaø 5 loâ ba ba boá meï nhaäp khaåu töø Ñaøi Loan , Thaùi Lan , nhieàu con bò nhieãm moät soá vi khuaån gaây beänh."
	viqrStr    = "Nhie^`u ba ba nha^.p kha^?u nhie^~m vi khua^?n. Tu+` dda^`u na(m dde^'n nay , Chi cu.c Ba?o ve^. nguo^`n lo+.i thu?y sa?n TP HCM dda~ pha't hie^.n , trong 71 lo^ ba ba gio^'ng va` 5 lo^ ba ba bo^' me. nha^.p kha^?u tu+` DDa`i Loan , Tha'i Lan , nhie^`u con bi. nhie^~m mo^.t so^' vi khua^?n ga^y be^.nh."
	unicodeStr = "Nhiều ba ba nhập khẩu nhiễm vi khuẩn. Từ đầu năm đến nay , Chi cục Bảo vệ nguồn lợi thủy sản TP HCM đã phát hiện , trong 71 lô ba ba giống và 5 lô ba ba bố mẹ nhập khẩu từ Đài Loan , Thái Lan , nhiều con bị nhiễm một số vi khuẩn gây bệnh."
)

func Test_TCVN3ToUnicode(t *testing.T) {
	if TCVN3ToUnicode(tcvn3Str) != unicodeStr {
		t.Fatal("failed")
	}
}

func Test_VNIToUnicode(t *testing.T) {
	if VNIToUnicode(vniStr) != unicodeStr {
		t.Fatal("failed")
	}
}

func Test_VIQRToUnicode(t *testing.T) {
	if VIQRToUnicode(viqrStr) != unicodeStr {
		t.Fatal("failed")
	}
}

func Test_UnicodeToTCVN3(t *testing.T) {
	if UnicodeToTCVN3(unicodeStr) != tcvn3Str {
		t.Fatal("failed")
	}
}

func Test_UnicodeToVNI(t *testing.T) {
	if UnicodeToVNI(unicodeStr) != vniStr {
		t.Fatal("failed")
	}
}

func Test_UnicodeToVIQR(t *testing.T) {
	if UnicodeToVIQR(unicodeStr) != viqrStr {
		t.Fatal("failed")
	}
}
