package router

type JHSHashShard struct {
	NodeShardNum  int
	TableShardNum int
}

type JHSShard interface {
	JHSFindForKey(key interface{}) (int, int)
	FindForKey(key interface{}) int
}

func (s *JHSHashShard) JHSFindForKey(key interface{}) (int, int) {
	h := HashValue(key)
	t_index := -1
	if s.TableShardNum > 0 {
		t_index = int(h % uint64(s.TableShardNum))
	}

	return t_index, int(h % uint64(s.NodeShardNum))
}

func (s *JHSHashShard) FindForKey(key interface{}) int {
	h := HashValue(key)

	return int(h % uint64(s.NodeShardNum))
}
