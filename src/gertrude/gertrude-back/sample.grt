+(X, s(Y)) => s(+(X, Y));
+(X, 0) => X;

*(X, s(Y)) => +(*(X, Y), X);
*(X, 0) => 0;

^(X, s(Y)) => *(^(X, Y), X);
^(X, 0) => s(0);

|

+(s(s(s(s(s(0))))), s(s(s(s(0)))))
asdfasdf
