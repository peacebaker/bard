package errs

// this is so common it deserves to go first; the rest are alphabetical
type UWutM8 struct {
	Message string
}

func (e UWutM8) Error() string {
	return e.Message
}

// thrown by sus when password hashing fails
// carried through guard to the backend
type HashFailed struct {
	Message string
	Err     error
}

func (e HashFailed) Error() string {
	return e.Message + "\n" + e.Err.Error()
}

// thrown when JSON decoding fails
type JSONDecodeFailed struct {
	Message string
	Err     error
}

func (e JSONDecodeFailed) Error() string {
	return e.Message + "\n" + e.Err.Error()
}

// thrown when JSON encoding fails
type JSONEncodeFailed struct {
	Message string
	Err     error
}

func (e JSONEncodeFailed) Error() string {
	return e.Message + "\n" + e.Err.Error()
}

// thrown when logins fail because of password validation issues
type LoginFailed struct {
	Message string
}

func (e LoginFailed) Error() string {
	return e.Message
}

// for mongodb failures
type MongoFailed struct {
	Message string
	Err     error
}

func (e MongoFailed) Error() string {
	return e.Message + "\n" + e.Err.Error()
}

type NeighborhoodNotFound struct {
	Message string
}

func (e NeighborhoodNotFound) Error() string {
	return e.Message
}

// thrown when a post request fails,
// message should include
type PostReqFailed struct {
	Message string
	Err     error
}

func (e PostReqFailed) Error() string {
	return e.Message + "\n" + e.Err.Error()
}

type RandSeedFailed struct {
	Message string
	Err     error
}

func (e RandSeedFailed) Error() string {
	return e.Message + "\n" + e.Err.Error()
}

type Unauthorized struct {
	Message string
}

func (e Unauthorized) Error() string {
	return e.Message
}

// thrown when a user already exists in the neighborhood
type UserExists struct {
	Message string
}

func (e UserExists) Error() string {
	return e.Message
}
