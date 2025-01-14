// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package types

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	AccessList        = types.AccessList
	AccessTuple       = types.AccessTuple
	Block             = types.Block
	Blocks            = types.Blocks
	Bloom             = types.Bloom
	Body              = types.Body
	Log               = types.Log
	Receipt           = types.Receipt
	Receipts          = types.Receipts
	ReceiptForStorage = types.ReceiptForStorage
	Transaction       = types.Transaction
	Transactions      = types.Transactions
	Header            = types.Header
	BlockNonce        = types.BlockNonce
	DynamicFeeTx      = types.DynamicFeeTx
	AccessListTx      = types.AccessListTx
	LegacyTx          = types.LegacyTx
	TxData            = types.TxData
	Signer            = types.Signer
	Withdrawal        = types.Withdrawal
	Withdrawals       = types.Withdrawals
)

var (
	NewLondonSigner        = types.NewLondonSigner
	MakeSigner             = types.MakeSigner
	SignTx                 = types.SignTx
	Sender                 = types.Sender
	NewTx                  = types.NewTx
	NewTransaction         = types.NewTransaction
	LatestSignerForChainID = types.LatestSignerForChainID
	SignNewTx              = types.SignNewTx
	MustSignNewTx          = types.MustSignNewTx
	NewBlock               = types.NewBlock
	NewBlockWithHeader     = types.NewBlockWithHeader
)
