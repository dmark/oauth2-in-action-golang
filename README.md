# oauth2-in-action-golang

**Summary**: [_OAuth 2 in Action_](https://www.manning.com/books/oauth-2-in-action) code samples, re-written in golang.

I am learning [golang](https://go.dev/). As an exercise, I will be re-writing the [code samples](https://github.com/oauthinaction/oauth-in-action-code) from _OAuth 2 in Action_, by [Justin Richer](https://github.com/jricher) and [Antonio Sanso](https://github.com/asanso), in Go.

The original sources are written in javascript / node.js. They use express, Underscore.js and Consolidate.js.

## 2025-12-29

What has changed? Other than golang itself!

1. The directory structure for each exercise has changed as follows: within each exercise directory, each of the actors (the client, the authorization server, and the resource server) have their own sub-directories, and each actor's HTML template file is stored in a `files/` directory therein.
1. In place of Express.js, we will use the standard library's `net/http` package.
1. In place of Underscore.js and Consolidate.js, we will use the the standard library's `html/template` package and other built-in functionality.

I've also added a Taskfile.yml so you can launch each exercise with a single `task` (modern-day `make` for old people like me) command. E.g.,

```
$ task ap-A-ex-0
```

The task names match the directory names for each exercise. You run `task [taskname]` from the project root directory. `brew install go-task` if you are on a Mac and using Homebrew.