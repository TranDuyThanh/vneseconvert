package vneseconvert

import (
	"sort"
	"strings"
)

var (
	unicode = []string{
		"À", "Á", "Â", "Ã", "È", "É", "Ê", "Ì", "Í", "Ò",
		"Ó", "Ô", "Õ", "Ù", "Ú", "Ý", "à", "á", "â", "ã",
		"è", "é", "ê", "ì", "í", "ò", "ó", "ô", "õ", "ù",
		"ú", "ý", "Ă", "ă", "Đ", "đ", "Ĩ", "ĩ", "Ũ", "ũ",
		"Ơ", "ơ", "Ư", "ư", "Ạ", "ạ", "Ả", "ả", "Ấ", "ấ",
		"Ầ", "ầ", "Ẩ", "ẩ", "Ẫ", "ẫ", "Ậ", "ậ", "Ắ", "ắ",
		"Ằ", "ằ", "Ẳ", "ẳ", "Ẵ", "ẵ", "Ặ", "ặ", "Ẹ", "ẹ",
		"Ẻ", "ẻ", "Ẽ", "ẽ", "Ế", "ế", "Ề", "ề", "Ể", "ể",
		"Ễ", "ễ", "Ệ", "ệ", "Ỉ", "ỉ", "Ị", "ị", "Ọ", "ọ",
		"Ỏ", "ỏ", "Ố", "ố", "Ồ", "ồ", "Ổ", "ổ", "Ỗ", "ỗ",
		"Ộ", "ộ", "Ớ", "ớ", "Ờ", "ờ", "Ở", "ở", "Ỡ", "ỡ",
		"Ợ", "ợ", "Ụ", "ụ", "Ủ", "ủ", "Ứ", "ứ", "Ừ", "ừ",
		"Ử", "ử", "Ữ", "ữ", "Ự", "ự", "Ỳ", "ỳ", "Ỵ", "ỵ",
		"Ỷ", "ỷ", "Ỹ", "ỹ",
	}

	tcvn3 = []string{
		"Aµ", "A¸", "¢", "A·", "EÌ", "EÐ", "£", "I×", "IÝ", "Oß",
		"Oã", "¤", "Oâ", "Uï", "Uó", "Yý", "µ", "¸", "©", "·",
		"Ì", "Ð", "ª", "×", "Ý", "ß", "ã", "«", "â", "ï",
		"ó", "ý", "¡", "¨", "§", "®", "IÜ", "Ü", "Uò", "ò",
		"¥", "¬", "¦", "­", "A¹", "¹", "A¶", "¶", "¢Ê", "Ê",
		"¢Ç", "Ç", "¢È", "È", "¢É", "É", "¢Ë", "Ë", "¡¾", "¾",
		"¡»", "»", "¡¼", "¼", "¡½", "½", "¡Æ", "Æ", "EÑ", "Ñ",
		"EÎ", "Î", "EÏ", "Ï", "£Õ", "Õ", "£Ò", "Ò", "£Ó", "Ó",
		"£Ô", "Ô", "£Ö", "Ö", "IØ", "Ø", "IÞ", "Þ", "Oä", "ä",
		"Oá", "á", "¤è", "è", "¤å", "å", "¤æ", "æ", "¤ç", "ç",
		"¤é", "é", "¥í", "í", "¥ê", "ê", "¥ë", "ë", "¥ì", "ì",
		"¥î", "î", "Uô", "ô", "Uñ", "ñ", "¦ø", "ø", "¦õ", "õ",
		"¦ö", "ö", "¦÷", "÷", "¦ù", "ù", "Yú", "ú", "Yþ", "þ",
		"Yû", "û", "Yü", "ü",
	}

	vni = []string{
		"AØ", "AÙ", "AÂ", "AÕ", "EØ", "EÙ", "EÂ", "Ì", "Í", "OØ",
		"OÙ", "OÂ", "OÕ", "UØ", "UÙ", "YÙ", "aø", "aù", "aâ", "aõ",
		"eø", "eù", "eâ", "ì", "í", "oø", "où", "oâ", "oõ", "uø",
		"uù", "yù", "AÊ", "aê", "Ñ", "ñ", "Ó", "ó", "UÕ", "uõ",
		"Ô", "ô", "Ö", "ö", "AÏ", "aï", "AÛ", "aû", "AÁ", "aá",
		"AÀ", "aà", "AÅ", "aå", "AÃ", "aã", "AÄ", "aä", "AÉ", "aé",
		"AÈ", "aè", "AÚ", "aú", "AÜ", "aü", "AË", "aë", "EÏ", "eï",
		"EÛ", "eû", "EÕ", "eõ", "EÁ", "eá", "EÀ", "eà", "EÅ", "eå",
		"EÃ", "eã", "EÄ", "eä", "Æ", "æ", "Ò", "ò", "OÏ", "oï",
		"OÛ", "oû", "OÁ", "oá", "OÀ", "oà", "OÅ", "oå", "OÃ", "oã",
		"OÄ", "oä", "ÔÙ", "ôù", "ÔØ", "ôø", "ÔÛ", "ôû", "ÔÕ", "ôõ",
		"ÔÏ", "ôï", "UÏ", "uï", "UÛ", "uû", "ÖÙ", "öù", "ÖØ", "öø",
		"ÖÛ", "öû", "ÖÕ", "öõ", "ÖÏ", "öï", "YØ", "yø", "Î", "î",
		"YÛ", "yû", "YÕ", "yõ",
	}

	viqr = []string{
		"A`", "A'", "A^", "A~", "E`", "E'", "E^", "I`", "I'", "O`",
		"O'", "O^", "O~", "U`", "U'", "Y'", "a`", "a'", "a^", "a~",
		"e`", "e'", "e^", "i`", "i'", "o`", "o'", "o^", "o~", "u`",
		"u'", "y'", "A(", "a(", "DD", "dd", "I~", "i~", "U~", "u~",
		"O+", "o+", "U+", "u+", "A.", "a.", "A?", "a?", "A^'", "a^'",
		"A^`", "a^`", "A^?", "a^?", "A^~", "a^~", "A^.", "a^.", "A('", "a('",
		"A(`", "a(`", "A(?", "a(?", "A(~", "a(~", "A(.", "a(.", "E.", "e.",
		"E?", "e?", "E~", "e~", "E^'", "e^'", "E^`", "e^`", "E^?", "e^?",
		"E^~", "e^~", "E^.", "e^.", "I?", "i?", "I.", "i.", "O.", "o.",
		"O?", "o?", "O^'", "o^'", "O^`", "o^`", "O^?", "o^?", "O^~", "o^~",
		"O^.", "o^.", "O+'", "o+'", "O+`", "o+`", "O+?", "o+?", "O+~", "o+~",
		"O+.", "o+.", "U.", "u.", "U?", "u?", "U+'", "u+'", "U+`", "u+`",
		"U+?", "u+?", "U+~", "u+~", "U+.", "u+.", "Y`", "y`", "Y.", "y.",
		"Y?", "y?", "Y~", "y~",
	}

	unicodeToTCVN3 = []string{}
	unicodeToVNI   = []string{}
	unicodeToVIQR  = []string{}
	tcvn3ToUnicode = []string{}
	vniToUnicode   = []string{}
	viqrToUnicode  = []string{}

	TCVN3ToUnicodeMap = map[string]string{}
	VNIToUnicodeMap   = map[string]string{}
	VIQRToUnicodeMap  = map[string]string{}

	TCVN3ToUnicodeReplacer *strings.Replacer
	VNIToUnicodeReplacer   *strings.Replacer
	VIQRToUnicodeReplacer  *strings.Replacer
	UnicodeToTCVN3Replacer *strings.Replacer
	UnicodeToVNIReplacer   *strings.Replacer
	UnicodeToVIQRReplacer  *strings.Replacer
)

func init() {

	for i, char := range unicode {
		unicodeToTCVN3 = append(unicodeToTCVN3, char, tcvn3[i])
		unicodeToVNI = append(unicodeToVNI, char, vni[i])
		unicodeToVIQR = append(unicodeToVIQR, char, viqr[i])

		TCVN3ToUnicodeMap[tcvn3[i]] = char
		VNIToUnicodeMap[vni[i]] = char
		VIQRToUnicodeMap[viqr[i]] = char
	}

	sort.Sort(ByLength(tcvn3))
	sort.Sort(ByLength(vni))
	sort.Sort(ByLength(viqr))

	for _, char := range tcvn3 {
		tcvn3ToUnicode = append(tcvn3ToUnicode, char, TCVN3ToUnicodeMap[char])
	}

	for _, char := range vni {
		vniToUnicode = append(vniToUnicode, char, VNIToUnicodeMap[char])
	}

	for _, char := range viqr {
		viqrToUnicode = append(viqrToUnicode, char, VIQRToUnicodeMap[char])
	}

	UnicodeToTCVN3Replacer = strings.NewReplacer(unicodeToTCVN3...)
	UnicodeToVNIReplacer = strings.NewReplacer(unicodeToVNI...)
	UnicodeToVIQRReplacer = strings.NewReplacer(unicodeToVIQR...)
	TCVN3ToUnicodeReplacer = strings.NewReplacer(tcvn3ToUnicode...)
	VNIToUnicodeReplacer = strings.NewReplacer(vniToUnicode...)
	VIQRToUnicodeReplacer = strings.NewReplacer(viqrToUnicode...)
}

func TCVN3ToUnicode(str string) string {
	return TCVN3ToUnicodeReplacer.Replace(str)
}
func VNIToUnicode(str string) string {
	return VNIToUnicodeReplacer.Replace(str)
}
func VIQRToUnicode(str string) string {
	return VIQRToUnicodeReplacer.Replace(str)
}
func UnicodeToTCVN3(str string) string {
	return UnicodeToTCVN3Replacer.Replace(str)
}
func UnicodeToVNI(str string) string {
	return UnicodeToVNIReplacer.Replace(str)
}
func UnicodeToVIQR(str string) string {
	return UnicodeToVIQRReplacer.Replace(str)
}

// var tcvn3 = []rune{
// 	0x00B8, 0x00B5, 0x00B6, 0x00B7, 0x00B9, 0x00A1, 0x00BE,
// 	0x00BB, 0x00BC, 0x00BD, 0x00C6, 0x00A2, 0x00CA, 0x00C7, 0x00C8,
// 	0x00C9, 0x00CB, 0x00A7, 0x00D0, 0x00CC, 0x00CE, 0x00CF, 0x00D1,
// 	0x00A3, 0x00D5, 0x00D2, 0x00D3, 0x00D4, 0x00D6, 0x00DD, 0x00D7,
// 	0x00D8, 0x00DC, 0x00DE, 0x00E3, 0x00DF, 0x00E1, 0x00E2, 0x00E4,
// 	0x00A4, 0x00E8, 0x00E5, 0x00E6, 0x00E7, 0x00E9, 0x00A5, 0x00ED,
// 	0x00EA, 0x00EB, 0x00EC, 0x00EE, 0x00F3, 0x00EF, 0x00F1, 0x00F2,
// 	0x00F4, 0x00A6, 0x00F8, 0x00F5, 0x00F6, 0x00F7, 0x00F9, 0x00B8,
// 	0x00B5, 0x00B6, 0x00B7, 0x00B9, 0x00A8, 0x00BE, 0x00BB, 0x00BC,
// 	0x00BD, 0x00C6, 0x00A9, 0x00CA, 0x00C7, 0x00C8, 0x00C9, 0x00CB,
// 	0x00AE, 0x00D0, 0x00CC, 0x00CE, 0x00CF, 0x00D1, 0x00AA, 0x00D5,
// 	0x00D2, 0x00D3, 0x00D4, 0x00D6, 0x00DD, 0x00D7, 0x00D8, 0x00DC,
// 	0x00DE, 0x00E3, 0x00DF, 0x00E1, 0x00E2, 0x00E4, 0x00AB, 0x00E8,
// 	0x00E5, 0x00E6, 0x00E7, 0x00E9, 0x00AC, 0x00ED, 0x00EA, 0x00EB,
// 	0x00EC, 0x00EE, 0x00F3, 0x00EF, 0x00F1, 0x00F2, 0x00F4, 0x00AD,
// 	0x00F8, 0x00F5, 0x00F6, 0x00F7, 0x00F9, 0x00FD, 0x00FA, 0x00FB,
// 	0x00FC, 0x00FE,
// }

// var unicode = []rune{
// 	0x00C1, 0x00C0, 0x1EA2, 0x00C3, 0x1EA0, 0x0102, 0x1EAE,
// 	0x1EB0, 0x1EB2, 0x1EB4, 0x1EB6, 0x00C2, 0x1EA4, 0x1EA6, 0x1EA8,
// 	0x1EAA, 0x1EAC, 0x0110, 0x00C9, 0x00C8, 0x1EBA, 0x1EBC, 0x1EB8,
// 	0x00CA, 0x1EBE, 0x1EC0, 0x1EC2, 0x1EC4, 0x1EC6, 0x00CD, 0x00CC,
// 	0x1EC8, 0x0128, 0x1ECA, 0x00D3, 0x00D2, 0x1ECE, 0x00D5, 0x1ECC,
// 	0x00D4, 0x1ED0, 0x1ED2, 0x1ED4, 0x1ED6, 0x1ED8, 0x01A0, 0x1EDA,
// 	0x1EDC, 0x1EDE, 0x1EE0, 0x1EE2, 0x00DA, 0x00D9, 0x1EE6, 0x0168,
// 	0x1EE4, 0x01AF, 0x1EE8, 0x1EEA, 0x1EEC, 0x1EEE, 0x1EF0, 0x00E1,
// 	0x00E0, 0x1EA3, 0x00E3, 0x1EA1, 0x0103, 0x1EAF, 0x1EB1, 0x1EB3,
// 	0x1EB5, 0x1EB7, 0x00E2, 0x1EA5, 0x1EA7, 0x1EA9, 0x1EAB, 0x1EAD,
// 	0x0111, 0x00E9, 0x00E8, 0x1EBB, 0x1EBD, 0x1EB9, 0x00EA, 0x1EBF,
// 	0x1EC1, 0x1EC3, 0x1EC5, 0x1EC7, 0x00ED, 0x00EC, 0x1EC9, 0x0129,
// 	0x1ECB, 0x00F3, 0x00F2, 0x1ECF, 0x00F5, 0x1ECD, 0x00F4, 0x1ED1,
// 	0x1ED3, 0x1ED5, 0x1ED7, 0x1ED9, 0x01A1, 0x1EDB, 0x1EDD, 0x1EDF,
// 	0x1EE1, 0x1EE3, 0x00FA, 0x00F9, 0x1EE7, 0x0169, 0x1EE5, 0x01B0,
// 	0x1EE9, 0x1EEB, 0x1EED, 0x1EEF, 0x1EF1, 0x00FD, 0x1EF3, 0x1EF7,
// 	0x1EF9, 0x1EF5,
// }
