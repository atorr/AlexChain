package blockchain

import(
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type(
	url = url.URL
	txn struct {
		sender string
		recipient string
		amount int
	}
	block struct {
		index int
		timestamp int
		txns []txn
		proof int
		prevHash string
	}
	chain []*block
	blockChainOps interface{
		func newBlock(proof int, prevHash string) block
		func newTxn(sender, recipient string, amount int) int
		func registerNode(address string)
		func validChain(chn chain) bool
		func resolveConflicts() bool
		func lastBlock() *block
	}
	blockChain struct{
		chn chain
		currentTxns []txn
		nodes []*url
	}
)

func (b *blockChain) newBlock(proof int, prevHash string) block {
	blk := &block{
		index: len(b.chain) + 1,
		timestamp: time.Now(),
		txns: b.currentTxns,
		proof: proof,
		prevHash: prevHash,
	}

	b.currentTxns = []txn{}

	b.chn = append(b.chn, blk)
	return blk
}

func (b *blockChain) newTxn(sender, recipient string, amount int) int {
	t := txn{
		sender: sender,
		recipient: recipient,
		amount: amount,
	}

	b.currentTxns = append(b.currentTxns, t)
	return b.lastBlock()['index'] + 1
}

func (b *blockChain) registerNode(address string) {
	parsed_url := url.Parse(address)
	b.nodes = append(b.nodes, parsed_url)
}

func (b *blockChain) validChain(chn chain) bool {
	prevBlock := b.chain[0]
	currentIndex := 1

	for currentIndex < len(b.chain) {
		blk = b.chain[currentIndex]

		if block.prevHash != hash(prevBlock) {
			return false
		}

		if !valid_proof(prevBlock.proof, blk.proof) {
			return false
		}

		prevBlock = blk
		currentIndex = currentIndex + 1
	}

	return true
}

func (b *blockChain) resolveConflicts() bool {
	var newChain chain
	neighborinos := b.nodes

	maxLength := len(b.chn)

	for _, node := range neighborinos {
		addr := fmt.Printf("http://%s/chain", node.String())
		response, err := http.Get(addr)
		if err != nil {
			return false
		}

		if response.Status != url.StatusOK {
			return false
		}

		length := response.Body.Length
		chn := response.Body.Chn

		if length > maxLength && b.validChain(chn) {
			maxLength = length
			newChain = chn
		}
	}

	if newChain {
		b.chn = newChain
		return true
	}

	return false
}
