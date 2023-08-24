package config

import "testing"

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name: "valid config",
			cfg: Config{
				BaseShortURL: "https://example.com",
				ServerAddr:   "localhost:8080",
			},
			wantErr: false,
		},
		{
			name: "invalid base short URL",
			cfg: Config{
				BaseShortURL: "",
				ServerAddr:   "localhost:8080",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.validate()

			if len(tt.cfg.ServerAddr) < 5 {
				t.Error("bad server address")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
