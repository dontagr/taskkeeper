package task

import "testing"

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"ok", "Buy milk", "Buy milk", false},
		{"trim", "  Buy milk  ", "Buy milk", false},
		{"empty", "", "", true},
		{"spaces", "   ", "", true},
		{"too long", string(make([]byte, 201)), "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateTitle(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("got %q want %q", got, tt.want)
			}
		})
	}
}
