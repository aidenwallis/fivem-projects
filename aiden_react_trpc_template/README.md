# aiden_react_trpc_template

A FiveM resource that combines that lets you build a strongly-typed, TypeScript server/client using [trpc](https://trpc.io) and [React](https://reactjs.org) for your [NUI](https://docs.fivem.net/docs/scripting-manual/nui-development/full-screen-nui/) interface.

This project makes use of [aiden_auth](../aiden_auth) for authenticating requests to the server, and you get full, end-to-end type safety when writing NUI code.

## Installation

To use this, add the following to your `server.cfg`

```
ensure aiden_auth
ensure aiden_react_trpc_template
```

Note, you must setup [aiden_auth](../aiden_auth) and have a running instance of the server first, and configure it in `config.json`.

Once `aiden_auth` handshakes and generates a session, it'll get passed all the way down to the react context and the UI will load.

## Developing

To develop with the tool, start the [aiden_auth](../aiden_auth/) server/DB locally, then run `yarn start`.

Start FiveM and connect to your server, if you ensured the resource, you should see black text in the top left, that's the resource. You can modify the `src/client` directory and write react code there, then `ensure aiden_react_trpc_template` to reload it.
