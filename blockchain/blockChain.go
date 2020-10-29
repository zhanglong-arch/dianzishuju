package blockchain

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"math/big"
)

//桶的名称，该桶用于装区块信息
var BUCKET_NAME = "blocks"
//表示最新的区块的key名
var LAST_KEY = "lasthash"
//存储区块数据的文件
var CHAINDB = "chain.db"

var CHAIN BlockChain

/**
 * 区块链结构体实例定义:用于表示代表一条区块链
 * 该区块链包含以下功能:
 		① 将新产生的区块与已有的区块链接起来，并保存
		② 可以查询某个区块的信息
		③ 可以将所有区块进行遍历，输出区块信息
*/
type BlockChain struct {
	LastHash []byte //最新区块hash
	BoltDb   *bolt.DB
}
/**
 * 查询所有的区块信息，并返回。将所有的区块放入到切片中
 */
func (bc BlockChain) QueryAllBlocks() []*Block{
	blocks := make([]*Block,0)
	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据出错")
		}
		eachKey := bc.LastHash
		preHashBig := new(big.Int)
		zeroBig := big.NewInt(0)
		for{
			eachBlockBytes := bucket.Get(eachKey)
			//反序列化以后得到的每一个区块
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//将遍历到每一个区块结构体指针放入到[]byte容器中
			blocks = append(blocks,eachBlock)

			//preHashBig.SetBytes(eachBlock.PrevHash)
			if preHashBig.Cmp(zeroBig) == 0{//判断区块链遍历是否已到创世区块，如果到创世区块，跳出循环，否则继续
				break
			}

			eachKey = eachBlock.PrevHash
		}

		return nil
	})
	return blocks
}


/**
 * 通过区块的高度查询某个具体的区块，返回区块实例
 */
func (bc BlockChain) QueryBlockByHeight(height int64) *Block{
	if height < 0 {//如果目标高度小于0，则说明参数不合法
		return nil
	}
	var block *Block
	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据失败")
		}
		hashKey := bc.LastHash
		for{
			lasstBlockBytes := bucket.Get(hashKey)
			eachBlock, _ := DeSerialize(lasstBlockBytes)
			if eachBlock.Height < height{//给定的数字超出区块链中的区块高度，直接返回
				break
			}
			if eachBlock.Height == height{//已找到目标区块
				block = eachBlock
				break
			}
			//遍历的当前的区块的高度与目标高度不一致，继续往前遍历
			//以eachBlock.PrevHash为key，使用Get获取上一个区块的数据
			hashKey = eachBlock.PrevHash
		}

		return nil
	})
	return block
}

/**
 * 用于创建一条区块链，并返回该区块链实例
	解释：由于区块链就是由一个一个的区块组成的，因此,如果要创建一条区块链，那么必须要先
		创建一个区块，该区块作为该条区块链的创世区块。
*/
func NewBlockChain() BlockChain {
	//0、打开存储区块数据的chain.db文件
	db, err := bolt.Open(CHAINDB, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	var bl BlockChain
	//先从区块链中都看是否创世区块已经存在
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {//程序第一次执行时，bucket为空
			//创建新桶
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		if len(lastHash) == 0 { //没有创世区块
			//1、创建创世区块
			genesis := CreateGenesisBlock() //创世区块
			//2、创建一个存储区块数据的文件
			fmt.Printf("genesis的Hash值:%x\n", genesis.Hash)
			bl = BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
			genesisBytes, _ := genesis.Serialize()
			bucket.Put(genesis.Hash, genesisBytes)
			bucket.Put([]byte(LAST_KEY), genesis.Hash)
		} else { //有创世区块
			lastHash := bucket.Get([]byte(LAST_KEY))
			lastBlockBytes := bucket.Get(lastHash) //创世区块的[]byte
			lastBlock, err := DeSerialize(lastBlockBytes)
			if err != nil {
				panic("读取区块链数据失败")
			}
			bl = BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}
		}
		return nil
	})
	CHAIN = bl
	return bl
}

/**
 * 调用BlockChain的该SaveBlock方法，该方法可以将一个生成的新区块保存到chain.db文件中
 */
func (bc BlockChain) SaveData(data []byte) (Block, error) {
	db := bc.BoltDb
	var e error
	var lastBlock *Block
	//先查询chain.db中存储的最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			e = errors.New("boltdb未创建，请重试!")
			return e
		}
		//lastHash := bucket.Get([]byte(LAST_KEY))
		lastBlockBytes := bucket.Get(bc.LastHash)
		lastBlock, _ = DeSerialize(lastBlockBytes)
		return nil
	})

	//1、先生成一个区块 把data存入到新生成的区块中
	newBlock := NewBlock(lastBlock.Height+1, data, lastBlock.Hash)

	//更新chain.db 把newblock存入到boltdb中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//key=hash, value=block的byte
		//区块序列化
		newBlockBytes, _ := newBlock.Serialize()
		//把区块信息保存到boltdb中
		bucket.Put(newBlock.Hash, newBlockBytes)
		//更新代表最后一个区块hash值的记录
		bucket.Put([]byte(LAST_KEY), newBlock.Hash)
		bc.LastHash = newBlock.Hash
		return nil
	})
	return newBlock, e
}