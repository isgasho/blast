// Copyright (c) 2019 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package manager

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/objx"

	"github.com/mosuka/blast/protobuf/raft"
)

func TestRaftFSM_GetNode(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	node1 := &raft.Node{
		Id: "node1",
		Metadata: &raft.Metadata{
			BindAddr: ":16060",
			GrpcAddr: ":17070",
			HttpAddr: ":18080",
			Leader:   false,
		},
	}
	node2 := &raft.Node{
		Id: "node2",
		Metadata: &raft.Metadata{
			BindAddr: ":16061",
			GrpcAddr: ":17071",
			HttpAddr: ":18081",
			Leader:   false,
		},
	}
	node3 := &raft.Node{
		Id: "node3",
		Metadata: &raft.Metadata{
			BindAddr: ":16062",
			GrpcAddr: ":17072",
			HttpAddr: ":18082",
			Leader:   false,
		},
	}

	fsm.applySetNode(node1)
	fsm.applySetNode(node2)
	fsm.applySetNode(node3)

	node, err := fsm.GetNode(&raft.Node{
		Id: "node2",
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedValue := &raft.Node{
		Id: "node2",
		Metadata: &raft.Metadata{
			BindAddr: ":16061",
			GrpcAddr: ":17071",
			HttpAddr: ":18081",
			Leader:   false,
		},
	}
	actualValue := node
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}

}

func TestRaftFSM_SetNode(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	node1 := &raft.Node{
		Id: "node1",
		Metadata: &raft.Metadata{
			BindAddr: ":16060",
			GrpcAddr: ":17070",
			HttpAddr: ":18080",
			Leader:   false,
		},
	}
	node2 := &raft.Node{
		Id: "node2",
		Metadata: &raft.Metadata{
			BindAddr: ":16061",
			GrpcAddr: ":17071",
			HttpAddr: ":18081",
			Leader:   false,
		},
	}
	node3 := &raft.Node{
		Id: "node3",
		Metadata: &raft.Metadata{
			BindAddr: ":16062",
			GrpcAddr: ":17072",
			HttpAddr: ":18082",
			Leader:   false,
		},
	}

	fsm.applySetNode(node1)
	fsm.applySetNode(node2)
	fsm.applySetNode(node3)

	node, err := fsm.GetNode(&raft.Node{
		Id: "node2",
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedValue := &raft.Node{
		Id: "node2",
		Metadata: &raft.Metadata{
			BindAddr: ":16061",
			GrpcAddr: ":17071",
			HttpAddr: ":18081",
			Leader:   false,
		},
	}
	actualValue := node
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}

	node2.Metadata.Leader = true
	fsm.applySetNode(node2)

	node, err = fsm.GetNode(&raft.Node{
		Id: "node2",
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedValue = &raft.Node{
		Id: "node2",
		Metadata: &raft.Metadata{
			BindAddr: ":16061",
			GrpcAddr: ":17071",
			HttpAddr: ":18081",
			Leader:   true,
		},
	}
	actualValue = node
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}
}

func TestRaftFSM_DeleteNode(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	node1 := &raft.Node{
		Id: "node1",
		Metadata: &raft.Metadata{
			BindAddr: ":16060",
			GrpcAddr: ":17070",
			HttpAddr: ":18080",
			Leader:   false,
		},
	}
	node2 := &raft.Node{
		Id: "node2",
		Metadata: &raft.Metadata{
			BindAddr: ":16061",
			GrpcAddr: ":17071",
			HttpAddr: ":18081",
			Leader:   false,
		},
	}
	node3 := &raft.Node{
		Id: "node3",
		Metadata: &raft.Metadata{
			BindAddr: ":16062",
			GrpcAddr: ":17072",
			HttpAddr: ":18082",
			Leader:   false,
		},
	}

	fsm.applySetNode(node1)
	fsm.applySetNode(node2)
	fsm.applySetNode(node3)

	node, err := fsm.GetNode(&raft.Node{
		Id: "node2",
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedValue := &raft.Node{
		Id: "node2",
		Metadata: &raft.Metadata{
			BindAddr: ":16061",
			GrpcAddr: ":17071",
			HttpAddr: ":18081",
			Leader:   false,
		},
	}
	actualValue := node
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}

	node2.Metadata.Leader = true
	fsm.applySetNode(node2)

	node, err = fsm.GetNode(&raft.Node{
		Id: "node2",
	})
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedValue = &raft.Node{
		Id: "node2",
		Metadata: &raft.Metadata{
			BindAddr: ":16061",
			GrpcAddr: ":17071",
			HttpAddr: ":18081",
			Leader:   true,
		},
	}
	actualValue = node
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}

	fsm.applyDeleteNode(&raft.Node{
		Id: "node2",
	})

	node, err = fsm.GetNode(&raft.Node{
		Id: "node2",
	})
	if err == nil {
		t.Errorf("expected error: %v", err)
	}

	actualValue = node
	if reflect.DeepEqual(nil, actualValue) {
		t.Errorf("expected content to see nil, saw %v", actualValue)
	}
}

func TestRaftFSM_pathKeys(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	keys := fsm.pathKeys("/a/b/c/d")

	expectedValue := []string{"a", "b", "c", "d"}
	actualValue := keys
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}
}

func TestRaftFSM_makeSafePath(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	safePath := fsm.makeSafePath("/a/b/c/d")

	expectedValue := "a.b.c.d"
	actualValue := safePath
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}

}

func TestRaftFSM_walk(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	data := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"c": "abc",
				"d": "abd",
			},
			"e": []interface{}{
				"ae1",
				"ae2",
			},
		},
	}

	val1 := fsm.normalize(objx.New(data))

	exp1 := data
	act1 := val1
	if !reflect.DeepEqual(exp1, act1) {
		t.Errorf("expected content to see %v, saw %v", exp1, act1)
	}
}

func TestRaftFSM_Get(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	fsm.applySet("/", map[string]interface{}{"a": 1}, false)

	value, err := fsm.Get("/a")
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedValue := 1
	actualValue := value
	if expectedValue != actualValue {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}
}

func TestRaftFSM_makeMap(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	value := fsm.makeMap("/hoge/fuga", map[string]interface{}{"hoo": "var"})

	expectedValue := map[string]interface{}{"hoge": map[string]interface{}{"fuga": map[string]interface{}{"hoo": "var"}}}
	actualValue := value
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}
}

func TestRaftFSM_Set(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	fsm.applySet("/", map[string]interface{}{"a": 1}, true)

	val1, err := fsm.Get("/a")
	if err != nil {
		t.Errorf("%v", err)
	}

	exp1 := 1
	act1 := val1
	if exp1 != act1 {
		t.Errorf("expected content to see %v, saw %v", exp1, act1)
	}

	fsm.applySet("/b/bb", map[string]interface{}{"b": 1}, false)

	val2, err := fsm.Get("/b")
	if err != nil {
		t.Errorf("%v", err)
	}

	exp2 := map[string]interface{}{"bb": map[string]interface{}{"b": 1}}
	act2 := val2.(map[string]interface{})
	if !reflect.DeepEqual(exp2, act2) {
		t.Errorf("expected content to see %v, saw %v", exp2, act2)
	}

	fsm.applySet("/", map[string]interface{}{"a": 1}, false)

	val3, err := fsm.Get("/")
	if err != nil {
		t.Errorf("%v", err)
	}

	exp3 := map[string]interface{}{"a": 1}
	act3 := val3
	if !reflect.DeepEqual(exp3, act3) {
		t.Errorf("expected content to see %v, saw %v", exp3, act3)
	}

	fsm.applySet("/", map[string]interface{}{"b": 2}, true)

	val4, err := fsm.Get("/")
	if err != nil {
		t.Errorf("%v", err)
	}

	exp4 := map[string]interface{}{"a": 1, "b": 2}
	act4 := val4
	if !reflect.DeepEqual(exp4, act4) {
		t.Errorf("expected content to see %v, saw %v", exp4, act4)
	}
}

func TestRaftFSM_Delete(t *testing.T) {
	tmp, err := ioutil.TempDir("", "")
	if err != nil {
		t.Errorf("%v", err)
	}
	defer func() {
		err := os.RemoveAll(tmp)
		if err != nil {
			t.Errorf("%v", err)
		}
	}()

	logger := log.New(os.Stderr, "", 0)

	fsm, err := NewRaftFSM(tmp, logger)
	if err != nil {
		t.Errorf("%v", err)
	}

	fsm.applySet("/", map[string]interface{}{"a": 1}, false)

	value, err := fsm.Get("/a")
	if err != nil {
		t.Errorf("%v", err)
	}

	expectedValue := 1
	actualValue := value
	if expectedValue != actualValue {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}

	fsm.applyDelete("/a")

	value, err = fsm.Get("/a")
	if err == nil {
		t.Errorf("expected nil: %v", err)
	}

	actualValue = value
	if nil != actualValue {
		t.Errorf("expected content to see %v, saw %v", expectedValue, actualValue)
	}
}
