**_THIS FILE GENERATE BY AI_**

# Git Commit Message Convention

## Format

```
<type>(<scope>): <short summary>

[optional body]

[optional footer]
```

## Types

Must be one of the following:

| Type         | Description                                                            |
| ------------ | ---------------------------------------------------------------------- |
| **feat**     | A new feature or functionality                                         |
| **fix**      | A bug fix                                                              |
| **docs**     | Documentation changes (e.g., README, comments, markdown files)         |
| **docs-api** | API documentation generation                                           |
| **style**    | Changes that don't affect code meaning (formatting, white-space, etc.) |
| **refactor** | Code changes that neither fix bugs nor add features                    |
| **perf**     | Performance improvements                                               |
| **test**     | Adding or modifying tests                                              |
| **chore**    | Build process, tooling changes, dependency updates, etc.               |
| **ci**       | Changes to CI/CD configuration and scripts                             |
| **proto**    | Generated code from protobuf files                                     |
| **revert**   | Reverting a previous commit                                            |
| **rename**   | File or directory renaming without code changes                        |

## Scope (Optional)

The scope provides context about which part of the codebase is affected.

Examples:

- `feat(auth)`: New feature in auth module
- `fix(api)`: Bug fix in API
- `chore(deps)`: Dependency upgrades

## Subject

- Use imperative, present tense ("add" not "added" or "adds")
- Don't capitalize the first letter
- No period (.) at the end
- Keep it concise, ideally under 50 characters

## Body (Optional)

- Use imperative, present tense
- Include motivation for the change and contrast with previous behavior
- Wrap lines at 72 characters

## Footer (Optional)

- Reference issues the commit closes (e.g., "Closes #123")
- Breaking changes should start with "BREAKING CHANGE:"

## Workflow-Specific Examples

### 1. Documentation Generation

```
docs(user): update user registration guide

Add section on password requirements and account verification process
```

### 2. API Documentation Generation

```
docs-api(payment): generate OpenAPI documentation

Generated API docs for new payment endpoints using Swagger annotations
```

### 3. Generated Go Code from Protobuf

```
proto(user): regenerate Go files from updated user.proto

Update models to include new user preference fields
```

### 4. Dependency Upgrades

```
chore(deps): upgrade golang from 1.19 to 1.20

Improves performance and security with latest Golang features
```

### 5. Script Development

```
feat(scripts): create automated database backup script

Add daily backups with configurable retention policy
```

### 6. Refactoring and File Renaming

```
refactor(handlers): simplify error handling in HTTP responses

Extract common error handling logic into shared middleware
```

```
rename(utils): standardize helper function file names

Change from camelCase to snake_case for consistency
```

## Additional Common Examples

```
feat(auth): implement multi-factor authentication
```

```
fix(login): resolve infinite loading on failed login attempts
```

```
chore(ci): update GitHub Actions workflow
```

```
style(components): format code according to style guide
```

```
test(api): add integration tests for user endpoints
```

```
perf(queries): optimize database indexes for faster lookups
```
