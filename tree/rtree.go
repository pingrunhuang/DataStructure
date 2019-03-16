package DataStructure

type Coordinate struct {
	lat float64
	lon float64
}

func (this *Coordinate) isOnLeft(other *Coordinate) bool {
	if this.lon < other.lon {
		return true
	}
	return false
}

func (this *Coordinate) isOnRight(other *Coordinate) bool {
	return !this.isOnLeft(other)
}

func (this *Coordinate) isLower(other *Coordinate) bool {
	if this.lat < other.lat {
		return true
	}
	return false
}

func (this *Coordinate) isUpper(other *Coordinate) bool {
	return !this.isLower(other)
}

type Point struct {
	name     string
	location Coordinate
}

type MinimumBoundingBox struct {
	id         int64
	lowerLeft  Coordinate
	upperRight Coordinate
	points     *[]Point
}

func (this *MinimumBoundingBox) contains(other *MinimumBoundingBox) bool {
	if this.lowerLeft.isOnLeft(&other.lowerLeft) && this.lowerLeft.isLower(&other.lowerLeft) &&
		this.upperRight.isOnRight(&other.upperRight) && this.upperRight.isUpper(&other.upperRight) {
		return true
	}
	return false
}

// RTreeInternalNode a list of mmb
type RTreeInternalNode struct {
	children *[]MinimumBoundingBox
}

// RTreeLeaveNode or not depends on the disk page: if the disk is full,
// the existed leave node will be splitted
type RTreeLeaveNode struct {
	mmb  *[]MinimumBoundingBox
	size int16
}

//RTree is My implementation of rtree
type RTree struct {
	root *RTreeInternalNode
}

//NewRTree initialization
func NewRTree(objects *[]MinimumBoundingBox) *RTree {
	return nil
}

func (tree *RTree) insert(node *MinimumBoundingBox) {

}

func (tree *RTree) searchPoint(point *Point) {

}
func (tree *RTree) searchBox(box *MinimumBoundingBox) {

}

func (tree *RTree) split(box *RTreeLeaveNode) {

}
