package crawls

import (
	"fmt"
	"log"
	"project/vnexpress/api/models"
	"project/vnexpress/config/driver/database"

	"github.com/gocolly/colly"
)

func Crawl() {
	fmt.Println("hello")

	c := colly.NewCollector(colly.AllowedDomains("www.vnexpress.net", "vnexpress.net"))
	// var articles []model.Article

	c.OnHTML("article.item-news", func(e *colly.HTMLElement) {
		title := e.ChildText(".title-news a")
		link := e.ChildAttr(".title-news a", "href")
		image := e.ChildAttr("img[itemprop=contentUrl]", "src")
		description := e.ChildText("p.description")
		if title == "" {
			log.Println("Skipping article with null data.")
			return
		}
		var existingArticle models.Article
		checkurl := database.DBConn.Where("url = ?", link).First(&existingArticle)
		if checkurl.Error == nil {
			// Nếu bài viết đã tồn tại, bỏ qua
			log.Println("Article already exists:", link)
			return
		}

		article := models.Article{Title: title, Link: link, Image: image, Description: description}

		result := database.DBConn.Create(&article)

		if result.Error != nil {
			log.Println("Error in saving data:", result.Error)
		} else {
			log.Println("Article saved successfully.")
		}

	})
	err := c.Visit("https://vnexpress.net/the-thao/euro-2024/tin-tuc")
	if err != nil {
		log.Fatalf("Lỗi khi truy cập URL: %v\n", err)
	}
	fmt.Println("Success!")
}
