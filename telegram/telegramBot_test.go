package telegram

import (
	"testing"

	telegrambot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Test_getChatID(t *testing.T) {
	type args struct {
		m *telegrambot.Message
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{
			name: "Test getChatID",
			args: args{
				m: &telegrambot.Message{
					Chat: &telegrambot.Chat{
						ID: -1001923478100,
						Type: "supergroup",
					},
				},
			},
			want: -1923478100,
		},
		{
			name: "Test Notsupergroup",
			args: args{
				m: &telegrambot.Message{
					Chat: &telegrambot.Chat{
						ID: -802956121,
						Type: "group",
					},
				},
			},
			want: -802956121,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getChatID(tt.args.m); got != tt.want {
				t.Errorf("getChatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
