package gn

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Message("test message")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "test message") {
		t.Errorf("Message() output = %q, want to contain %q", output, "test message")
	}
}

func TestWarn(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Warn("warning message")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "⚠️") {
		t.Errorf("Warn() should contain warning icon")
	}
	if !strings.Contains(output, "warning message") {
		t.Errorf("Warn() output = %q, want to contain %q", output, "warning message")
	}
}

func TestInfo(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Info("info message")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "ℹ️") {
		t.Errorf("Info() should contain info icon")
	}
	if !strings.Contains(output, "info message") {
		t.Errorf("Info() output = %q, want to contain %q", output, "info message")
	}
}

func TestProgress(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Progress("progress message")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "⏳") {
		t.Errorf("Progress() should contain progress icon")
	}
	if !strings.Contains(output, "progress message") {
		t.Errorf("Progress() output = %q, want to contain %q", output, "progress message")
	}
}

func TestSuccess(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Success("success message")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "✅") {
		t.Errorf("Success() should contain success icon")
	}
	if !strings.Contains(output, "success message") {
		t.Errorf("Success() output = %q, want to contain %q", output, "success message")
	}
}

func TestMessageWithVars(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Message("test %s %d", "hello", 42)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "test hello 42") {
		t.Errorf("Message() with vars output = %q, want to contain %q", output, "test hello 42")
	}
}

func TestUserMsg_colorize(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		contains []string
	}{
		{
			name:     "title tag",
			msg:      "This is a <title>title</title> message",
			contains: []string{"title"},
		},
		{
			name:     "warn tag",
			msg:      "This is a <warn>warning</warn> message",
			contains: []string{"warning"},
		},
		{
			name:     "em tag",
			msg:      "This is an <em>emphasized</em> message",
			contains: []string{"emphasized"},
		},
		{
			name:     "err tag",
			msg:      "This is an <err>error</err> message",
			contains: []string{"error"},
		},
		{
			name:     "multiple tags",
			msg:      "<title>Title</title> with <warn>warning</warn> and <em>emphasis</em>",
			contains: []string{"Title", "warning", "emphasis"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			um := userMsg{
				msgType: unknownMsgType,
				msg:     tt.msg,
			}
			result := um.colorize(tt.msg)

			for _, substr := range tt.contains {
				if !strings.Contains(result, substr) {
					t.Errorf("colorize() = %q, want to contain %q", result, substr)
				}
			}
		})
	}
}

func TestUserMsg_print(t *testing.T) {
	tests := []struct {
		name     string
		userMsg  userMsg
		contains []string
	}{
		{
			name: "info message",
			userMsg: userMsg{
				msgType: infoMsgType,
				msg:     "info test",
			},
			contains: []string{"ℹ️", "info test"},
		},
		{
			name: "success message",
			userMsg: userMsg{
				msgType: successMsgType,
				msg:     "success test",
			},
			contains: []string{"✅", "success test"},
		},
		{
			name: "warning message",
			userMsg: userMsg{
				msgType: warningMsgType,
				msg:     "warning test",
			},
			contains: []string{"⚠️", "warning test"},
		},
		{
			name: "error message",
			userMsg: userMsg{
				msgType: errorMsgType,
				msg:     "error test",
			},
			contains: []string{"❌", "error test"},
		},
		{
			name: "progress message",
			userMsg: userMsg{
				msgType: progressMsgType,
				msg:     "progress test",
			},
			contains: []string{"⏳", "progress test"},
		},
		{
			name: "message with vars",
			userMsg: userMsg{
				msgType: infoMsgType,
				msg:     "value: %d",
				vars:    []any{42},
			},
			contains: []string{"ℹ️", "value: 42"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			tt.userMsg.print()

			w.Close()
			os.Stdout = old

			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			for _, substr := range tt.contains {
				if !strings.Contains(output, substr) {
					t.Errorf("print() output = %q, want to contain %q", output, substr)
				}
			}
		})
	}
}

func TestMsgType(t *testing.T) {
	// Test that msgType constants are defined correctly
	tests := []struct {
		name    string
		msgType msgType
		value   msgType
	}{
		{"unknown", unknownMsgType, 0},
		{"error", errorMsgType, 1},
		{"warning", warningMsgType, 2},
		{"info", infoMsgType, 3},
		{"progress", progressMsgType, 4},
		{"success", successMsgType, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.msgType != tt.value {
				t.Errorf("%s msgType = %d, want %d", tt.name, tt.msgType, tt.value)
			}
		})
	}
}

func ExampleMessage() {
	Message("This is a general message")
}

func ExampleWarn() {
	Warn("This is a warning: %s", "something went wrong")
}

func ExampleInfo() {
	Info("Processing <em>%d</em> items", 100)
}

func ExampleProgress() {
	Progress("Loading <title>data</title>...")
}

func ExampleSuccess() {
	Success("Operation completed successfully!")
}
