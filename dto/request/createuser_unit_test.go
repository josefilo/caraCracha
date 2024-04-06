package request

import (
    "testing"
)

func TestCreateUser_Validate(t *testing.T) {
    tests := []struct {
        name    string
        user    createUser
        wantErr bool
    }{
        {
            name: "valid user",
            user: createUser{
                email:    "test@example.com",
                password: "password",
                name:     "Test User",
                birthDate: "2000-01-01",
            },
            wantErr: false,
        },
        {
            name: "invalid email",
            user: createUser{
                email:    "invalid email",
                password: "password",
                name:     "Test User",
                birthDate: "2000-01-01",
            },
            wantErr: true,
        },
        {
			name: "invalid birth date",
			user: createUser{
				email:    "test@example.com",
                password: "password",
                name:     "Test User",
                birthDate: "01-01-2000",
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