# git-credential-github-app-sm

[![release](https://github.com/4ears-me/git-credential-github-app-sm/actions/workflows/release.yml/badge.svg)](https://github.com/4ears-me/git-credential-github-app-sm/actions/workflows/release.yml)
![GitHub License](https://img.shields.io/github/license/4ears-me/git-credential-github-app-sm)
![GitHub Tag](https://img.shields.io/github/v/tag/4ears-me/git-credential-github-app-sm)
[![Go Report Card](https://goreportcard.com/badge/github.com/4ears-me/git-credential-github-app-sm)](https://goreportcard.com/report/github.com/4ears-me/git-credential-github-app-sm)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/4ears-me/git-credential-github-app-sm)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/4ears-me/git-credential-github-app-sm?tab=doc)

This is a simple git credential provider for authenticating GitHub apps by installation. The secret must
be stored in AWS Secrets Manager with a string value of the format:

```json
{
  "appId": <application ID>,
  "installationID": <installation ID>,
  "privateKey": <private key in PEM format>
}
```

## Usage
Arguments:
* `-secret-arn <ARN>` - the ARN of the secret
* `-role <ARN>` - (Optional) role ARN to assume before retrieving the secret
* `-token-command "<command>"` - (Optional, requires `-role`) command to generate an OIDC token to use when assuming the role

You can configure the git credential helper either through modifying your `gitconfig` or running:

```shell
git config [--global] credential.https://github.com.helper 'github-app-sm -secret-arn <arn> -role <arn> -token-command "<command>"'
```

For more information on configuring helpers, refer to the [git documentation](https://git-scm.com/docs/gitcredentials).

We also recommend using a credential cache to limit the number of requests to GitHub by running this BEFORE the previous command:

```shell
git config credential.https://github.com.helper
```

### Docker
If you need access to the helper in docker, it can be complicated to install Go (if you don't use it) or get the right
binary. We provide x86-64 and arm64 docker images at `https://gallery.ecr.aws/y1c6o4t8/git-credential-github-app-sm`.

You can use these images like this:

```dockerfile
FROM public.ecr.aws/y1c6o4t8/git-credential-github-app-sm:v0.1.7 AS git-helper

# Build your image
FROM ubuntu:latest
# Do some stuff that you need

# COPY it somewhere on your path
COPY --from=git-helper /git-credential-github-app-sm /usr/bin
```