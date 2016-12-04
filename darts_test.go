package darts

import (
	"reflect"
	"testing"
)

func TestDarts_BuildFromStrs(t *testing.T) {
	type args struct {
		words [][]rune
	}
	tests := []struct {
		name    string
		d       *Darts
		args    args
		wantErr bool
	}{
		{
			"正常纯中文构建-1", &Darts{},
			args{[][]rune{[]rune("你好"), []rune("你很好"), []rune("你真的很好"), []rune("我不好"), []rune{}}}, false,
		},
		{
			"无效失败构建", &Darts{}, args{[][]rune{}}, true,
		},
		{
			"中英文混杂构建", &Darts{}, args{[][]rune{[]rune("Hey man!"), []rune("Hey 兄弟!"), []rune("hey guys!")}}, false,
		},
	}
	for _, tt := range tests {
		if err := tt.d.BuildFromStrs(tt.args.words); (err != nil) != tt.wantErr {
			t.Errorf("%q. Darts.BuildFromStrs() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestDarts_ExactMatchSearch(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		d    *Darts
		data [][]rune
		args args
		want bool
	}{
		{
			"中文单词匹配成功-1", &Darts{},
			[][]rune{[]rune("你好"), []rune("你好吗"), []rune("你不好"), []rune("死神"), []rune("工程党"), []rune("呼呼"), []rune("工程文化")},
			args{"你好"}, true,
		},
		{
			"中英文单词匹配成功-1", &Darts{},
			[][]rune{[]rune("hey你好"), []rune("Hey你好吗"), []rune("Oh,你不好"), []rune("Ha 死神"), []rune("工程党,yes"), []rune("god,呼呼"), []rune("工程文化")},
			args{"Ha 死神"}, true,
		},
	}

	for _, tt := range tests {
		if err := tt.d.BuildFromStrs(tt.data); err != nil {
			t.Errorf("%q. Darts.BuildFromStrs() fail", tt.name)
		}

		if got := tt.d.ExactMatchSearch(tt.args.key); got != tt.want {
			t.Errorf("%q. Darts.ExactMatchSearch() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestDarts_CommonPrefixSearch(t *testing.T) {
	type args struct {
		size int64
		key  string
	}
	tests := []struct {
		name string
		d    *Darts
		data [][]rune
		args args
		want []string
	}{
		{
			"中文单词匹配成功-1", &Darts{},
			[][]rune{[]rune("你好"), []rune("你好吗"), []rune("你不好"), []rune("死神"), []rune("工程党"), []rune("呼呼"), []rune("工程文化")},
			args{2, "你"},
			[]string{"你好", "你好吗"},
		},

		{
			"中英文单词匹配成功-1", &Darts{},
			[][]rune{[]rune("a你好"), []rune("a你好吗"), []rune("你不好"), []rune("a死神"), []rune("工程党"), []rune("呼呼"), []rune("工程文化")},
			args{-1, "a"},
			[]string{"a你好", "a你好吗", "a死神"},
		},

		{
			"无匹配-1", &Darts{},
			[][]rune{[]rune("a你好"), []rune("a你好吗"), []rune("你不好"), []rune("a死神"), []rune("工程党"), []rune("呼呼"), []rune("工程文化")},
			args{-1, "b"},
			nil,
		},
	}
	for _, tt := range tests {
		if err := tt.d.BuildFromStrs(tt.data); err != nil {
			t.Errorf("%q. Darts.BuildFromStrs() fail", tt.name)
		}

		got := tt.d.CommonPrefixSearch(tt.args.size, tt.args.key)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Darts.CommonPrefixSearch() got = %v, want %v", tt.name, got, tt.want)
		}
	}
}
