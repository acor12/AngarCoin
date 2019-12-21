package node

import (
	bolt "go.etcd.io/bbolt"
)

var (
	bucketName = []byte("blocks")
	tip        = []byte("tip")
)

// Blockchain struct representation
type Blockchain struct {
	Tip []byte
	db  *bolt.DB
}

//Close database conection
func (bc *Blockchain) Close() {
	bc.db.Close()
}

//GetBlock returns block by hash
func (bc *Blockchain) GetBlock(hash []byte) (block *Block) {
	// TODO
	bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)

		bucket.Get(hash)
		return nil
	})
	return
}

//Iterator create a blockchain iterator instance
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.Tip, bc}
}

//AddBlock appends block to the blockchain
func (bc *Blockchain) AddBlock(block Block) {
	bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("l"), block.Hash)
		return nil
	})
}

// InitializeBlockchain loads or create blockchain file
// and returns a blockchain instance
func InitializeBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open("blockchain", 0600, nil)
	if err != nil {
		panic("Opening database error")
	}

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			bucket, err = tx.CreateBucket(bucketName)
			err = bucket.Put([]byte("blockhash"), []byte("serialized block"))
			err = bucket.Put([]byte("l"), []byte("last hash"))
			tip = []byte("l")
			return err
		}

		tip = bucket.Get([]byte("l"))
		return nil
	})
	return &Blockchain{
		Tip: tip,
		db:  db,
	}
}

//BlockchainIterator struct
type BlockchainIterator struct {
	nextHash   []byte
	blockchain *Blockchain
}

// Next returns the next block in the blockchain
func (bci *BlockchainIterator) Next() (block *Block) {
	block = bci.blockchain.GetBlock(bci.nextHash)
	bci.nextHash = block.PrevHash
	return block
}
