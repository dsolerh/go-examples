package replacer

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_importReplacer_replace(t *testing.T) {
	type fields struct {
		replacements Replacements
	}
	type args struct {
		s           *bufio.Scanner
		currentLine []byte
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantBytes   []byte
		wantReplace bool
		wantErr     bool
	}{
		{
			name: "should return the same line and no replace",
			fields: fields{
				replacements: map[string][]byte{
					"potato": []byte("tomato"),
				},
			},
			args: args{
				s:           bufio.NewScanner(strings.NewReader("onion\npeach\n")),
				currentLine: []byte("carrot"),
			},
			wantBytes:   []byte("carrot"),
			wantReplace: false,
		},
		{
			name: "should replace the single import line",
			fields: fields{
				replacements: map[string][]byte{
					"potato": []byte("tomato"),
				},
			},
			args: args{
				s:           bufio.NewScanner(strings.NewReader("does not matter")),
				currentLine: []byte(`import "potato"`),
			},
			wantBytes:   []byte(`import "tomato"`),
			wantReplace: true,
		},
		{
			name: "should replace the multi import block",
			fields: fields{
				replacements: map[string][]byte{
					"potato": []byte("tomato"),
				},
			},
			args: args{
				s:           bufio.NewScanner(strings.NewReader("\t\"carrots\"\n\t\"peach\"\n\t\"potato\"\n)\n")),
				currentLine: []byte("import ("),
			},
			wantBytes:   []byte("import (\n\t\"carrots\"\n\t\"peach\"\n\t\"tomato\"\n)"),
			wantReplace: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &importReplacer{
				replacements: tt.fields.replacements,
			}
			got, got1, err := i.replace(tt.args.s, tt.args.currentLine)
			if (err != nil) != tt.wantErr {
				t.Errorf("importReplacer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.wantBytes) {
				t.Errorf("importReplacer.replace() got = %s, want %s", got, tt.wantBytes)
			}
			if got1 != tt.wantReplace {
				t.Errorf("importReplacer.replace() got1 = %v, want %v", got1, tt.wantReplace)
			}
		})
	}
}
