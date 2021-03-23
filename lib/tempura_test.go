package lib

import (
	"testing"
)

func Test_Fill(t *testing.T) {
	type args struct {
		input    Input
		template string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Replace the field",
			args: args{
				input:    Input{"Test": "test"},
				template: "{{.Test}}",
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "Throw a syntax error",
			args: args{
				input:    Input{"Test": "test"},
				template: "{{.Test}",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Throw a set template",
			args: args{
				input:    Input{"Test": "test"},
				template: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempura := Tempura{tt.args.template, "", tt.args.input}

			got, err := tempura.Fill()
			if (err != nil) != tt.wantErr {
				t.Errorf("Fill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FillFromFile(t *testing.T) {
	type args struct {
		inputFilePath    string
		templateFilePath string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantErr1 bool
		wantErr2 bool
	}{
		{
			name: "Replace the field",
			args: args{
				inputFilePath:    "../examples/if-else/input.json",
				templateFilePath: "../examples/if-else/template.tmpl",
			},
			want: `isTrue is True!

isFalse is False!

Cat: nyan

`,
			wantErr1: false,
			wantErr2: false,
		},
		{
			name: "Throw a no such file or directory for input file",
			args: args{
				inputFilePath:    "../nothing_path.json",
				templateFilePath: "../examples/if-else/template.tmpl",
			},
			want:     "",
			wantErr1: true,
			wantErr2: true,
		},
		{
			name: "Throw a no such file or directory for template file",
			args: args{
				inputFilePath:    "../examples/if-else/input.json",
				templateFilePath: "../nothing_path.json",
			},
			want:     "",
			wantErr1: false,
			wantErr2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempura := Tempura{"", tt.args.templateFilePath, nil}

			err := tempura.SetInputFromJSON(tt.args.inputFilePath)
			if (err != nil) != tt.wantErr1 {
				t.Errorf("SetInputFromJSON() error = %v, wantErr %v", err, tt.wantErr1)
				return
			}
			got, err := tempura.Fill()
			if (err != nil) != tt.wantErr2 {
				t.Errorf("Fill() error = %v, wantErr %v", err, tt.wantErr2)
				return
			}
			if got != tt.want {
				t.Errorf("Fill() = %v, want %v", got, tt.want)
			}
		})
	}
}
