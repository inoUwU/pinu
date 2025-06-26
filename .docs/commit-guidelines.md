# Conventional Commits Usage Guide

> 📘 This document includes material from the [Conventional Commits v1.0.0 Specification](https://www.conventionalcommits.org/en/v1.0.0/), licensed under [Creative Commons Attribution 3.0 International (CC BY 3.0)](https://creativecommons.org/licenses/by/3.0/).  
> Copyright © The Conventional Commits Specification Authors.

---

## ✍️ Commit Message Format

```

<type>(<scope>)\<optional !>: <description>

\[optional body]

\[optional footer(s)]

````

- The message must start with a **type** (`feat`, `fix`, `docs`, etc.).
- A **scope** may be included in parentheses.
- If the commit includes a breaking change, **`!`** must be added before the colon.
- A **description** must follow after a colon and a space.

---

## ✅ Examples

### Basic Feature and Fix

```text
feat(login): add support for OAuth2

fix(parser): handle null inputs correctly
````

### With Body

```text
feat(profile): allow avatar image uploads

This adds support for uploading avatar images using multipart/form-data.
Thumbnails are generated automatically using sharp.
```

### With Footer

```text
fix(auth): reject expired tokens

Tokens are now checked for expiration during validation.

Closes: #42
```

### Breaking Change (in header)

```text
refactor(api)!: drop support for legacy v1 endpoints

BREAKING CHANGE: API clients must now use the v2 endpoints.
```

### Breaking Change (footer only)

```text
chore(deps): update express to v5

BREAKING CHANGE: Express 5 no longer supports middleware chaining via `next('route')`.
```

---

## 📌 Common `type` Values

| Type       | Meaning                       |
| ---------- | ----------------------------- |
| `feat`     | A new feature                 |
| `fix`      | A bug fix                     |
| `docs`     | Documentation only changes    |
| `style`    | Code style (formatting, etc.) |
| `refactor` | Code refactoring              |
| `perf`     | Performance improvement       |
| `test`     | Adding or fixing tests        |
| `chore`    | Build process or tool changes |

You may define and use additional types if needed.

---

## 🔒 Enforcement (Optional)

To enforce Conventional Commits automatically, consider using:

- [`commitlint`](https://github.com/conventional-changelog/commitlint)
- [`husky`](https://typicode.github.io/husky)

These tools help prevent invalid commit messages during local development or CI workflows.

---

## 🔗 References

- [Official Specification](https://www.conventionalcommits.org/en/v1.0.0/)
- [RFC 2119 Keywords](https://www.rfc-editor.org/rfc/rfc2119)
- [CC BY 4.0 License](https://creativecommons.org/licenses/by/4.0/)
