#include <iostream>

int main()
{
	int S {};
	int T {};
	int F {};
	int horaChegada {};

	std::cin >> S >> T >> F;

	int chegada = (S + T + F) % 24;

	if (chegada < 0) {
		chegada += 24;
	}

	std::cout << chegada << std::endl;

	return 0;
}

