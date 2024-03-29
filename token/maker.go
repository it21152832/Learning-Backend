package token

import "time"

//maker is an interface for managing tokens
type Maker interface{
	//createToken creates a new token for a specific username and duration
		CreateToken(email string, duration time.Duration)(string,error)


	//verifyToken check if the the token is valid or not
		VerifyToken(token string) (*Payload, error)
}