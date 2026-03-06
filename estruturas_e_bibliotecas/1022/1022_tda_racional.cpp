#include <iostream>
#include <string>
#include <cmath>

int mdc(int a, int b) {
	if (b == 0) {
		return std::abs(a);
	}
	return mdc(b, a % b);
}

int main()
{
	int N {};
	std::cin >> N;

	int N1 {};
	int N2 {};
	int D1 {};
	int D2 {};
	char barra1, barra2, op;

	for (int i = 0; i < N; i++) {
		std::cin >> N1 >> barra1 >> D1 >> op >> N2 >> barra2 >> D2;

		int numerador {};
		int denominador {};

		if (op == '+') {
			numerador = (N1*D2 + N2*D1);
			denominador = (D1*D2);
		} else if (op == '-') {
			numerador = (N1*D2 - N2*D1);
			denominador = (D1*D2);
		} else if (op == '*') {
			numerador = (N1*N2);
			denominador = (D1*D2);
		} else if (op == '/') {
			numerador = (N1*D2);
			denominador = (N2*D1);
		}
		
		int max_divisor = mdc(numerador, denominador);
		int numerador_simplificado = numerador/max_divisor;
		int denominador_simplificado = denominador/max_divisor;

		std::cout << numerador << "/" << denominador << " = " << numerador_simplificado << "/" << denominador_simplificado << std::endl;
	}	
}
