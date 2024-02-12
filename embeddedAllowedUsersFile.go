package main

import (
	_ "embed"
)

//go:embed allowedUsers/allowedUsers.json
var allowedUsers []byte
