X => Y;
// Y => Z;
// +(X, s(Y)) => s(+(X, Y));
// +(X, 0) => X;

// *(X, s(Y)) => +(*(X, Y), Y);
// *(X, 0) => 0;

// ^(X, s(Y)) => *(^(X, Y), X);
// ^(X, 0) => 1;