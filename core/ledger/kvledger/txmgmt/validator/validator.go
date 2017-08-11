/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package validator

import (
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
)

// Validator validates the transactions present in a block and returns a batch that should be used to update the state
type Validator interface {
	ValidateAndPrepareBatch(blockAndPvtdata *ledger.BlockAndPvtData, doMVCCValidation bool) (*privacyenabledstate.UpdateBatch, error)
}

// ErrMissingPvtdata is to be thrown if a collection is present in the public read-write set
// but the corresponding pvt data is missing in the pvt data supplied with the block for validation
// TODO When we implement functionalities planned for the phase-2, this assumption may not be true
type ErrMissingPvtdata struct {
	Msg string
}

// ErrPvtdataHashMissmatch is to be thrown if the hash of a collection present in the public read-write set
// does not match with the corresponding pvt data  supplied with the block for validation
type ErrPvtdataHashMissmatch struct {
	Msg string
}

func (e *ErrMissingPvtdata) Error() string {
	return e.Msg
}

func (e *ErrPvtdataHashMissmatch) Error() string {
	return e.Msg
}
