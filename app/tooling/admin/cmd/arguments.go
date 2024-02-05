package cmd

const (
	network        = "network"
	keyCoin        = "key-coin"
	keyPath        = "key-path"
	keyStorePath   = "key-store-path"
	passPhrase     = "passphrase"
	contractID     = "contract-id"
	balance        = "balance"
	addMoney       = "add-money"
	removeMoney    = "remove-money"
	money          = "money"
	transaction    = "transaction"
	wallet         = "wallet"
	vaultAddress   = "vault-address"
	mountPath      = "mount-path"
	token          = "token"
	keysFolder     = "keys-folder"
	credentialFile = "credential-file"
	dbUser         = "db-user"
	dbPass         = "db-pass"
	dbHost         = "db-host"
	dbName         = "db-name"
)

var shortName = map[string]string{
	// global (persistent)
	network:      "n",
	keyCoin:      "K",
	keyPath:      "k",
	keyStorePath: "P",
	passPhrase:   "p",
	contractID:   "c",

	// contract (child of global)
	balance:     "b",
	addMoney:    "a",
	removeMoney: "r",
	money:       "m",

	// transaction (child of global)
	transaction: "t",

	// wallet (child of global)
	wallet: "w",

	// vault (persistent) (child of global)
	vaultAddress: "a",
	mountPath:    "m",
	token:        "t",

	// vault_addkeys (child of vault)
	keysFolder: "f",

	// vault_init (child of vault)
	credentialFile: "C",
}
