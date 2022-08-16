# algo-multi-sig-manager
This repository contains the backend code for storing Multi sig transaction infor and responsible for broadcasting it to the blockchain once all signs are done

## How To Build the Project
You need executed the following commands.

This command will be executed only once to fetch all the dependencies.<br>
`go mod tidy` 

To the run the project run <br>
`go run cmd/main.go`

# Go Packages

## Controllers
- authctrl (Authorization)
- multisigaccountctrl (MultiSig Account)
- transactionctrl (Transaction)
- signedtransactionctrl (Signed Transaction)

## Service
- authsvc (Authorization)
- multisigaccountsvc (MultiSig Account)
- transactionsvc (Transaction)
- signedtransactionsvc (Signed Transaction)

## Utils
- algoutil (General Algorand utilities)
- apiutil (General API and Gin utilities)
- dbutil (Shared logic and General DB utilities)
- jwtutil (JWToken utilities)
- paginateutil (Utility to handle pagination for lists)
- testutil (Utilities to write unit tests easily)
- viperutil (Utilities to load and bind config from viper)
