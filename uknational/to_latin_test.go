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
		actual, n, err := transform.String(ToLatin(), input)
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal(expected))
		Expect(n).To(Equal(len(input)))
	},

	// Постанова від 27 січня 2010 р. N 55

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

	// // poor man's position handling:
	// "Є": "Ye", // на початку слова
	// "є": "ie", // в інших позиціях
	PEntry("Є - на початку слова", "Єа", "Yea"),
	PEntry("є - на початку слова", "єа", "yea"),
	PEntry("є - всередині слова", "аєа", "aiea"),
	PEntry("є - в кінці слова", "ає", "aie"),

	// // poor man's position handling:
	// "Ї": "Yi", // на початку слова
	// "ї": "i",  // в інших позиціях
	PEntry("Ї - на початку слова", "Їа", "Yia"),
	PEntry("ї - на початку слова", "їа", "yia"),
	PEntry("ї - всередині слова", "аїа", "aia"),
	PEntry("ї - в кінці слова", "аї", "ai"),

	// // poor man's position handling:
	// "Й": "Y", // на початку слова
	// "й": "i", // в інших позиціях
	PEntry("Й - на початку слова", "Йа", "Ya"),
	PEntry("й - на початку слова", "йа", "ya"),
	PEntry("й - всередині слова", "айа", "aia"),
	PEntry("й - в кінці слова", "ай", "ai"),

	// // poor man's position handling:
	// "Ю": "Yu", // на початку слова
	// "ю": "iu", // в інших позиціях
	PEntry("Ю - на початку слова", "юа", "Yua"),
	PEntry("ю - на початку слова", "юа", "yua"),
	PEntry("ю - всередині слова", "аюа", "aiua"),
	PEntry("ю - в кінці слова", "аю", "aiu"),

	// // poor man's position handling:
	// "Я": "Ya", // на початку слова
	// "я": "ia", // в інших позиціях
	PEntry("Я - на початку слова", "яа", "Yaa"),
	PEntry("я - на початку слова", "яа", "yaa"),
	PEntry("я - всередині слова", "аяа", "aiaa"),
	PEntry("я - в кінці слова", "ая", "aia"),
)
