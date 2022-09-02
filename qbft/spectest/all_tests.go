package spectest

import (
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/commit"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/controller"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/decided"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/messages"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/prepare"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/proposal"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/proposer"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/roundchange"
	"testing"
)

type SpecTest interface {
	TestName() string
	Run(t *testing.T)
}

var AllTests = []SpecTest{
	controller.StartInstanceFirstHeight(),
	controller.StartInstancePreviousDecided(),
	controller.StartInstancePreviousNotDecided(),
	controller.StartInstanceInvalidValue(),
	controller.FirstDecided(),
	controller.InvalidIdentifier(),
	controller.NoInstanceRunning(),
	controller.NotFirstDecided(),
	controller.NotDecided(),
	controller.ProcessMsgError(),
	controller.SavedAndBroadcastedDecided(),

	proposer.FourOperators(),
	proposer.SevenOperators(),
	proposer.TenOperators(),
	proposer.ThirteenOperators(),

	messages.RoundChangeDataInvalidJustifications(),
	messages.RoundChangeDataInvalidPreparedRound(),
	messages.RoundChangeDataInvalidPreparedValue(),
	messages.RoundChangePrePreparedJustifications(),
	messages.RoundChangeNotPreparedJustifications(),
	messages.CommitDataEncoding(),
	messages.MsgNilIdentifier(),
	messages.MsgNonZeroIdentifier(),
	messages.MsgTypeUnknown(),
	messages.PrepareDataEncoding(),
	messages.ProposeDataEncoding(),
	messages.MsgDataNil(),
	messages.MsgDataNonZero(),
	messages.SignedMsgSigTooShort(),
	messages.SignedMsgSigTooLong(),
	messages.SignedMsgNoSigners(),
	messages.SignedMsgDuplicateSigners(),
	messages.SignedMsgMultiSigners(),
	messages.GetRoot(),
	messages.SignedMessageEncoding(),
	messages.CreateProposal(),
	messages.CreateProposalPreviouslyPrepared(),
	messages.CreateProposalNotPreviouslyPrepared(),
	messages.CreatePrepare(),
	messages.CreateCommit(),
	messages.CreateRoundChange(),
	messages.CreateRoundChangePreviouslyPrepared(),
	messages.RoundChangeDataEncoding(),
	messages.PrepareDataInvalid(),
	messages.CommitDataInvalid(),
	messages.ProposalDataInvalid(),

	tests.HappyFlow(),
	tests.SevenOperators(),
	tests.TenOperators(),
	tests.ThirteenOperators(),

	proposal.HappyFlow(),
	proposal.NotPreparedPreviouslyJustification(),
	proposal.PreparedPreviouslyJustification(),
	proposal.DifferentJustifications(),
	proposal.JustificationsNotHeighest(),
	proposal.JustificationsValueNotJustified(),
	proposal.DuplicateMsg(),
	proposal.FirstRoundJustification(),
	proposal.FutureRoundAcceptedProposalNotPrevPrepared(),
	proposal.FutureRoundAcceptedProposal(),
	proposal.PastRoundValidProposalPrevNotPrepared(),
	proposal.PastRoundValidProposal(),
	proposal.PastRound(),
	proposal.ImparsableProposalData(),
	proposal.InvalidRoundChangeJustificationPrepared(),
	proposal.InvalidRoundChangeJustification(),
	proposal.PreparedPreviouslyNoRCJustificationQuorum(),
	proposal.NoRCJustification(),
	proposal.PreparedPreviouslyNoPrepareJustificationQuorum(),
	proposal.PreparedPreviouslyDuplicatePrepareMsg(),
	proposal.PreparedPreviouslyDuplicatePrepareQuorum(),
	proposal.PreparedPreviouslyDuplicateRCMsg(),
	proposal.PreparedPreviouslyDuplicateRCQuorum(),
	proposal.DuplicateRCMsg(),
	proposal.InvalidPrepareJustificationValue(),
	proposal.InvalidPrepareJustificationRound(),
	proposal.InvalidProposalData(),
	proposal.InvalidValueCheck(),
	proposal.MultiSigner(),
	proposal.PostDecided(),
	proposal.PostPrepared(),
	proposal.SecondProposalForRound(),
	proposal.WrongHeight(),
	proposal.WrongProposer(),
	proposal.WrongSignature(),
	proposal.UnknownSigner(),

	prepare.DuplicateMsg(),
	prepare.HappyFlow(),
	prepare.ImparsableProposalData(),
	prepare.InvalidPrepareData(),
	prepare.MultiSigner(),
	prepare.NoPreviousProposal(),
	prepare.OldRound(),
	prepare.FutureRound(),
	prepare.PostDecided(),
	prepare.WrongData(),
	prepare.WrongHeight(),
	prepare.WrongSignature(),
	prepare.UnknownSigner(),

	commit.CurrentRound(),
	commit.FutureRound(),
	commit.PastRound(),
	commit.DuplicateMsg(),
	commit.HappyFlow(),
	commit.InvalidCommitData(),
	commit.PostDecided(),
	commit.WrongData1(),
	commit.WrongData2(),
	commit.MultiSignerWithOverlap(),
	commit.MultiSignerNoOverlap(),
	commit.DuplicateSigners(),
	commit.NoPrevAcceptedProposal(),
	commit.WrongHeight(),
	commit.ImparsableCommitData(),
	commit.WrongSignature(),
	commit.UnknownSigner(),
	commit.InvalidValCheck(),

	decided.UnknownSigner(),
	decided.WrongSignature(),
	decided.WrongHeight(),
	decided.PostDecided(),
	decided.SecondMsg(),
	decided.PastRound(),
	decided.NoPrevAcceptedProposal(),
	decided.InvalidValCheckData(),
	decided.InvalidData(),
	decided.ImparsableData(),
	decided.FutureRound(),
	decided.DuplicateSigners(),
	decided.DuplicateMsg(),
	decided.PrevCommitOverlap(),
	decided.CurrentRound(),

	roundchange.HappyFlow(),
	roundchange.WrongHeight(),
	roundchange.WrongSig(),
	roundchange.UnknownSigner(),
	roundchange.MultiSigner(),
	roundchange.QuorumNotPrepared(),
	roundchange.QuorumPrepared(),
	roundchange.PeerPrepared(),
	roundchange.PeerPreparedDifferentHeights(),
	roundchange.JustificationWrongValue(),
	roundchange.JustificationWrongRound(),
	roundchange.JustificationNoQuorum(),
	roundchange.JustificationMultiSigners(),
	roundchange.JustificationInvalidSig(),
	roundchange.JustificationInvalidRound(),
	roundchange.JustificationInvalidPrepareData(),
	roundchange.JustificationDuplicateMsg(),
	roundchange.InvalidRoundChangeData(),
	roundchange.F1DifferentFutureRounds(),
	roundchange.F1DifferentFutureRoundsNotPrepared(),
	roundchange.PastRound(),
	roundchange.DuplicateMsgQuorum(),
	roundchange.DuplicateMsgQuorumPreparedRCFirst(),
	roundchange.DuplicateMsg(),
	roundchange.ImparsableRoundChangeData(),
	roundchange.NotProposer(),
	roundchange.ValidJustification(),
	roundchange.F1DuplicateSigner(),
	roundchange.F1DuplicateSignerNotPrepared(),
	roundchange.F1Speedup(),
	roundchange.F1SpeedupPrevPrepared(),
	roundchange.AfterProposal(),
}
