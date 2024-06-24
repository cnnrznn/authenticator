# authenticator

`authenticator` is a byte-sized CLI for storing time-based one time passwords (TOTP).

TOTP is an algorithm for performing multi-factor authentication (MFA).
At a high level, the algorithm is as follows.

1. The service provider and user securely share a secret.
2. Time passes, the user is asked to authenticate using this secret.
3. The user and service provider generate a 6-digit token based on the current time from the epoch, and the secret.
4. If the user provides a matching token to the server, the user is authenticated by the service provider.

## References

- https://www.ietf.org/rfc/rfc6238.txt
- https://github.com/google/google-authenticator/wiki/Key-Uri-Format

## TODO

- Ability to delete entries
- Report a column "expires in" for the user
