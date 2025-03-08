package main

import (
	"flag"
	"log"

	"github.com/4ears-me/github-app-credential-helper/common"
)

func main() {
	secretArn := flag.String("secret-arn", "", "Secret ARN")
	role := flag.String("role", "", "Role ARN if a role should be assumed")
	tokenCommand := flag.String("token-command", "", "OIDC token command if using web identity")
	flag.Parse()

	subcommand := flag.Arg(0)

	if subcommand == "get" {
		get(secretArn, role, tokenCommand)
	}
}

func get(secretArn *string, role *string, tokenCommand *string) {
	if secretArn == nil || *secretArn == "" {
		log.Fatal("-secret-arn is required")
	}

	if !common.ShouldRun() {
		return
	}

	provider := secretsManagerProvider{
		secretArn:    *secretArn,
		role:         role,
		tokenCommand: tokenCommand,
	}

	helper := common.NewAuthenticator(&provider)
	output, err := helper.Authenticate()
	if err != nil {
		log.Fatal(err)
	}
	print(output)
}
