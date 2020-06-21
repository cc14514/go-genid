/*************************************************************************
 * Copyright (C) 2016-2019 PDX Technologies, Inc. All Rights Reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * @Time   : 2020/6/16 10:58 上午
 * @Author : liangc
 *************************************************************************/

package genid

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"math/big"
	"sync"
	"time"
)

var (
	head, seed, last []byte

	lock sync.Mutex
)

func init() {
	new(sync.Once).Do(func() {
		key, _, _, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			panic(err)
		}
		t := time.Now()
		h, _ := new(big.Int).SetString(t.Format("20060102"), 10)
		random := new(big.Int).Mul(new(big.Int).SetBytes(key), big.NewInt(t.UnixNano())).Bytes()
		// head = yyyymmdd
		head = h.Bytes()
		// seed = append(head , random)
		seed = append(head, random...)
		last = new(big.Int).SetBytes(append(head, seed...)).Bytes()
		//log.Println("genid-inited", len(last), last)
	})
}

type ID []byte

func (id ID) String() string {
	return new(big.Int).SetBytes(id).String()
}

func (id ID) Hex() string {
	return hex.EncodeToString(id)
}

func (id ID) Hash() string {
	s1 := sha1.New()
	s1.Write(id)
	h := s1.Sum(nil)
	return hex.EncodeToString(h)
}

func GenID() ID {
	lock.Lock()
	defer lock.Unlock()
	id := new(big.Int).Add(new(big.Int).SetBytes(last), big.NewInt(1))
	last = id.Bytes()
	// TODO When last mod N equal 0 then reflash head
	return id.Bytes()
}
