package kana

type KanaMapper struct {
	r2k map[string]string
	k2r map[string]string
}

func initKanaMapper(_ [][2]string) *KanaMapper {
	return nil
}

var HiraganaMapper *KanaMapper
var DakutenMapper *KanaMapper
var CombinationMapper *KanaMapper

func init() {
	HiraganaMapper = initKanaMapper([][2]string{
		// a kanas
		{"あ", "a"}, {"い", "i"}, {"う", "u"}, {"え", "e"}, {"お", "o"},
		// ka kanas
		{"か", "ka"}, {"き", "ki"}, {"く", "ku"}, {"け", "ke"}, {"こ", "ko"},
		// sa kanas
		{"さ", "sa"}, {"し", "shi"}, {"す", "su"}, {"せ", "se"}, {"そ", "so"},
		// ta kanas
		{"た", "ta"}, {"ち", "chi"}, {"つ", "tsu"}, {"て", "te"}, {"と", "to"},
		// na kanas
		{"な", "na"}, {"に", "ni"}, {"ぬ", "nu"}, {"ね", "ne"}, {"の", "no"},
		// ha kanas
		{"は", "ha"}, {"ひ", "hi"}, {"ふ", "fu/hu"}, {"へ", "he"}, {"ほ", "ho"},
		// ma kanas
		{"ま", "ma"}, {"み", "mi"}, {"む", "mu"}, {"め", "me"}, {"も", "mo"},
		// y kanas
		{"や", "ya"}, {"ゆ", "yu"}, {"よ", "yo"},
		// ra kana
		{"ら", "ra"}, {"り", "ri"}, {"る", "ru"}, {"れ", "re"}, {"ろ", "ro"},
		// w kana
		{"わ", "wa"}, {"を", "wo"}, {"ん", "n"},
	})

	DakutenMapper = initKanaMapper([][2]string{
		// K -> G
		{"が", "ga"}, {"ぎ", "gi"}, {"ぐ", "gu"}, {"げ", "ge"}, {"ご", "go"},
		// S -> Z
		{"ざ", "za"}, {"じ", "ji"}, {"ず", "zu"}, {"ぜ", "ze"}, {"ぞ", "zo"},
		// T -> D
		{"だ", "da"}, {"ぢ", "zi/di"}, {"づ", "zu/du"}, {"で", "de"}, {"ど", "do"},
		// H -> B
		{"ば", "ba"}, {"び", "bi"}, {"ぶ", "bu"}, {"べ", "be"}, {"ぼ", "bo"},
		// H -> P
		{"ぱ", "pa"}, {"ぴ", "pi"}, {"ぷ", "pu"}, {"ぺ", "pe"}, {"ぽ", "po"},
	})

	CombinationMapper = initKanaMapper([][2]string{
		{"きゃ", "kya"}, {"きゅ", "kyu"}, {"きょ", "kyo"},
		{"ぎゃ", "gya"}, {"ぎゅ", "gyu"}, {"ぎょ", "gyo"},
		{"しゃ", "sha"}, {"しゅ", "shu"}, {"しょ", "sho"},
		{"じゃ", "jya"}, {"じゅ", "jyu"}, {"じょ", "jyo"},
		{"ちゃ", "cha"}, {"ちゅ", "chu"}, {"ちょ", "cho"},
		{"ぢゃ", "dya"}, {"ぢゅ", "dyu"}, {"ぢょ", "dyo"},
		{"にゃ", "nya"}, {"にゅ", "nyu"}, {"にょ", "nyo"},
		{"ひゃ", "hya"}, {"ひゅ", "hyu"}, {"ひょ", "hyo"},
		{"びゃ", "bya"}, {"びゅ", "byu"}, {"びょ", "byo"},
		{"ぴゃ", "pya"}, {"ぴゅ", "pyu"}, {"ぴょ", "pyo"},
		{"みゃ", "mya"}, {"みゅ", "myu"}, {"みょ", "myo"},
		{"りゃ", "rya"}, {"りゅ", "ryu"}, {"りょ", "ryo"},
	})

	// いっしょ
}
