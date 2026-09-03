# Contributing to Blog Project

Thank you for your interest in contributing! Every contribution helps make this project better.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [How to Contribute](#how-to-contribute)
- [Coding Standards](#coding-standards)
- [Commit Guidelines](#commit-guidelines)
- [Pull Request Process](#pull-request-process)

## Code of Conduct

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Getting Started

### Prerequisites

- **Go** 1.25 or higher
- **PostgreSQL** running instance
- **Git** installed on your machine

### Development Setup

1. **Fork the repository**

   Click the "Fork" button on the top right of the GitHub repository page.

2. **Clone your fork**

   ```bash
   git clone https://github.com/<your-username>/blog-project.git
   cd blog-project
   ```

3. **Add upstream remote**

   ```bash
   git remote add upstream https://github.com/surajit/blog-project.git
   ```

4. **Set up environment variables**

   ```bash
   cp .env.sample .env
   ```

   Edit `.env` with your database credentials.

5. **Install dependencies**

   ```bash
   go mod tidy
   ```

6. **Run the server**

   ```bash
   go run cmd/server/main.go
   ```

## How to Contribute

### Reporting Bugs

- Check [existing issues](https://github.com/surajit/blog-project/issues) first to avoid duplicates.
- Open a new issue with a clear title and description.
- Include steps to reproduce, expected behavior, and actual behavior.

### Suggesting Features

- Open an issue with the **feature request** label.
- Clearly describe the feature, its use case, and expected behavior.

### Submitting Code Changes

1. Create a new branch from `main`:

   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following the [coding standards](#coding-standards).

3. Test your changes thoroughly.

4. Commit your changes following the [commit guidelines](#commit-guidelines).

5. Push to your fork:

   ```bash
   git push origin feature/your-feature-name
   ```

6. Open a Pull Request against `main`.

## Coding Standards

### General

- Follow standard Go conventions and idioms.
- Run `gofmt` and `goimports` before committing.
- Keep functions small and focused on a single responsibility.
- Use meaningful variable and function names.

### Project Structure

- **Handlers** — HTTP request/response logic only.
- **Services** — Business logic and validation.
- **Repositories** — Database operations using GORM.
- **Models** — Data structures and relationships.

### Error Handling

- Always handle errors explicitly.
- Return meaningful error messages.
- Use the `utils.JSON` helper for consistent API responses.

### Comments

- Add comments for exported functions and complex logic.
- Keep comments concise and up-to-date.

## Commit Guidelines

### Format

```
<type>: <description>
```

### Types

| Type | Description |
|------|-------------|
| `feat` | A new feature |
| `fix` | A bug fix |
| `docs` | Documentation changes |
| `style` | Code style changes (formatting, no logic change) |
| `refactor` | Code restructuring without changing behavior |
| `test` | Adding or updating tests |
| `chore` | Maintenance tasks (dependencies, CI, etc.) |

### Examples

```
feat: add user profile endpoint
fix: resolve null pointer in post handler
docs: update API endpoints in README
refactor: extract validation logic to service layer
```

## Pull Request Process

1. **Title** — Use a clear, descriptive title matching the commit format.
2. **Description** — Explain what changed and why.
3. **Size** — Keep PRs focused. One feature or fix per PR.
4. **Tests** — Add tests for new features when applicable.
5. **Review** — Wait for at least one review before merging.
6. **Merge** — Squash and merge into `main`.

### PR Checklist

- [ ] Code follows project coding standards
- [ ] Changes are tested locally
- [ ] Documentation is updated if needed
- [ ] Commit messages follow guidelines
- [ ] No merge conflicts with `main`

---

Thank you for contributing!
