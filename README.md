# oauth2-in-action-golang

**Summary**: [_OAuth 2 in Action_](https://www.manning.com/books/oauth-2-in-action) code samples, re-written in golang.

I am learning [golang](https://go.dev/). As an exercise, I will be re-writing the [code samples](https://github.com/oauthinaction/oauth-in-action-code) from _OAuth 2 in Action_, by [Justin Richer](https://github.com/jricher) and [Antonio Sanso](https://github.com/asanso), in Go.

The original sources are written in javascript / node.js. They use express, Underscore.js and Consolidate.js.

## Submitting PRs

You are welcome to submit PRs and issues, but keep in mind I am doing this as an exercise in learning Go, so I may not merge any PRs or resolve any issues until I have completed the project. Also bear in mind the "Assumptions" below.

## Assumptions

The point of the book is to learn about the inner workings of OAuth2, and not to learn any particular programming language. For this reason, these translated golang versions of the book's exercises will keep things as simple as possible, meaning:

* Using native features in golang and the standard library wherever possible (e.g., `net/http`, not Gin; `html/template`, not Templ)
* Not necessarily doing things the "correct" way, either from a professional software development perspective (I am not a professional software developer), or from an idiomatic golang perspective.

## 2026-01-01

* Happy New Year!
* My `Taskfile.yml` syntax was wrong. The file has been fixed, and optimized to handle the eventual long list of exercises.
* The code and templates for the authorization server and resource server Appendix code are updated. They're a mess, but they work. I'll spend some time cleaning these up before moving on, as these will server as the template for the rest of the exercises.

## 2025-12-29

What has changed? Other than golang itself!

1. The directory structure for each exercise has changed, compared with the original javascript, as follows: within each exercise directory, each of the actors (the client, the authorization server, and the resource server) have their own sub-directories, and each actor's HTML template file is stored in a `files/` directory therein.
1. In place of Express.js, we will use the standard library's `net/http` package.
1. In place of Underscore.js and Consolidate.js, we will use the the standard library's `html/template` package and other built-in functionality.
1. The template files for the sample code from Appendix A have been converted to use `html/template` syntax.
1. The code for the OAuth client in Go is complete. Running the client `main.go` results in the same user experience as the original JavaScript.

I've also added a Taskfile.yml so you can launch each exercise with a single `task` (modern-day `make` for old people like me) command. E.g.,

```
$ task ap-A-ex-0
```

The task names match the directory names for each exercise. Run `task [taskname]` from the project root directory. `brew install go-task` if you are on a Mac and using Homebrew. Using `task` is optional. You can always `go run` the individual `main.go` files for each exercise.