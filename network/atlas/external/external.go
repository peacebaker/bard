package external

// someday, this will be a served from a db
// until then, we'll just import the variable

var Atlas = Server{
	Host: "localhost",
	Port: "41100",
}

var Backend = Server{
	Host: "localhost",
	Port: "41000",
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Phonebook struct {
	Chat   Server `json:"chat"`
	Feed   Server `json:"feed"`
	Guard  Server `json:"guard"`
	HOA    Server `json:"hoa"`
	House  Server `json:"house"`
	Postal Server `json:"postal"`
	Sus    Server `json:"sus"`
}

// each backend server will maintain an internal list of guardhouses, but no other neighborhood information
// when a request comes in, the backend server will ask the guard
// for the location of the guard server for the requested neighborhood
// the guard server will return a phonebook, which has addresses
// hopefully if we just keep the phonebook cached in memory, it won't cause issues to do this repeatedly.
// I'd like to keep the phonebooks in volatile memory ONLY on the backends for increased security.
