package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	FeaturedPosts []featuredPostData
	AboutPosts    []aboutPostData
}

type postPage struct {
	Title    string
	Subtitle string
	Image    string
	Name     string
	Page     []pagepostData
}

type featuredPostData struct {
	Item             string
	Title            string
	Subtitle         string
	Button           string
	Text             string
	Date             string
	Icon             string
	Features__button string
}

type aboutPostData struct {
	Title    string
	Text     string
	Subtitle string
	Icon     string
	Image    string
	Date     string
}

type pagepostData struct {
	P string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		FeaturedPosts: featuredPosts(),
		AboutPosts:    aboutPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := postPage{
		Title:    "The Road Ahead",
		Subtitle: "The road ahead might be paved - it might not be.",
		Image:    "../static/images/the_road_ahead/img.png",
		Name:     "main image",
		Page:     pagePosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Item:             "features__item",
			Button:           "",
			Features__button: "",
			Title:            "The Road Ahead",
			Subtitle:         "The road ahead might be paved - it might not be.",
			Icon:             "../static/images/features/icon/icon_l.png",
			Text:             "Mat Vogels",
			Date:             "September 25, 2015",
		},
		{
			Item:             "features__item features__item-r",
			Features__button: "button features__btn",
			Button:           "Adventure",
			Title:            "From Top Down",
			Subtitle:         "Once a year, go someplace you’ve never been before.",
			Icon:             "../static/images/features/icon/icon_r.png",
			Text:             "William Wong",
			Date:             "September 25, 2015",
		},
	}
}

func aboutPosts() []aboutPostData {
	return []aboutPostData{
		{
			Image:    "../static/images/about/about1.png",
			Title:    "Still Standing Tall",
			Subtitle: "Life begins at the end of your comfort zone.",
			Icon:     "../static/images/about/icon/icon_l.png",
			Text:     "William Wong",
			Date:     "9/25/2015",
		},
		{
			Image:    "../static/images/about/about2.png",
			Title:    "Sunny Side Up",
			Subtitle: "No place is ever as bad as they tell you it’s going to be.",
			Icon:     "../static/images/about/icon/icon_r.png",
			Text:     "Mat Vogels",
			Date:     "9/25/2015",
		},
		{
			Image:    "../static/images/about/about3.png",
			Title:    "Water Falls",
			Subtitle: "We travel not to escape life, but for life not to escape us.",
			Icon:     "../static/images/about/icon/icon_r.png",
			Text:     "Mat Vogels",
			Date:     "9/25/2015",
		},
		{
			Image:    "../static/images/about/about4.png",
			Title:    "Through the Mist",
			Subtitle: "Travel makes you see what a tiny place you occupy in the world.",
			Icon:     "../static/images/about/icon/icon_l.png",
			Text:     "William Wong",
			Date:     "9/25/2015",
		},
		{
			Image:    "../static/images/about/about5.png",
			Title:    "Awaken Early",
			Subtitle: "Not all those who wander are lost.",
			Icon:     "../static/images/about/icon/icon_r.png",
			Text:     "Mat Vogels",
			Date:     "9/25/2015",
		},
		{
			Image:    "../static/images/about/about6.png",
			Title:    "Try it Always",
			Subtitle: "The world is a book, and those who do not travel read only one page.",
			Icon:     "../static/images/about/icon/icon_l.png",
			Text:     "William Wong",
			Date:     "9/25/2015",
		},
	}
}

func pagePosts() []pagepostData {
	return []pagepostData{
		{
			P: "Dark spruce forest frowned on either side the frozen waterway. The trees had been stripped by a recent wind of their white covering of frost, and they seemed to lean towards each other, black and ominous, in the fading light. A vast silence reigned over the land. The land itself was a desolation, lifeless, without movement, so lone and cold that the spirit of it was not even that of sadness. There was a hint in it of laughter, but of a laughter more terrible than any sadness-a laughter that was mirthless as the smile of the sphinx, a laughter cold as the frost and partaking of the grimness of infallibility. It was the masterful and incommunicable wisdom of eternity laughing at the futility of life and the effort of life. It was the Wild, the savage, frozen-hearted Northland Wild.",
		},
		{
			P: "But there was life, abroad in the land and defiant. Down the frozen waterway toiled a string of wolfish dogs. Their bristly fur was rimed with frost. Their breath froze in the air as it left their mouths, spouting forth in spumes of vapour that settled upon the hair of their bodies and formed into crystals of frost. Leather harness was on the dogs, and leather traces attached them to a sled which dragged along behind. The sled was without runners. It was made of stout birch-bark, and its full surface rested on the snow. The front end of the sled was turned up, like a scroll, in order to force down and under the bore of soft snow that surged like a wave before it. On the sled, securely lashed, was a long and narrow oblong box. There were other things on the sled—blankets, an axe, and a coffee-pot and frying-pan; but prominent, occupying most of the space, was the long and narrow oblong box.",
		},
		{
			P: "In advance of the dogs, on wide snowshoes, toiled a man. At the rear of the sled toiled a second man. On the sled, in the box, lay a third man whose toil was over,—a man whom the Wild had conquered and beaten down until he would never move nor struggle again. It is not the way of the Wild to like movement. Life is an offence to it, for life is movement; and the Wild aims always to destroy movement. It freezes the water to prevent it running to the sea; it drives the sap out of the trees till they are frozen to their mighty hearts; and most ferociously and terribly of all does the Wild harry and crush into submission man—man who is the most restless of life, ever in revolt against the dictum that all movement must in the end come to the cessation of movement.",
		},
		{
			P: "But at front and rear, unawed and indomitable, toiled the two men who were not yet dead. Their bodies were covered with fur and soft-tanned leather. Eyelashes and cheeks and lips were so coated with the crystals from their frozen breath that their faces were not discernible. This gave them the seeming of ghostly masques, undertakers in a spectral world at the funeral of some ghost. But under it all they were men, penetrating the land of desolation and mockery and silence, puny adventurers bent on colossal adventure, pitting themselves against the might of a world as remote and alien and pulseless as the abysses of space.",
		},
	}
}
