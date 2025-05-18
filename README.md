# mjwtauth

A pluggable, reusable JWT authentication package for Go, supporting both HS256 (symmetric) and RS256 (asymmetric, RSA) signing methods. Designed for easy integration into any Go REST API or service.

---

## Features

* **User Signup & Login** with password hashing (bcrypt)
* **JWT Generation & Verification** (HS256 or RS256)
* **Pluggable Claims Customization**: Inject your own claims/scopes per user
* **User & Repository Abstractions**: Bring your own user struct and storage backend
* **Clean interfaces** for easy extension and testability
* **Ready for microservices**: Use public key verification for distributed systems
* **Minimal dependencies** (only `golang-jwt/jwt/v5` and `bcrypt`)

---

## Installation

```bash
go get github.com/morgansundqvist/mjwt-auth
```

---

## Quickstart

### 1. **Implement the User and UserRepository interfaces in your app:**

```go
type User struct {
    ID           string
    Username     string
    PasswordHash string
}

func (u *User) GetID() string           { return u.ID }
func (u *User) GetUsername() string     { return u.Username }
func (u *User) GetPasswordHash() string { return u.PasswordHash }

type InMemoryUserRepo struct {
    users map[string]*User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{users: make(map[string]*User)}
}

func (r *InMemoryUserRepo) FindByUsername(username string) (mjwtauth.AuthUser, error) {
    u, ok := r.users[username]
    if !ok {
        return nil, errors.New("user not found")
    }
    return u, nil
}

func (r *InMemoryUserRepo) CreateUser(username, hash string) (mjwtauth.AuthUser, error) {
    id := uuid.New().String()
    user := &User{ID: id, Username: username, PasswordHash: hash}
    r.users[username] = user
    return user, nil
}
```

---

### 2. **Create a signer (HS256 or RS256):**

#### **HS256 Example**

```go
signer := mjwtauth.NewHS256Signer([]byte("your-secret-key"))
```

#### **RS256 Example**

```go
privKeyPEM, _ := os.ReadFile("private.pem")
pubKeyPEM, _ := os.ReadFile("public.pem")
signer, err := mjwtauth.NewRS256SignerFromPEM(privKeyPEM, pubKeyPEM)
if err != nil {
    log.Fatal(err)
}
```

---

### 3. **Set up a Claims Customizer (optional)**

```go
claimsCustomizer := func(user mjwtauth.AuthUser, claims map[string]interface{}) error {
    claims["role"] = "admin" // Add custom role or scopes
    return nil
}
```

---

### 4. **Create the AuthService**

```go
repo := NewInMemoryUserRepo()
auth := mjwtauth.NewAuthService(repo, signer, claimsCustomizer)
```

---

### 5. **Signup and Login**

```go
user, err := auth.Signup("morgan", "mysecretpassword")
token, err := auth.Login("morgan", "mysecretpassword")
fmt.Println("JWT token:", token)
```

---

### 6. **Verify a JWT Token**

```go
claims, err := auth.VerifyToken(token)
if err != nil {
    fmt.Println("Invalid token:", err)
} else {
    fmt.Println("Claims:", claims)
}
```

---

## API

### Interfaces

#### `AuthUser`

Implement in your app’s user type.

```go
type AuthUser interface {
    GetID() string
    GetUsername() string
    GetPasswordHash() string
}
```

#### `UserRepository`

Your user storage backend.

```go
type UserRepository interface {
    FindByUsername(username string) (AuthUser, error)
    CreateUser(username, hashedPassword string) (AuthUser, error)
}
```

#### `JWTSigner`

Implemented by the package for HS256 or RS256.

```go
type JWTSigner interface {
    Sign(claims map[string]interface{}) (string, error)
    Verify(tokenString string) (jwt.MapClaims, error)
}
```

#### `ClaimsCustomizer`

A callback to modify/add claims for the JWT token.

```go
type ClaimsCustomizer func(user AuthUser, claims map[string]interface{}) error
```

---

## Advanced Usage

* **Custom Password Policies**: Add your own validation before calling `Signup`.
* **Support for context.Context**: Future releases may add context for cancellation/tracing.
* **Refresh Tokens**: Not included by default; you can extend the service for your needs.

---

## Key Rotation & Security

* **HS256**: All verifiers and issuers must share the same secret.
* **RS256**: Only the AuthService holds the private key; verifiers need the public key.
* Rotate keys by deploying new PEM files and restarting services.

---

## Testing

* You can easily mock the interfaces for unit tests.
* Run tests with:

  ```bash
  go test ./...
  ```

---

## Contributing

PRs and issues welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

---

## License

MIT © Morgan Sundqvist, 2025

---

## Credits

* Uses [`golang-jwt/jwt`](https://github.com/golang-jwt/jwt)
* Inspired by best practices in scalable Go authentication

