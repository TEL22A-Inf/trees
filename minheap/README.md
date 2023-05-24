# Heaps

## Definition

Ein *Min-Heap* ist ein binärer Baum, der folgende Eigenschaften erfüllt:

* Der Baum ist vollständig. D.h. alle Ebenen sind vollständig gefüllt,
  bis auf die letzte, die von links nach rechts gefüllt ist.
* Der Schlüssel eines Knotens ist immer kleiner als die Schlüssel seiner Kinder.
* Zwischen den Kindern eines Knotens gibt es keine Ordnungsbeziehung.

Definition eines *Max-Heaps* analog.

## Eigenschaften

* Die Höhe eines Heaps mit $n$ Elementen ist $O(log\ n)$.
  * Daraus folgt, dass die Operationen `Insert` und `RemoveMin` in $O(log\ n)$
    ausgeführt werden können.
* Das Minimum eines Min-Heaps ist die Wurzel.
  * D.h. `GetMin` kann in $O(1)$ ausgeführt werden.

## Operationen

### `Insert`

1. Füge das Element am Ende des Heaps ein.
2. Führe `BubbleUp` aus, um die Heap-Eigenschaft wiederherzustellen.

### `RemoveMin`

1. Vertausche die Wurzel mit dem letzten Element des Heaps.
2. Entferne das letzte Element.
3. Führe `BubbleDown` aus, um die Heap-Eigenschaft wiederherzustellen.

### `GetMin`

1. Liefere die Wurzel.

### `BubbleUp`

1. Vergleiche den Schlüssel des Knotens mit dem Schlüssel seines Elternknotens.
2. Falls der Schlüssel des Knotens kleiner ist als der des Elternknotens, vertausche die Knoten.

### `BubbleDown`

1. Vergleiche den Schlüssel des Knotens mit dem Schlüssel seines kleineren Kindes.
2. Falls der Schlüssel des Knotens größer ist als der des kleineren Kindes, vertausche die Knoten.

### Speicherung des Baums, Position der Knoten

* Der Baum kann als Slice gespeichert werden.
  * Eine Pointerstruktur wie bei binären Suchbäumen ist nicht notwendig.
  * Das liegt an der Vollständigkeit des Baumes, es gibt keine Lücken.
* Die Position eines Knotens kann wie folgt berechnet werden:
  * Der Index der Wurzel ist 0.
  * Der Index des linken Kindes eines Knotens $n$ ist $2n+1$.
  * Der Index des rechten Kindes eines Knotens $n$ ist $2n+2$.
  * Der Index des Elternknotens eines Knotens $n$ ist $\lfloor\frac{n-1}{2}\rfloor$.
