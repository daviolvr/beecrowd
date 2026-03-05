#include <iostream>

int main() 
{
	int n = 0;
	int velocidade = 0;
	
	while (std::cin >> n) {
		int maior = 0;
		for (int i = 0; i < n; i++) {
			int velocidade;
			std::cin >> velocidade;
			if (velocidade > maior) {
				maior = velocidade;
			}	
		}

		if (maior < 10) {
			std::cout << "1" << std::endl;
		} else if (maior >= 10 && maior < 20) {
			std::cout << "2" << std::endl;
		} else {
			std::cout << "3" << std::endl;
		}
	}
	return 0;	
}
