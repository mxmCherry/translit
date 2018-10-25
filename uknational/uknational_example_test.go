package uknational_test

import (
	"fmt"

	"github.com/mxmCherry/translit/uknational"
)

func ExampleConverter() {
	uk := uknational.Converter()

	fmt.Println(uk.Convert("Алушта"))      // Alushta
	fmt.Println(uk.Convert("Андрій"))      // Andrii
	fmt.Println(uk.Convert("Борщагівка"))  // Borshchahivka
	fmt.Println(uk.Convert("Борисенко"))   // Borysenko
	fmt.Println(uk.Convert("Вінниця"))     // Vinnytsia
	fmt.Println(uk.Convert("Володимир"))   // Volodymyr
	fmt.Println(uk.Convert("Гадяч"))       // Hadiach
	fmt.Println(uk.Convert("Богдан"))      // Bohdan
	fmt.Println(uk.Convert("Згурський"))   // Zghurskyi
	fmt.Println(uk.Convert("Ґалаґан"))     // Galagan
	fmt.Println(uk.Convert("Ґорґани"))     // Gorgany
	fmt.Println(uk.Convert("Донецьк"))     // Donetsk
	fmt.Println(uk.Convert("Дмитро"))      // Dmytro
	fmt.Println(uk.Convert("Рівне"))       // Rivne
	fmt.Println(uk.Convert("Олег"))        // Oleh
	fmt.Println(uk.Convert("Есмань"))      // Esman
	fmt.Println(uk.Convert("Єнакієве"))    // Yenakiieve
	fmt.Println(uk.Convert("Гаєвич"))      // Haievych
	fmt.Println(uk.Convert("Короп'є"))     // Koropie
	fmt.Println(uk.Convert("Житомир"))     // Zhytomyr
	fmt.Println(uk.Convert("Жанна"))       // Zhanna
	fmt.Println(uk.Convert("Жежелів"))     // Zhezheliv
	fmt.Println(uk.Convert("Закарпаття"))  // Zakarpattia
	fmt.Println(uk.Convert("Казимирчук"))  // Kazymyrchuk
	fmt.Println(uk.Convert("Медвин"))      // Medvyn
	fmt.Println(uk.Convert("Михайленко"))  // Mykhailenko
	fmt.Println(uk.Convert("Іванків"))     // Ivankiv
	fmt.Println(uk.Convert("Іващенко"))    // Ivashchenko
	fmt.Println(uk.Convert("Їжакевич"))    // Yizhakevych
	fmt.Println(uk.Convert("Кадиївка"))    // Kadyivka
	fmt.Println(uk.Convert("Мар'їне"))     // Marine
	fmt.Println(uk.Convert("Йосипівка"))   // Yosypivka
	fmt.Println(uk.Convert("Стрий"))       // Stryi
	fmt.Println(uk.Convert("Олексій"))     // Oleksii
	fmt.Println(uk.Convert("Київ"))        // Kyiv
	fmt.Println(uk.Convert("Коваленко"))   // Kovalenko
	fmt.Println(uk.Convert("Лебедин"))     // Lebedyn
	fmt.Println(uk.Convert("Леонід"))      // Leonid
	fmt.Println(uk.Convert("Миколаїв"))    // Mykolaiv
	fmt.Println(uk.Convert("Маринич"))     // Marynych
	fmt.Println(uk.Convert("Ніжин"))       // Nizhyn
	fmt.Println(uk.Convert("Наталія"))     // Nataliia
	fmt.Println(uk.Convert("Одеса"))       // Odesa
	fmt.Println(uk.Convert("Онищенко"))    // Onyshchenko
	fmt.Println(uk.Convert("Полтава"))     // Poltava
	fmt.Println(uk.Convert("Петро"))       // Petro
	fmt.Println(uk.Convert("Решетилівка")) // Reshetylivka
	fmt.Println(uk.Convert("Рибчинський")) // Rybchynskyi
	fmt.Println(uk.Convert("Суми"))        // Sumy
	fmt.Println(uk.Convert("Соломія"))     // Solomiia
	fmt.Println(uk.Convert("Тернопіль"))   // Ternopil
	fmt.Println(uk.Convert("Троць"))       // Trots
	fmt.Println(uk.Convert("Ужгород"))     // Uzhhorod
	fmt.Println(uk.Convert("Уляна"))       // Uliana
	fmt.Println(uk.Convert("Фастів"))      // Fastiv
	fmt.Println(uk.Convert("Філіпчук"))    // Filipchuk
	fmt.Println(uk.Convert("Харків"))      // Kharkiv
	fmt.Println(uk.Convert("Христина"))    // Khrystyna
	fmt.Println(uk.Convert("Біла Церква")) // Bila Tserkva
	fmt.Println(uk.Convert("Стеценко"))    // Stetsenko
	fmt.Println(uk.Convert("Чернівці"))    // Chernivtsi
	fmt.Println(uk.Convert("Шевченко"))    // Shevchenko
	fmt.Println(uk.Convert("Шостка"))      // Shostka
	fmt.Println(uk.Convert("Кишеньки"))    // Kyshenky
	fmt.Println(uk.Convert("Щербухи"))     // Shcherbukhy
	fmt.Println(uk.Convert("Гоща"))        // Hoshcha
	fmt.Println(uk.Convert("Гаращенко"))   // Harashchenko
	fmt.Println(uk.Convert("Юрій"))        // Yurii
	fmt.Println(uk.Convert("Корюківка"))   // Koriukivka
	fmt.Println(uk.Convert("Яготин"))      // Yahotyn
	fmt.Println(uk.Convert("Ярошенко"))    // Yaroshenko
	fmt.Println(uk.Convert("Костянтин"))   // Kostiantyn
	fmt.Println(uk.Convert("Знам'янка"))   // Znamianka
	fmt.Println(uk.Convert("Феодосія"))    // Feodosiia

	// Output:
	// Alushta
	// Andrii
	// Borshchahivka
	// Borysenko
	// Vinnytsia
	// Volodymyr
	// Hadiach
	// Bohdan
	// Zghurskyi
	// Galagan
	// Gorgany
	// Donetsk
	// Dmytro
	// Rivne
	// Oleh
	// Esman
	// Yenakiieve
	// Haievych
	// Koropie
	// Zhytomyr
	// Zhanna
	// Zhezheliv
	// Zakarpattia
	// Kazymyrchuk
	// Medvyn
	// Mykhailenko
	// Ivankiv
	// Ivashchenko
	// Yizhakevych
	// Kadyivka
	// Marine
	// Yosypivka
	// Stryi
	// Oleksii
	// Kyiv
	// Kovalenko
	// Lebedyn
	// Leonid
	// Mykolaiv
	// Marynych
	// Nizhyn
	// Nataliia
	// Odesa
	// Onyshchenko
	// Poltava
	// Petro
	// Reshetylivka
	// Rybchynskyi
	// Sumy
	// Solomiia
	// Ternopil
	// Trots
	// Uzhhorod
	// Uliana
	// Fastiv
	// Filipchuk
	// Kharkiv
	// Khrystyna
	// Bila Tserkva
	// Stetsenko
	// Chernivtsi
	// Shevchenko
	// Shostka
	// Kyshenky
	// Shcherbukhy
	// Hoshcha
	// Harashchenko
	// Yurii
	// Koriukivka
	// Yahotyn
	// Yaroshenko
	// Kostiantyn
	// Znamianka
	// Feodosiia
}
