package ssv

import (
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/types"
)

// DutyRunners is a map of duty runners mapped by msg id hex.
type DutyRunners map[types.BeaconRole]Runner

// DutyRunnerForMsgID returns a Runner from the provided msg ID, or nil if not found
func (ci DutyRunners) DutyRunnerForMsgID(msgID types.MessageID) Runner {
	role := msgID.GetRoleType()
	return ci[role]
}

type Network interface {
	Broadcast(message types.Encoder) error
}

// Storage is a persistent storage for the SSV
type Storage interface {
}

// AttesterCalls interface has all attester duty specific calls
type AttesterCalls interface {
	// GetAttestationData returns attestation data by the given slot and committee index
	GetAttestationData(slot phase0.Slot, committeeIndex phase0.CommitteeIndex) (*phase0.AttestationData, error)
	// SubmitAttestation submit the attestation to the node
	SubmitAttestation(attestation *phase0.Attestation) error
}

// ProposerCalls interface has all block proposer duty specific calls
type ProposerCalls interface {
	// GetBeaconBlock returns beacon block by the given slot and committee index
	GetBeaconBlock(slot phase0.Slot, committeeIndex phase0.CommitteeIndex, graffiti, randao []byte) (*bellatrix.BeaconBlock, error)
	// SubmitBeaconBlock submit the block to the node
	SubmitBeaconBlock(block *bellatrix.SignedBeaconBlock) error
}

// AggregatorCalls interface has all attestation aggregator duty specific calls
type AggregatorCalls interface {
	// SubmitAggregateSelectionProof returns an AggregateAndProof object
	SubmitAggregateSelectionProof(slot phase0.Slot, committeeIndex phase0.CommitteeIndex, slotSig []byte) (*phase0.AggregateAndProof, error)
	// SubmitSignedAggregateSelectionProof broadcasts a signed aggregator msg
	SubmitSignedAggregateSelectionProof(msg *phase0.SignedAggregateAndProof) error
}

// SyncCommitteeCalls interface has all sync committee duty specific calls
type SyncCommitteeCalls interface {
	// GetSyncMessageBlockRoot returns beacon block root for sync committee
	GetSyncMessageBlockRoot() (phase0.Root, error)
	// SubmitSyncMessage submits a signed sync committee msg
	SubmitSyncMessage(msg *altair.SyncCommitteeMessage) error
}

// SyncCommitteeContributionCalls interface has all sync committee contribution duty specific calls
type SyncCommitteeContributionCalls interface {
	// GetSyncSubcommitteeIndex returns sync committee indexes for aggregator
	GetSyncSubcommitteeIndex(slot phase0.Slot, pubKey phase0.BLSPubKey) ([]uint64, error)
	// IsSyncCommitteeAggregator returns tru if aggregator
	IsSyncCommitteeAggregator(proof []byte) (bool, error)
	// SyncCommitteeSubnetID returns sync committee subnet ID from subcommittee index
	SyncCommitteeSubnetID(subCommitteeID uint64) (uint64, error)
	// GetSyncCommitteeContribution returns
	GetSyncCommitteeContribution(slot phase0.Slot, subnetID uint64, pubKey phase0.BLSPubKey) (*altair.SyncCommitteeContribution, error)
	// SubmitSignedContributionAndProof broadcasts to the network
	SubmitSignedContributionAndProof(contribution *altair.SignedContributionAndProof) error
}

type DomainCalls interface {
	DomainData(epoch phase0.Epoch, domain phase0.DomainType) (phase0.Domain, error)
}

type BeaconNode interface {
	// GetBeaconNetwork returns the beacon network the node is on
	GetBeaconNetwork() types.BeaconNetwork
	AttesterCalls
	ProposerCalls
	AggregatorCalls
	SyncCommitteeCalls
	SyncCommitteeContributionCalls
	DomainCalls
}
