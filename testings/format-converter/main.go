package testy

type Format struct{}

// func convert(inputFormat Format, params any, fns ...func(in Format, params any) Format) Format {
// 	intermediate := fns(inputFormat, params)
// 	intermediat2 := funcs(inputFormat)
// 	return intermediat2
// }

// func GoogleToSRT(input Format, params any) {
// 	return convert(input, GoogleToParagraph, ParagraphToSrt)
// }

// func AWSToSRT(input Format, params any) {
// 	return convert(input, GoogleToParagraph, ParagraphToSrt)
// }

// func GoogleToSRT2(input Format, params any) {
// 	return convert(input, GoogleToParagraph, P2Txt, Txt2Srt)
// }

// func Convert(input Format, params any) {
// 	inputFormat := string
// 	outFormat := string
// 	switch {
// 	case inputFormat == "google" && outFormat == "srt":
// 		return GoogleToSRT(input, params)
// 	case inputFormat == "aws" && outFormat == "srt":
// 		return AWSToSRT(input, params)
// 	}
// }

// func main() {
// 	googleAPIFormat := Format{}
// 	srtFormat := Convert(googleAPIFormat, {in:,out:,higl})

// 	srt := srtFormat.ToString()
// 	print(srt)
// }
