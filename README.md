# Kryptonim-Interview

Cześć!

W tym projekcie zrealizowałem zadanie rekrutacyjne, które obejmuje implementację kilku funkcji API oraz testy integracyjne. Poniżej przedstawiam szczegóły realizacji.

## Implementacja

### 1. Endpoints:
- **/rates** – implementacja endpointu zgodnie z wymaganiami.
- **/exchange** – implementacja endpointu zgodnie z wymaganiami.

### 2. Testy:
- Testy integracyjne zostały zaimplementowane przy użyciu mikroskopijnej wersji Cucumber w Go.
- Zaimplementowany mechanizm mockowania dla zewnętrznego API.

### 3. Pokrycie testami:
- Zgodnie z opisem zadania zadbałem o pełne pokrycie testami (testy integracyjne).

### 4. Dockeryzacja:
- W projekcie znajduje się Docker, który umożliwia uruchamianie testów integracyjnych lokalnie.
- **Entrypoint** w Dockerze umożliwia automatyczne odświeżanie testów po każdej zmianie w kodzie.
- Bardzo lubie takie rozwiazanie przy pracy w podejściu TDD jest to bardzo wygodne, jeśli chodzi o to czemu BDD - to dlatego ze to po prostu jest świetne, im szybciej się to ma w projekcie tym lepiej bo po prostu ułatwia pisanie, czytania i przyspiesza tworzenie testów. Czyste testy z go są po prostu dramatyczne przy rosnącej aplikacji.

## Uruchomienie lokalne testów integracyjnych

1. Zbudowanie kontenerów:

    ```bash
    docker-compose build
    ```

2. Uruchomienie kontenerów:

    ```bash
    docker-compose up
    ```

**Dodatkowo**: W pliku `entrypoint.sh` zawarłem mechanizm umożliwiający tagowanie scenariuszy testowych, dzięki czemu można łatwiej pracować nad konkretnymi przypadkami, szczególnie przy długich testach.

## Podejście do architektury

### Wzorce projektowe
W projekcie zastosowałem kilka popularnych wzorców projektowych, które ułatwiają rozdzielanie odpowiedzialności i poprawiają testowalność oraz skalowalność aplikacji:

1. **Command Pattern**: Zastosowany wzorzec Command pozwala na późniejszy prosty deployment razem z plikami do k8s. 

2. **Repository Pattern**: Zaimplementowałem wzorzec Repository do oddzielenia logiki dostępu do danych (baza danych) od reszty aplikacji. Dzięki temu kod staje się bardziej testowalny i łatwiejszy do utrzymania.

3. **Service Pattern**: Wzorzec Service został wykorzystany do oddzielenia logiki biznesowej od reszty aplikacji. To sprawia, że kod staje się bardziej modularny i umożliwia łatwe zarządzanie oraz testowanie różnych funkcjonalności aplikacji.

### Lazy Loading
Pewnie będzie pytanie o lazy loading - ze względu na doświadczenie w pracy z aplikacjami chmurowymi. Zasoby są ładowane tylko wtedy, gdy są naprawdę potrzebne, co ma znaczenie w kontekście środowisk chmurowych chociażby - szybszy deployment, jednorazowe otwieranie połączenia do db i kilka innych jeszcze ale to do pogadania/dyskusji. 

## Brak implementacji autoryzacji

W tej wersji aplikacji **nie zaimplementowałem systemu autoryzacji**. Skupiłem się na implementacji core'owej logiki API i testów integracyjnych, zakładając, że kwestie związane z autoryzacją (np. JWT, OAuth2) mogą być dodane później, w miarę jak aplikacja by rosła i przechodziła do środowisk produkcyjnych.

## Struktura projektu

Pod względem struktury projektu, zachowałem większą elastyczność, ponieważ na etapie wczesnej implementacji aplikacji (gdy jest jeszcze mała) nie jestem fanem sztywnych zasad dotyczących organizacji plików i folderów. Zamiast tego, struktura jest raczej luźna i opiera się na intuicji oraz potrzebach rozwoju aplikacji.

## Logowanie

Zdecydowałem się na stosowanie logów w formacie `[211616]` jako unikalnego identyfikatora dla każdej operacji. Umożliwia to łatwe śledzenie operacji w logach, szczególnie w systemach rozproszonych (np. w chmurze). Dzięki temu mogę szybko zidentyfikować, gdzie w kodzie wystąpił problem, korzystając z unikalnego identyfikatora, nie piszę ich z reki mam dodaną regułe na keyword "rtime" w IDE i mi generuje te cyferki z daty/sekundy etc. To o tyle fajne ze moge potem wziąć taki kod i zrobić ctrl+f w projekcie i odrazu wiem w kodzie gdzie to poszło :)  

## Dlaczego tagowanie scenariuszy?

W testach integracyjnych, szczególnie w TDD, korzystam z tagowania scenariuszy testowych. Dzięki temu mogę łatwiej pracować nad poszczególnymi przypadkami, weryfikować je i debugować w prostszy sposób. Jest to szczególnie przydatne, gdy mamy do czynienia z dużą liczbą testów.

## Komentarze w kodzie

W kodzie pozostawiłem kilka komentarzy, które mają charakter informacyjny. Wyjaśniają one specyficzne decyzje projektowe i operacje, które mogą być istotne w kontekście dalszej pracy nad projektem.

## Technologie

- **Go (Golang)** – język programowania, w którym napisana jest aplikacja.
- **Docker** – do konteneryzacji aplikacji i testów integracyjnych.
- **Cucumber (Go)** – framework do pisania testów integracyjnych z podejściem BDD.
- **PostgreSQL** – baza danych używana do przechowywania danych.

---

## Podsumowanie

Wiem że pewnie możecie wymagać dokeryzacji która odpali wam to lokalnie, tylko po co? Nie odpala się dziś apek lokalnie z dockera gdy działami z chmurami, przygotowałem wszystko pod to żeby dało się tego użyć w konfigach do k8s (wiadomo nie ma Dockerfile dla samej binarki, i nie wrzucałem migrate command bo bym go tutaj i tak nie użył) No i zostawiłem cmd do api z wbitym adresem na 0.0.0.0 a ogólnie to by szło z configu z resztą konfiguracji routera. 

Mam nadzieję, że to spełnia oczekiwania, chętnie pogadam bo starałem się tak zahaczyć o rózne tematy żeby nie było strasznie sucho skoro zadanie ma być do przejrzenia i potem do porozmawiania to trakuje to raczej jako próbkę kodu :) 
