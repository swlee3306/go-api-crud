# Public configuration boundary

Reviewed against the repository source on 2026-09-14. This is a configuration
review, not a verified container deployment guide.

## Current code versus legacy Compose

`config/database.go` reads the variables in the right column. The existing
`docker-compose.yml` supplies the different names in the left column.

| Existing Compose variable | Variable read by current Go code |
| --- | --- |
| `DB_CONNECTION` | `DB_DRIVER` |
| `DB_DATABASE` | `DB_NAME` |
| `DB_USERNAME` | `DB_USER` |

`DB_HOST`, `DB_PORT` and `DB_PASSWORD` are also read by the code. `main.go` reads
`JWT_SECRET`, `SERVER_HOST` and `SERVER_PORT`. Missing values can fall back to
development defaults; do not rely on those defaults for a deployment.

The legacy Compose file includes fixed credential candidates, publishes database
and Redis ports, and mounts local files for initialization and TLS. Its production
use is unknown. Changing only a password is not sufficient to establish that the
stack is correct or safe. Do not run it as a portfolio smoke test.

## Separate public guidance from an actual deployment

- The README uses explicit `REPLACE_WITH_...` placeholders. They must not be used
  as actual credentials. Provide a unique signing secret and appropriate database
  credentials through your chosen secret-management mechanism.
- Keep any deployment-specific configuration outside the tracked repository.
  Do not overwrite the existing Compose file or `.env.example` in a shared or
  operating checkout based solely on this review.
- If producing a new Compose example, use a distinct filename and an explicit
  project name. Verify variable names, least-privilege DB access, required mounts,
  port bindings and image/build compatibility before recommending it.
- Configuration rendering must use synthetic secrets only: rendered Compose output
  can contain resolved values. Do not paste real rendered configuration into logs.
- Container startup, database initialization, HTTP authentication, persistence and
  cleanup need a separate isolated test. No containers were started in this review.

## Unchanged and unresolved

The existing Compose file, Dockerfile, `.env.example` and runtime Go defaults were
not changed. In particular, the Dockerfile copies `.env.example` into the image as
`.env`; this review does not assert that this file is loaded at runtime.
Replacing defaults in the runtime is a behavior change requiring focused tests.
Historical credentials remain in Git history until separately addressed, and no
credential validity, expiry, rotation or live deployment was checked.
