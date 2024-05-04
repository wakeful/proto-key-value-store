package memory

import (
	"testing"
)

func TestStore_Get(t *testing.T) {
	store := &Store{
		data: map[string]string{
			"question": "what is the answer to life, the universe, and everything?",
			"answer":   "42",
		},
	}

	tests := []struct {
		name       string
		key        string
		want       string
		wantStatus bool
	}{
		{
			name:       "get question",
			key:        "question",
			want:       "what is the answer to life, the universe, and everything?",
			wantStatus: true,
		},
		{
			name:       "get answer",
			key:        "answer",
			want:       "42",
			wantStatus: true,
		},
		{
			name:       "get missing",
			key:        "missing",
			want:       "",
			wantStatus: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotStatus := store.Get(tt.key)
			if got != tt.want {
				t.Errorf("Get() got = %v, want_key %v", got, tt.want)
			}

			if gotStatus != tt.wantStatus {
				t.Errorf("Get() gotStatus = %v, want_key %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestStore_Set(t *testing.T) {
	store := &Store{
		data: map[string]string{
			"question": "what is the answer to life, the universe, and everything?",
			"answer":   "42",
		},
	}

	tests := []struct {
		name      string
		key       string
		value     string
		wantKey   string
		wantValue string
	}{
		{
			name:      "set question",
			key:       "question",
			value:     "what is the answer to life, the universe, and everything? ",
			wantKey:   "question",
			wantValue: "what is the answer to life, the universe, and everything? ",
		},
		{
			name:      "set answer",
			key:       "answer",
			value:     "42",
			wantKey:   "answer",
			wantValue: "42",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotValue := store.Set(tt.key, tt.value)
			if gotKey != tt.wantKey {
				t.Errorf("Set() gotKey = %v, want_key %v", gotKey, tt.wantKey)
			}

			if gotValue != tt.wantValue {
				t.Errorf("Set() gotValue = %v, want_key %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestNewMemoryStore(t *testing.T) {
	store := NewMemoryStore()
	if store == nil {
		t.Error("NewMemoryStore() returned nil")
	}
	if len(store.data) != 0 {
		t.Errorf("NewMemoryStore() data length = %v, want 0", len(store.data))
	}
}
