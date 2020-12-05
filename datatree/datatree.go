package datatree

import (
	"fmt"
	"strings"
)

//Datatree a structure for manage data
type Datatree struct {
	child   map[string]*Datatree
	father  *Datatree
	watcher int
}

// NewZnode return root of a datatree
func NewZnode() *Datatree {
	var newznode Datatree
	newznode.father = nil
	newznode.child = make(map[string]*Datatree)
	newznode.watcher = -1
	return &newznode
}

//SplitDir split dir into name
func SplitDir(dir string) []string {
	dirqueue := strings.Split(dir, "/")
	for i := 0; i < len(dirqueue); i++ {
		dirqueue[i] = strings.Replace(dirqueue[i], "/", "", -1)
	}
	return dirqueue
}

//CreateZnode create a new znode on datatree
func CreateZnode(dir string, root *Datatree) {
	dirq := SplitDir(dir)
	now := root
	for i := 0; i < len(dirq)-1; i++ {
		next, exist := (*now).child[dirq[i]]
		if exist {
			now = next
		} else {
			fmt.Println("dir not exist")
			return
		}
	}
	next, exist := (*now).child[dirq[len(dirq)-1]]
	if exist {
		fmt.Println("znode already exist")
		now = next //meaningless, just to use next
	} else {
		newz := NewZnode()
		newz.father = now
		(*now).child[dirq[len(dirq)-1]] = newz
		fmt.Println("create " + dir + " successfully")
	}
}

//DeleteZnode delete a new znode on datatree
func DeleteZnode(dir string, root *Datatree) {
	dirq := SplitDir(dir)
	now := root
	for i := 0; i < len(dirq); i++ {
		next, exist := (*now).child[dirq[i]]
		if exist {
			now = next
		} else {
			fmt.Println("dir not exist")
			return
		}
	}
	(*now).father = nil
	now = nil

}

//LookZnode shows all child of a direcory
func LookZnode(dir string, root *Datatree) {
	/*dirq := SplitDir(dir)
	now := root
	for i := 0; i < len(dirq); i++ {
		next, exist := (*now).child[dirq[i]]
		if exist {
			now = next
		} else {
			fmt.Println("dir not exist")
			return
		}
	}
	(*now).father = nil
	now = nil*/
}

//CreateWatcher create a new znode on datatree
func CreateWatcher(dir string, root *Datatree, clientid int) {
	dirq := SplitDir(dir)
	now := root
	for i := 0; i < len(dirq); i++ {
		next, exist := (*now).child[dirq[i]]
		if exist {
			now = next
		} else {
			fmt.Println("dir not exist")
			return
		}
	}
	(*now).watcher = clientid
}
