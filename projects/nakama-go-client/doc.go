// Package nkclient provides Go clients for connecting to a Nakama server.
//
// The package exposes two constructors at the top level:
//
//   - NewClient(serverKey, host, opts...)  — player-facing API client
//     for the apigrpc surface, used to authenticate end users and perform
//     gameplay calls.
//   - NewAdminClient(host, opts...)        — operator-facing Console API
//     client, used for administrative tasks.
//
// Common connection knobs are supplied via Option (WithHost, WithPort,
// WithSSL, WithTimeout, WithHTTPClient).
package nkclient
