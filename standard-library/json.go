package standard_library

import "encoding/json"

type User struct {
	Name     string `json:"name"`
	Password string `json:"-"` // json omit password
}

func (u *User) UnmarshalJSON(data []byte) error {
	type UserTemp User

	temp := struct {
		Password string `json:"password"`
		*UserTemp
	}{
		UserTemp: (*UserTemp)(u),
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	u.Password = temp.Password
	return nil
}
