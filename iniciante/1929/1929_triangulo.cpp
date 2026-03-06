#include <iostream>

bool triangulo(int a, int b, int c) 
{
	return (a + b > c) && (a + c > b) && (b + c > a);
}

int main()
{
	int A, B, C, D;
	std::cin >> A >> B >> C >> D;	

	if (triangulo(A,B,C) || triangulo(A,B,D) || triangulo(A,C,D) || triangulo (B,C,D)) {
		std::cout << "S" << std::endl;
	} else {
		std::cout << "N" << std::endl;
	}

	return 0;
}
