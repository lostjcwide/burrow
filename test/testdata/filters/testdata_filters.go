package filters

import (
	edb "github.com/eris-ltd/eris-db/erisdb"
	ep "github.com/eris-ltd/eris-db/erisdb/pipe"
	"github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/state"
)

var testDataJson = `{
  "chain_data": {
    "priv_validator": {
      "address": "37236DF251AB70022B1DA351F08A20FB52443E37",
      "pub_key": [1, "CB3688B7561D488A2A4834E1AEE9398BEF94844D8BDBBCA980C11E3654A45906"],
      "priv_key": [1, "6B72D45EB65F619F11CE580C8CAED9E0BADC774E9C9C334687A65DCBAD2C4151CB3688B7561D488A2A4834E1AEE9398BEF94844D8BDBBCA980C11E3654A45906"],
      "last_height": 0,
      "last_round": 0,
      "last_step": 0
    },
    "genesis": {
      "chain_id": "my_tests",
      "accounts": [
        {
          "address": "1000000000000000000000000000000000000000",
          "amount": 0
        },
        {
          "address": "0000000000000000000000000000000000000001",
          "amount": 1
        },
        {
          "address": "0000000000000000000000000000000000000002",
          "amount": 2
        },
        {
          "address": "0000000000000000000000000000000000000003",
          "amount": 3
        },
        {
          "address": "0000000000000000000000000000000000000004",
          "amount": 4
        },
        {
          "address": "0000000000000000000000000000000000000005",
          "amount": 5
        },
        {
          "address": "0000000000000000000000000000000000000006",
          "amount": 6
        },
        {
          "address": "0000000000000000000000000000000000000007",
          "amount": 7
        },
        {
          "address": "0000000000000000000000000000000000000008",
          "amount": 8
        },
        {
          "address": "0000000000000000000000000000000000000009",
          "amount": 9
        },
        {
          "address": "000000000000000000000000000000000000000A",
          "amount": 10
        },
        {
          "address": "000000000000000000000000000000000000000B",
          "amount": 11
        },
        {
          "address": "000000000000000000000000000000000000000C",
          "amount": 12
        },
        {
          "address": "000000000000000000000000000000000000000D",
          "amount": 13
        },
        {
          "address": "000000000000000000000000000000000000000E",
          "amount": 14
        },
        {
          "address": "000000000000000000000000000000000000000F",
          "amount": 15
        }
      ],
      "validators": [
        {
          "pub_key": [1, "CB3688B7561D488A2A4834E1AEE9398BEF94844D8BDBBCA980C11E3654A45906"],
          "amount": 5000000000,
          "unbond_to": [
            {
              "address": "93E243AC8A01F723DE353A4FA1ED911529CCB6E5",
              "amount": 5000000000
            }
          ]
        }
      ]
    }
  },
  "input": {
    "filters0": [
      {
        "field": "balance",
        "op": "==",
        "value": "0"
      }
    ],
    "filters1": [
      {
        "field": "balance",
        "op": ">",
        "value": "12"
      }
    ],
    "filters2": [
      {
        "field": "balance",
        "op": ">=",
        "value": "5"
      },
      {
        "field": "balance",
        "op": "<",
        "value": "8"
      }
    ]
  },
  "output": {
    "accounts0": {
      "accounts": [
        {
          "address": "1000000000000000000000000000000000000000",
          "pub_key": null,
          "sequence": 0,
          "balance": 0,
          "code": "",
          "storage_root": ""
        }
      ]
    },
    "accounts1": {
      "accounts": [
        {
          "address": "000000000000000000000000000000000000000D",
          "pub_key": null,
          "sequence": 0,
          "balance": 13,
          "code": "",
          "storage_root": ""
        },
        {
          "address": "000000000000000000000000000000000000000E",
          "pub_key": null,
          "sequence": 0,
          "balance": 14,
          "code": "",
          "storage_root": ""
        },
        {
          "address": "000000000000000000000000000000000000000F",
          "pub_key": null,
          "sequence": 0,
          "balance": 15,
          "code": "",
          "storage_root": ""
        }
      ]
    },
    "accounts2": {
      "accounts": [
        {
          "address": "0000000000000000000000000000000000000005",
          "pub_key": null,
          "sequence": 0,
          "balance": 5,
          "code": "",
          "storage_root": ""
        },
        {
          "address": "0000000000000000000000000000000000000006",
          "pub_key": null,
          "sequence": 0,
          "balance": 6,
          "code": "",
          "storage_root": ""
        },
        {
          "address": "0000000000000000000000000000000000000007",
          "pub_key": null,
          "sequence": 0,
          "balance": 7,
          "code": "",
          "storage_root": ""
        }
      ]
    }
  }
}`

var serverDuration uint = 100

type (
	ChainData struct {
		PrivValidator *state.PrivValidator `json:"priv_validator"`
		Genesis       *state.GenesisDoc    `json:"genesis"`
	}

	Input struct {
		Filters0 []*ep.FilterData `json:"filters0"`
		Filters1 []*ep.FilterData `json:"filters1"`
		Filters2 []*ep.FilterData `json:"filters2"`
	}

	Output struct {
		Accounts0 *ep.AccountList `json:"accounts0"`
		Accounts1 *ep.AccountList `json:"accounts1"`
		Accounts2 *ep.AccountList `json:"accounts2"`
	}

	TestData struct {
		ChainData *ChainData `json:"chain_data"`
		Input     *Input     `json:"input"`
		Output    *Output    `json:"output"`
	}
)

func LoadTestData() *TestData {
	codec := edb.NewTCodec()
	testData := &TestData{}
	err := codec.DecodeBytes(testData, []byte(testDataJson))
	if err != nil {
		panic(err)
	}
	return testData
}