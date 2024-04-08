package request

import (
	"testing"
)

func TestCreateUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    CreateUser
		wantErr bool
	}{
		{
			name: "valid user",
			user: CreateUser{
				Email:     "test@example.com",
				Password:  "password",
				Name:      "Test User",
				BirthDate: "2000-01-01",
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			user: CreateUser{
				Email:     "invalid email",
				Password:  "password",
				Name:      "Test User",
				BirthDate: "2000-01-01",
			},
			wantErr: true,
		},
		{
			name: "invalid birth date",
			user: CreateUser{
				Email:     "test@example.com",
				Password:  "password",
				Name:      "Test User",
				BirthDate: "01-01-2000",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("createUser.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
