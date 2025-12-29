# oauth2-in-action-golang

**Summary**: [_OAuth 2 in Action_](https://www.manning.com/books/oauth-2-in-action) code samples, re-written in golang.

I am learning [golang](https://go.dev/). As an exercise, I will be re-writing the [code samples](https://github.com/oauthinaction/oauth-in-action-code) from _OAuth 2 in Action_, by [Justin Richer](https://github.com/jricher) and [Antonio Sanso](https://github.com/asanso), in go.

The original sources are written in javascript / node.js. They use express, Underscore.js and Consolidate.js.

In place of Express.js, we will use the standard library's `net/http` package.

In golang, the standard library's `html/template` package and other built-in functionality in place of Underscore.js and Consolidate.js.