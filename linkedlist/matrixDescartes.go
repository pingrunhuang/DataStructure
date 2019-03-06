package DataStructure

// Descartes definition:
// given 2 matrix A and B and 2 cols i and j, notation descartes(A,B,i,j) will generate a new matrix C where C[*][i] == C[*][i+j]
// using linked list to define a matrix

func descartes(h1 *tableHead, h2 *tableHead, i int, j int) *tableHead {
	t1 := h1.next
	t2 := h2.next
	resultRow := 0
	resultCol := (*h1).col + (*h2).col
	// new return the pointer
	result := new(tableHead)

	ptr := result.next
	for t1.next != nil {
		for (*t1.data)[i-1] == (*t2.data)[j-1] {
			// make return the actual object not a pointer
			newRow := make([]int, h1.col+h2.col)
			copy(newRow[:h1.col], *t1.data)
			copy(newRow[h1.col:h1.col+h1.col], *t2.data)
			(*ptr).data = &newRow
			ptr = (*ptr).next
			t2 = (*t2).next
			resultRow++
		}
		ptr = (*ptr).next
		t1 = (*t1).next
		resultRow++
	}
	(*result).col = resultCol
	(*result).row = resultRow
	return result
}
