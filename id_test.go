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
 * @Time   : 2020/6/16 11:01 上午
 * @Author : liangc
 *************************************************************************/

package genid

import "testing"

func TestID(t *testing.T) {
	id := GenID()
	t.Log(len(id), id)
	t.Log(id.Hex())
	t.Log(id.Hash())
}

func BenchmarkIDHex(b *testing.B) {
	m := make(map[string]interface{}, 1000000)
	for i := 0; i < b.N; i++ {
		if _, ok := m[GenID().Hex()]; ok {
			panic("duplicate")
		}
	}
}

func BenchmarkIDHash(b *testing.B) {
	m := make(map[string]interface{}, 1000000)
	for i := 0; i < b.N; i++ {
		if _, ok := m[GenID().Hash()]; ok {
			panic("duplicate")
		}
	}
}
