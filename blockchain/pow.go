package blockchain

import (
	"BeegoDemo/util"
	"bytes"
	"crypto/sha256"
	"math/big"
)

const DIFFFICULTY = 16

/**
 * 工作量证明结构体
 */

type ProofOfWork struct {
	//目标值
	Target *big.Int
	//工作量证明算法对应的哪个区块
	Block Block
}

/**
 * 实例化一个pow算法实例
 */
func NewPow(block Block) ProofOfWork{
	target := big.NewInt(1)//初始值
	target.Lsh(target, 255-DIFFFICULTY)
	pow := ProofOfWork{
		Target: target,
		Block: block,
	}
	return pow
}

/**
 * pow算法：寻找符合条件的nonce值
 */
func (p ProofOfWork) Run() ([]byte,int64){
	var nonce int64
	//var bigBlock *big.Int
	bigBlock := new(big.Int)
	var block256Hash []byte
	for {
		block := p.Block

		heightBytes, _ := util.IntToBytes(block.Height)
		timeBytes, _ := util.IntToBytes(block.TimeStamp)
		versionBytes := util.StringToBytes(block.Version)
		nonceBytes, _ := util.IntToBytes(nonce)

		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		},[]byte{})
		sha256Hash := sha256.New()
		sha256Hash.Write(blockBytes)
		block256Hash = sha256Hash.Sum(nil)
		//fmt.Printf("挖矿中，当前尝试Nonce值：%d\n",nonce)

		//sha256Hash（区块+nonce） 对应的大整数
		bigBlock = bigBlock.SetBytes(block256Hash)
		//fmt.Printf("目标值：%x\n",p.Target)
		//fmt.Printf("hash值：%x\n",bigBlock)

		if p.Target.Cmp(bigBlock) == 1 {//如果满足条件时，退出循环
			break
		}
		nonce++ //如果条件不满足，nonce值+1， 继续下次循环
		//time.Sleep(100)
	}
	return block256Hash,nonce
}