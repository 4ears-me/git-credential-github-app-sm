# git-credential-github-app-sm

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
