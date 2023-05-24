package minheap

// MinHeap ist ein Datentyp für  Heaps, bei denen die kleinste Zahl an der Wurzel steht.
// Die Elemente werden in einem Array gespeichert.
type MinHeap struct {
	elements []int
}

// NewMinHeap erzeugt einen neuen leeren Min-Heap.
func NewMinHeap() *MinHeap {
	return &MinHeap{elements: make([]int, 0)}
}

// IsEmpty gibt true zurück, wenn der Heap leer ist.
func (h *MinHeap) IsEmpty() bool {
	return len(h.elements) == 0
}

// Insert fügt ein Element in den Heap ein.
func (h *MinHeap) Insert(value int) {
	h.elements = append(h.elements, value)
	h.BubbleUp(len(h.elements) - 1)
}

// RemoveMin entfernt das kleinste Element aus dem Heap.
// Wenn der Heap leer ist, hat die Funktion keinen Effekt.
func (h *MinHeap) RemoveMin() {
	if h.IsEmpty() {
		return
	}
	last := len(h.elements) - 1
	h.elements[0] = h.elements[last]
	h.elements = h.elements[:last]
	h.BubbleDown(0)
}

// GetMin gibt das kleinste Element zurück.
// Wenn der Heap leer ist, gibt es eine Panic.
func (h *MinHeap) GetMin() int {
	if h.IsEmpty() {
		panic("heap is empty")
	}
	return h.elements[0]
}

// BubbleUp schiebt das Element an der Position pos nach oben, bis die Heap-Bedingung erfüllt ist.
func (h *MinHeap) BubbleUp(pos int) {
	parentpos := (pos - 1) / 2
	if pos == 0 || pos >= len(h.elements) || h.elements[pos] >= h.elements[parentpos] {
		return
	}
	h.elements[pos], h.elements[parentpos] = h.elements[parentpos], h.elements[pos]
	h.BubbleUp(parentpos)
}

// BubbleDown schiebt das Element an der Position pos nach unten, bis die Heap-Bedingung erfüllt ist.
func (h *MinHeap) BubbleDown(pos int) {
	leftchildpos := 2*pos + 1
	rightchildpos := 2*pos + 2
	if leftchildpos >= len(h.elements) {
		return
	}
	childpos := leftchildpos
	if rightchildpos < len(h.elements) && h.elements[rightchildpos] < h.elements[leftchildpos] {
		childpos = rightchildpos
	}
	if h.elements[pos] <= h.elements[childpos] {
		return
	}
	h.elements[pos], h.elements[childpos] = h.elements[childpos], h.elements[pos]
	h.BubbleDown(childpos)
}
