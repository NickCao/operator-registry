package cache

import (
	"github.com/tidwall/btree"
)

type BundleKey struct {
	PackageName string
	ChannelName string
	Name        string
}

func bundleKeyComparator(a, b BundleKey) bool {
	if a.ChannelName != b.ChannelName {
		return a.ChannelName < b.ChannelName
	}
	if a.PackageName != b.PackageName {
		return a.PackageName < b.PackageName
	}
	return a.Name < b.Name
}

type bundleKeys struct {
	t *btree.BTreeG[BundleKey]
}

func newBundleKeys() bundleKeys {
	return bundleKeys{btree.NewBTreeG[BundleKey](bundleKeyComparator)}
}

func (b bundleKeys) Set(k BundleKey) {
	b.t.Set(k)
}

func (b bundleKeys) Len() int {
	return b.t.Len()
}

func (b bundleKeys) Walk(f func(k BundleKey) error) error {
	it := b.t.Iter()
	for it.Next() {
		if err := f(it.Item()); err != nil {
			return err
		}
	}
	return nil
}
