# safesql

Go safe SQL implementation

In order to use this library the following steps must be taken:

- Set up your CI/tests so that you can:
  - Ban imports of a specific package/function
  - Create an allowlist of call sites that can use that package/function
- Create an atomic change that:
  - Converts all calls to `database/sql` into calls to `safesql`. This can easily be achieved with the `legacyconversions` package and automated patching.
  - Prevents new calls to `legacyconversions` from being added and bans import of the `database/sql` package. This should ideally be true for all transitive dependencies.
  - Only allows `safesql` to import `database/sql`.
- After submitting that change, gradually migrate `legacyconversions` calls to use `safesql` functions or be promoted to `uncheckedconversions`. If you chose the latter make sure the strings that you promote are controlled by the programmer and never by the user.
