package mjwtauth

// UserRepository is implemented by your app to provide user storage/retrieval.
type UserRepository interface {
	FindByUsername(username string) (AuthUser, error)
	CreateUser(username, hashedPassword string) (AuthUser, error)
	// Optionally: FindByID, UpdateUser, etc.
}
