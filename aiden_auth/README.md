# aiden_auth

A way to authorize FiveM sessions... outside of FiveM. This project is written in TypeScript for FiveM integration, and a HTTP server in Go. It runs another HTTP service on your server, and exposes two HTTP addresses.

This is my first time messing with FiveM, I'm using this project as a way to understand the underlying APIs better.

## Convars

### `aiden_auth_internal_addr`

Set this convar to point the FiveM scripts to your running instance of the Go server, by default, you should set this to `http://localhost:3341`. **Do not include a trailing slash.**

## Exports

This resource provides two [exports](https://docs.fivem.net/docs/scripting-manual/runtimes/lua/#using-exports) that you may use to invoke functions from within the resource.

### `getSessionToken`

Returns the current session token, if there is none, it will return an empty string.

### `forceSessionRefresh`

You should use this to force a reload of the session token, if for some reason the one you have becomes invalid before the expiry. **You should never need to do this, the resource manages token refreshing for you!**

## Dependencies

You will need to know how to setup HTTP servers, use firewalls, and setup a relational database. If you're not confident with these things, you probably shouldn't use this project.

## Building the app

To build the go app, run

```bash
go build -o app cmd/service/main.go
```

Then, run `./app.exe run` or `./app run`, depending on whether you're on Windows Powershell or Linux/Mac Shell.

If you're developing locally, you can skip this skip, just run `go run cmd/service/main.go run`, **you should not use this in prod, let the compiler do its thing!**

## How does it work?

The idea is to push your player session state to a microservice that allows us to uniquely verify people on our own systems, without having to talk to the FiveM server, or have to send anything to the main server to build entire user flows.

We achieve this by building a new API service, that exposes *two* HTTP servers:

- *Private API*: **This server must __not__ be exposed to the public internet or your end users in any way.** This entire flow works because we have reasonable trust that the authority creating these sessions is the FiveM process, and to achieve that, we need an internal API. There is the option to point this at a [Unix socket](https://en.wikipedia.org/wiki/Unix_domain_socket), which eliminates both TCP overhead, and is more secure as you can apply Linux filesystem permissions to this API.

- *Public API*: **This server is safe to expose to the public internet!** You will use this API to build custom FiveM client implementations that can use any system you desire. You can simply throw this API a token, and it'll tell you if it's (a) valid, and (b) who is it. You can also let your other server processes interact with this API, if you want to use this as a session verifier on your own REST API, etc.

### Why would I want to even do this?

There's a few benefits to this:

- You can decouple your server-side code from the actual runtime of the server, that means you can deploy and make updates to that service, without any players on your server ever noticing.

- The FiveM main process has a lot going on, especially with a lot of mods, offloading that work (this even gives you the ability to offload to another server entirely!) means you can remove some excess traffic and bring better performance to that very expensive process.

## Developing

### Generating mocks

To generate mocks when you change code, run

```bash
go generate ./...
```
