#include <iostream>

int main()
{
    int n;

    while (true) {
        std::cin >> n;

        if (n == 0)
            break;

        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {

                int value = abs(i - j) + 1;

                if (value < 10)
                    std::cout << "  ";
                else if (value < 100)
                    std::cout << " ";

                std::cout << value;

                if (j < n)
                    std::cout << " ";
            }

            std::cout << std::endl;
        }

        std::cout << std::endl;
    }

    return 0;
}