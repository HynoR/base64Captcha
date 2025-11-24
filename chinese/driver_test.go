package chinese

import (
	"image/color"
	"reflect"
	"testing"
)

func TestNewDriver(t *testing.T) {
	type args struct {
		height          int
		width           int
		noiseCount      int
		showLineOptions int
		length          int
		source          string
		bgColor         *color.RGBA
		fonts           []string
	}
	tests := []struct {
		name string
		args args
		want *Driver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriver(tt.args.height, tt.args.width, tt.args.noiseCount, tt.args.showLineOptions, tt.args.length, tt.args.source, tt.args.bgColor, tt.args.fonts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_GenerateIdQuestionAnswer(t *testing.T) {
	tests := []struct {
		name        string
		d           *Driver
		wantId      string
		wantContent string
		wantAnswer  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotContent, gotAnswer := tt.d.GenerateIdQuestionAnswer()
			if gotId != tt.wantId {
				t.Errorf("Driver.GenerateIdQuestionAnswer() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotContent != tt.wantContent {
				t.Errorf("Driver.GenerateIdQuestionAnswer() gotContent = %v, want %v", gotContent, tt.wantContent)
			}
			if gotAnswer != tt.wantAnswer {
				t.Errorf("Driver.GenerateIdQuestionAnswer() gotAnswer = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}

func TestDriver_DrawCaptcha(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		d       *Driver
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.d.DrawCaptcha(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Driver.DrawCaptcha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetFont(t *testing.T) {
	f := GetFont()
	if f == nil {
		t.Error("GetFont() returned nil")
	}
}
