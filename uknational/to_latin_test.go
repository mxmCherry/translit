package uknational_test

import (
	"golang.org/x/text/transform"

	. "github.com/mxmCherry/translit/uknational"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable(
	"ToLatin",

	func(input, expected string) {
		actual, n, err := transform.String(ToLatin().Transformer(), input)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal(expected))
		Expect(n).To(Equal(len(input)))
	},

	// Постанова від 27 січня 2010 р. N 55
	// http://zakon.rada.gov.ua/laws/show/55-2010-п

	// Таблиця транслітерації українського алфавіту латиницею
	Entry("Алушта", "Алушта", "Alushta"),
	Entry("Андрій", "Андрій", "Andrii"),
	Entry("Борщагівка", "Борщагівка", "Borshchahivka"),
	Entry("Борисенко", "Борисенко", "Borysenko"),
	Entry("Вінниця", "Вінниця", "Vinnytsia"),
	Entry("Володимир", "Володимир", "Volodymyr"),
	Entry("Гадяч", "Гадяч", "Hadiach"),
	Entry("Богдан", "Богдан", "Bohdan"),
	Entry("Згурський", "Згурський", "Zghurskyi"),
	Entry("Ґалаґан", "Ґалаґан", "Galagan"),
	Entry("Ґорґани", "Ґорґани", "Gorgany"),
	Entry("Донецьк", "Донецьк", "Donetsk"),
	Entry("Дмитро", "Дмитро", "Dmytro"),
	Entry("Рівне", "Рівне", "Rivne"),
	Entry("Олег", "Олег", "Oleh"),
	Entry("Есмань", "Есмань", "Esman"),
	Entry("Єнакієве", "Єнакієве", "Yenakiieve"),
	Entry("Гаєвич", "Гаєвич", "Haievych"),
	Entry("Короп'є", "Короп'є", "Koropie"),
	Entry("Житомир", "Житомир", "Zhytomyr"),
	Entry("Жанна", "Жанна", "Zhanna"),
	Entry("Жежелів", "Жежелів", "Zhezheliv"),
	Entry("Закарпаття", "Закарпаття", "Zakarpattia"),
	Entry("Казимирчук", "Казимирчук", "Kazymyrchuk"),
	Entry("Медвин", "Медвин", "Medvyn"),
	Entry("Михайленко", "Михайленко", "Mykhailenko"),
	Entry("Іванків", "Іванків", "Ivankiv"),
	Entry("Іващенко", "Іващенко", "Ivashchenko"),
	Entry("Їжакевич", "Їжакевич", "Yizhakevych"),
	Entry("Кадиївка", "Кадиївка", "Kadyivka"),
	Entry("Мар'їне", "Мар'їне", "Marine"),
	Entry("Йосипівка", "Йосипівка", "Yosypivka"),
	Entry("Стрий", "Стрий", "Stryi"),
	Entry("Олексій", "Олексій", "Oleksii"),
	Entry("Київ", "Київ", "Kyiv"),
	Entry("Коваленко", "Коваленко", "Kovalenko"),
	Entry("Лебедин", "Лебедин", "Lebedyn"),
	Entry("Леонід", "Леонід", "Leonid"),
	Entry("Миколаїв", "Миколаїв", "Mykolaiv"),
	Entry("Маринич", "Маринич", "Marynych"),
	Entry("Ніжин", "Ніжин", "Nizhyn"),
	Entry("Наталія", "Наталія", "Nataliia"),
	Entry("Одеса", "Одеса", "Odesa"),
	Entry("Онищенко", "Онищенко", "Onyshchenko"),
	Entry("Полтава", "Полтава", "Poltava"),
	Entry("Петро", "Петро", "Petro"),
	Entry("Решетилівка", "Решетилівка", "Reshetylivka"),
	Entry("Рибчинський", "Рибчинський", "Rybchynskyi"),
	Entry("Суми", "Суми", "Sumy"),
	Entry("Соломія", "Соломія", "Solomiia"),
	Entry("Тернопіль", "Тернопіль", "Ternopil"),
	Entry("Троць", "Троць", "Trots"),
	Entry("Ужгород", "Ужгород", "Uzhhorod"),
	Entry("Уляна", "Уляна", "Uliana"),
	Entry("Фастів", "Фастів", "Fastiv"),
	Entry("Філіпчук", "Філіпчук", "Filipchuk"),
	Entry("Харків", "Харків", "Kharkiv"),
	Entry("Христина", "Христина", "Khrystyna"),
	Entry("Біла Церква", "Біла Церква", "Bila Tserkva"),
	Entry("Стеценко", "Стеценко", "Stetsenko"),
	Entry("Чернівці", "Чернівці", "Chernivtsi"),
	Entry("Шевченко", "Шевченко", "Shevchenko"),
	Entry("Шостка", "Шостка", "Shostka"),
	Entry("Кишеньки", "Кишеньки", "Kyshenky"),
	Entry("Щербухи", "Щербухи", "Shcherbukhy"),
	Entry("Гоща", "Гоща", "Hoshcha"),
	Entry("Гаращенко", "Гаращенко", "Harashchenko"),
	Entry("Юрій", "Юрій", "Yurii"),
	Entry("Корюківка", "Корюківка", "Koriukivka"),
	Entry("Яготин", "Яготин", "Yahotyn"),
	Entry("Ярошенко", "Ярошенко", "Yaroshenko"),
	Entry("Костянтин", "Костянтин", "Kostiantyn"),
	Entry("Знам'янка", "Знам'янка", "Znamianka"),
	Entry("Феодосія", "Феодосія", "Feodosiia"),

	// 1. Буквосполучення "зг" відтворюється латиницею як "zgh"
	// (наприклад,  Згорани - Zghorany, Розгон - Rozghon) на
	// відміну від "zh" - відповідника української літери
	// "ж".
	Entry("Згорани", "Згорани", "Zghorany"),
	Entry("Розгон", "Розгон", "Rozghon"),

	// 2. М'який знак і апостроф латиницею не відтворюються.
	Entry("ь", "ь", ""),
	Entry("Ь", "Ь", ""),
	Entry("'", "'", ""),
	Entry("’", "’", ""),

	// на початку слова: ye
	// в інших позиціях: ie
	Entry("Є - на початку слова", "Єа", "Yea"),
	Entry("є - на початку слова", "єа", "yea"),
	Entry("є - всередині слова", "аєа", "aiea"),
	Entry("є - в кінці слова", "ає", "aie"),

	// на початку слова: yi
	// в інших позиціях: i
	Entry("Ї - на початку слова", "Їа", "Yia"),
	Entry("ї - на початку слова", "їа", "yia"),
	Entry("ї - всередині слова", "аїа", "aia"),
	Entry("ї - в кінці слова", "аї", "ai"),

	// на початку слова: y
	// в інших позиціях: i
	Entry("Й - на початку слова", "Йа", "Ya"),
	Entry("й - на початку слова", "йа", "ya"),
	Entry("й - всередині слова", "айа", "aia"),
	Entry("й - в кінці слова", "ай", "ai"),

	// на початку слова: yu
	// в інших позиціях: iu
	Entry("Ю - на початку слова", "Юа", "Yua"),
	Entry("ю - на початку слова", "юа", "yua"),
	Entry("ю - всередині слова", "аюа", "aiua"),
	Entry("ю - в кінці слова", "аю", "aiu"),

	// на початку слова: ya
	// в інших позиціях: ia
	Entry("Я - на початку слова", "Яа", "Yaa"),
	Entry("я - на початку слова", "яа", "yaa"),
	Entry("я - всередині слова", "аяа", "aiaa"),
	Entry("я - в кінці слова", "ая", "aia"),
)
