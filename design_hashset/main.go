package main

func main() {

}

type MyHashSet struct {
	mp map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		map[int]struct{}{},
	}
}

func (this *MyHashSet) Add(key int) {
	this.mp[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	if _, ok := this.mp[key]; ok {
		delete(this.mp, key)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if _, ok := this.mp[key]; ok {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
