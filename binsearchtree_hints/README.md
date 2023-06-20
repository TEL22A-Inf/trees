Binäre Suchbäume
================

Beispielimplementierung eines binären Suchbaums.

Anmerkungen zu den Aufgaben
---------------------------

In den Dateien `tree_aufgaben.go` und `element_aufgaben.go` sind einige Methoden bzw. Funktionen zu den Datentypen
`Tree` und `Element` vorgegeben, die Sie vervollständigen sollen.
Dazu gibt es noch einige Anmerkungen:

* Die Methoden `GetMinNode` und `GetMaxNode` für `Element` sind hilfreich zur Implementierung
  von `RemoveValue` in `Tree`.

* Die Methoden bzw. Funktionen ab `Height` werden für die Umsetzung eines AVL-Baums benötigt.

* Damit die AVL-Baum-Implementierung funktioniert, muss die Methode `InsertValue` in `Tree`
  korrekt angepasst werden.

* Damit die AVL-Baum-Implementierung wirklich performant wird, muss auch die Struktur
  `Element` angepasst werden.

* Damit der Baum tatsächlich nützlich wird, sollte der einfache Wert `Value` durch ein
  Schlüssel-Wert-Paar ersetzt werden. Dazu muss die Struktur `Element` angepasst werden.
  Anschließend hat es auch Sinn, in `Tree` eine Methode zur Suche nach Schlüsseln
  umzusetzen.
