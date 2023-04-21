package server

import (
	"github.com/vishnushankarsg/metachain/chain"
	"github.com/vishnushankarsg/metachain/consensus"
	consensusDev "github.com/vishnushankarsg/metachain/consensus/dev"
	consensusDummy "github.com/vishnushankarsg/metachain/consensus/dummy"
	consensusIBFT "github.com/vishnushankarsg/metachain/consensus/ibft"
	consensusPolyBFT "github.com/vishnushankarsg/metachain/consensus/polybft"
	"github.com/vishnushankarsg/metachain/secrets"
	"github.com/vishnushankarsg/metachain/secrets/awsssm"
	"github.com/vishnushankarsg/metachain/secrets/gcpssm"
	"github.com/vishnushankarsg/metachain/secrets/hashicorpvault"
	"github.com/vishnushankarsg/metachain/secrets/local"
	"github.com/vishnushankarsg/metachain/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = "polybft"
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
