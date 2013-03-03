+(I1, I2) => #+(I1, I2);
/(I1, I2) => #/(I1, I2);
<+(I1, I2) => #<=(I1, I2);
not(true) => false;
not(false) => true;
and(true, B) => B;
and(false, B) => false;
skip => .;

cells(k(~>(=(X, I), K)), state(Sigma))
=>
cells(k(K), state(replace(Sigma, I, X)));

ifThenElse(true, S1, S2) => S1;
ifThenElse(false, S1, S2) => S2;

replace(empty, I, X) => mapItem(X, I);
replace(mapItem(X, J), I, X) => mapItem(X, I);
replace(mapItem(Y, J), I, X) => mapItem(Y, J) when Y != X;

|

cells(k(~>(=(x, 5), .)), state(empty))