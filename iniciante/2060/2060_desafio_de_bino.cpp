#include <iostream>

bool multiplode(int numero, int divisor)
{
	if (numero % divisor == 0) {
		return true;
	}

	return false;
}

int main()
{
	int N {};
	std::cin >> N;

	int mult2 {0};
	int mult3 {0};
	int mult4 {0};
	int mult5 {0};

	for (int i = 0; i < N; i++) {
		int L {};
		std::cin >> L;

		if (multiplode(L, 2)) {
			mult2++;
		}

		if (multiplode(L, 3)) {
			mult3++;
		}

		if (multiplode(L, 4)) {
			mult4++;
		}

		if (multiplode(L,5)) {
			mult5++;
		}	
	}
	
	std::cout << mult2 << " Multiplo(s) de 2" << std::endl;
	std::cout << mult3 << " Multiplo(s) de 3" << std::endl;
	std::cout << mult4 << " Multiplo(s) de 4" << std::endl;
	std::cout << mult5 << " Multiplo(s) de 5" << std::endl;

	return 0;
}
