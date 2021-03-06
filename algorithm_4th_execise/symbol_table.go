package practice

type symbolTable struct {
	keys   []interface{}
	values []interface{}
	N      int
}

func (st *symbolTable) put(key interface{}, value interface{}) {
	if key == nil || value == nil {
		return
	}
	index := st.rank(key)
	if index < st.N && st.keys[index] == key {
		st.values[index] = value
		return
	}
	st.N++
	if st.N > len(st.keys) {
		st.keys = append(st.keys, make([]interface{}, 10)...)
		st.values = append(st.values, make([]interface{}, 10)...)
	}
	for i := st.N; i > index; i-- {
		st.keys[i] = st.keys[i-1]
		st.values[i] = st.values[i-1]
	}
	st.keys[index] = key
	st.values[index] = value
}

func (st *symbolTable) get(key interface{}) (value interface{}) {
	if key == nil {
		return nil
	}
	index := st.rankV2(key, 0, st.N-1)
	if index < st.N && st.keys[index] == key {
		return st.values[index]
	}
	return nil
}

func (st *symbolTable) delete(key interface{}) {
	if key == nil {
		return
	}
	index := st.rank(key)
	if index < st.N && st.keys[index] == key {
	} else {
		return
	}
	for i := index; i < st.N-1; i++ {
		st.keys[i] = st.keys[i+1]
		st.values[i] = st.values[i+1]
	}
	st.N--
}

func (st *symbolTable) getKeys() (ks []interface{}) {
	return st.keys
}

// Binary search
func (st *symbolTable) rank(key interface{}) (index int) {
	low, high := 0, st.N-1

	for low <= high {
		mid := low + (high-low)/2
		v := st.keys[mid]
		r := Compare(key, v)
		if r == -1 {
			high = mid - 1
		} else if r == 1 {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}

// recursion binary search
func (st *symbolTable) rankV2(key interface{}, low, high int) (index int) {
	if high < low {
		return low
	}
	mid := low + (high-low)/2
	v := st.keys[mid]
	r := Compare(key, v)
	if r == -1 {
		return st.rankV2(key, low, mid-1)
	} else if r == 1 {
		return st.rankV2(key, mid+1, high)
	} else {
		return mid
	}
}
