# Contributing to mjwtauth

First off, thanks for taking the time to contribute!
Here are some guidelines to help you get started and make the process smooth for everyone.

---

## How Can I Contribute?

### 1. **Reporting Bugs or Requesting Features**

* Please [open an issue](https://github.com/morgansundqvist/mjwt-auth/issues) describing the problem or your suggestion.
* For bugs, include steps to reproduce, expected behavior, actual behavior, and any error messages.

### 2. **Submitting Pull Requests**

* Fork the repository and create your branch from `main`.
* Follow the existing code style.
* Add/update tests for your changes.
* Run `go test ./...` and make sure all tests pass.
* Document your changes in the code and update the README if needed.
* Open a pull request with a clear description of your change.

---

## Code Style

* Use idiomatic Go formatting (`go fmt`).
* Exported functions/types should have [Godoc](https://blog.golang.org/godoc) comments.
* Try to keep functions small and focused.
* Use error wrapping/context for better error messages.

---

## Running Tests

```bash
go test ./...
```

Add tests for any new features or bug fixes in `_test.go` files.
Coverage and good testing is strongly encouraged!

---

## Making Backwards-Incompatible Changes

* Major breaking changes should go through an issue and discussion first.
* Follow [Semantic Versioning](https://semver.org/).
* For breaking changes, bump the major version (`v2`, `v3`, etc).

---

## Feature Suggestions

* Feel free to open an issue to discuss major new features before starting to code.
* For significant changes, a design proposal is appreciated.

---

## License

By contributing, you agree that your contributions will be licensed under the MIT license.

---

Thank you for helping make `mjwt-auth` better!

---

Let me know if you want to add a **code of conduct**, example PR templates, or more detailed design proposal instructions!
