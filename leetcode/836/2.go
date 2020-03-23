func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(
		rec1[2] <= rec2[0] ||  //left
		rec1[3] <= rec2[1] || //bottom
		rec2[2] <= rec1[0] || //right
		rec2[3] <= rec1[1]     //top
	)
}	