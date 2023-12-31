// Copyright 2019 storyicon@foxmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import "testing"

func TestContains(t *testing.T) {
    type args struct {
        arr []string
        s   string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name: "test0",
            args: args{
                arr: []string{"a", "b"},
                s:   "b",
            },
            want: true,
        },
        {
            name: "test1",
            args: args{
                arr: []string{"a"},
                s:   "b",
            },
            want: false,
        },
    }
    for _, tt := range tests {
        if got := Contains(tt.args.arr, tt.args.s); got != tt.want {
            t.Errorf("%q. Contains() = %v, want %v", tt.name, got, tt.want)
        }
    }
}
