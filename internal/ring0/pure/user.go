package pure

import "time"

// NOTE: Standard library non-I/O utility packages (time, errors, math) are 
// explicitly permitted in the Pure layer. Only I/O packages are banned.
type User struct {
	ID        string
	Email     string
	IsActive  bool
	CreatedAt time.Time
}