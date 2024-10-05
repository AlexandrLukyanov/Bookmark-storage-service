package main

import "fmt"

type bookMarkMap = map[string]string

func main() {
	bookmarks := bookMarkMap{}
	fmt.Println("Приложение для закладок")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBokkmark(bookmarks)
		case 4:
			break Menu

		}
	}

}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант ")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку ")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookMarkMap) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок ")

	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}
}

func addBookmark(bookmarks bookMarkMap) {
	var newBookmarkKey string
	var newBookmarkValue string
	fmt.Print("Введите название :")
	fmt.Scan(&newBookmarkKey)
	fmt.Print("Введите ссылку :")
	fmt.Scan(&newBookmarkValue)
	bookmarks[newBookmarkKey] = newBookmarkValue
}

func deleteBokkmark(bookmarks bookMarkMap) {
	var bookmarkKeyToDelete string
	fmt.Print("Введите название :")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
}
