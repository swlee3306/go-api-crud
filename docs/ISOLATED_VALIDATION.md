# Isolated MySQL validation

Verified on 2026-09-14 using a dedicated Colima VM with no host-directory mounts
or SSH-agent forwarding. Only this example application and an empty MySQL 8
database ran on an internal Docker network, with synthetic test credentials.
No operational database or company service was contacted.

- `go test ./...`, `go vet ./...`, and `go build ./...` passed.
- The initial MySQL migration failed with error 1830: required `category_id`
  conflicted with an `ON DELETE SET NULL` foreign key.
- Both category relationship declarations now use `ON DELETE RESTRICT`.
  A referenced category cannot be hard-deleted; posts must first be reassigned
  or removed. The required category ID and its JSON type are unchanged.
- `TestRequiredCategoryRestrictsDeletion` failed before the fix and passed
  afterward. It checks both GORM relationship declarations.
- Rebuilt application and MySQL containers became healthy. An HTTP request
  inside the application container returned healthy application/database checks.
- Running the image with an empty JWT secret exited with code 1 before database
  access, without including a secret value in the error.

This is an isolated startup/migration/security-configuration smoke test, not a
production deployment or comprehensive CRUD/authentication/end-to-end test.
Host-to-container forwarding was unavailable in this isolated VM setup; the
health request was therefore made inside the container, not via the host port.
