package main

import (
	"fmt"
	"math/big"
)

func findAliquotSum(bigA *big.Int)(*big.Int){
	if bigA.Cmp(big.NewInt(1))==0{
		return new(big.Int)
	}
	sqrt := &big.Int{}

	sqrt.Sqrt(bigA)
	factorSum := big.NewInt(1)
	
	var q, r big.Int 
	for i:= big.NewInt(2); i.Cmp(sqrt)<=0 ; i.Add(big.NewInt(1), i) {
		q, r := q.QuoRem(bigA, i, &r)
		if(r.Cmp(new(big.Int))==0){
			factorSum.Add(i, factorSum)
			if(i.Cmp(q)!=0){
				factorSum.Add(q, factorSum)
			}
		}
	}
	return factorSum
	
}

func findAliquotSequence(a int64){
	n := 1;
	for bigA:=big.NewInt(a);  bigA.Cmp(new(big.Int))!=0; bigA = findAliquotSum(bigA){
		fmt.Printf("The %d aliquot number in this sequence is: %d \n" , n , bigA)
		n++;
	}
}

func main(){
	findAliquotSequence(138)
}
