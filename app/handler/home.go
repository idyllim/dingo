package handler

import (
	"github.com/dinever/dingo/app/model"
	"github.com/dinever/golf"
	"log"
	"strconv"
)

//func getArticleListSize() int {
//	size, _ := strconv.Atoi(model.GetSetting("article_size"))
//	if size < 1 {
//		size = 5
//	}
//	return size
//}

func getAllTags() []*model.Tag {
	tags, err := model.GetAllTags()
	if err != nil {
		log.Printf("[Error] Can not get tags: %v", err.Error())
	}
	return tags
}

//func updateSidebarData(data map[string]interface{}) {
//	popNumStr, err := model.GetSettingValue("popular_size")
//	if err != nil {
//		popNumStr = "5"
//	}
//	popNum, _ := strconv.Atoi(popNumStr)
//	if popNum < 1 {
//		popNum = 4
//	}
//	cmtNumStr, err := model.GetSettingValue("recent_comment_size")
//	if err != nil {
//		cmtNumStr = "5"
//	}
//	cmtNum, _ := strconv.Atoi(cmtNumStr)
//	if cmtNum < 1 {
//		cmtNum = 3
//	}
//	recentNumStr, err := model.GetSettingValue("popular_size")
//	if err != nil {
//		recentNumStr = "5"
//	}
//	recentNum, _ := strconv.Atoi(recentNumStr)
//	if recentNum < 1 {
//		recentNum = 4
//	}
//	data["Popular"], _, _ = model.GetPostList(1, int64(popNum), false, true, "created_at")
//	data["RecentArticles"], _, _ = model.GetPostList(1, int64(popNum), false, true, "published_at")
//	data["RecentComment"], _, _ = model.GetPostList(1, int64(popNum), false, true, "updated_at")
//	data["Tags"], _ = model.GetAllTags()
//}

func HomeHandler(ctx *Golf.Context) {
	p, _ := ctx.Param("page")
	page, _ := strconv.Atoi(p)
	articles, pager, _ := model.GetPostList(int64(page), 5, false, true, "published_at")
	// theme := model.GetSetting("site_theme")
	data := map[string]interface{}{
		"Articles": articles,
		"Pager":    pager,
		"Tags":     getAllTags(),
	}
	//	updateSidebarData(data)
	ctx.Loader("theme").Render("index.html", data)
}

func ContentHandler(ctx *Golf.Context) {
	slug, _ := ctx.Param("slug")
	article, err := model.GetPostBySlug(slug)
	log.Printf("Error: %v", err)
	if err != nil {
		ctx.Abort(404)
		return
	}
	article.Hits++
	data := map[string]interface{}{
		"Title":    article.Title,
		"Article":  article,
		"Content":  article,
		"Comments": article.Comments,
	}
	//		updateSidebarData(data)
	if article.IsPage {
		ctx.Loader("theme").Render("page.html", data)
	} else {
		ctx.Loader("theme").Render("article.html", data)
	}
}

//
//func CommentHandler(ctx *Golf.Context) {
//	id, _ := ctx.Param("id")
//	cid, _ := strconv.Atoi(id)
//	if cid < 1 || model.GetContentById(cid) == nil {
//		ctx.JSON(map[string]interface{}{
//			"res": false,
//		})
//	}
//	c := new(model.Comment)
//	c.Author = ctx.Request.FormValue("user")
//	c.Email = ctx.Request.FormValue("email")
//	c.Url = ctx.Request.FormValue("url")
//	c.Content = ctx.Request.FormValue("content")
//	c.Avatar = utils.Gravatar(c.Email, "50")
//	c.Pid, _ = strconv.Atoi(ctx.Request.FormValue("pid"))
//	c.Ip = ctx.Request.RemoteAddr
//	c.UserAgent = ctx.Request.UserAgent()
//	c.IsAdmin = false
//	msg := validateComment(c)
//	if msg == "" {
//		model.CreateComment(cid, c)
//		ctx.JSON(map[string]interface{}{
//			"res":     true,
//			"comment": c.ToJson(),
//		})
//		model.CreateMessage("comment", c)
//	} else {
//		ctx.JSON(map[string]interface{}{
//			"res": false,
//			"msg": msg,
//		})
//	}
//}
//
//func validateComment(c *model.Comment) string {
//	if utils.IsEmptyString(c.Author) || utils.IsEmptyString(c.Content) {
//		return "Name, Email and Content are required fields."
//	}
//	if !utils.IsEmail(c.Email) {
//		return "Email format not valid."
//	}
//	if !utils.IsEmptyString(c.Url) && !utils.IsURL(c.Url) {
//		return "Website URL format not valid."
//	}
//	return ""
//}
//
//func TagHandler(ctx *Golf.Context) {
//	p, _ := ctx.Param("page")
//	page, _ := strconv.Atoi(p)
//	t, _ := ctx.Param("tag")
//	tag, _ := url.QueryUnescape(t)
//	size := getArticleListSize()
//	articles, pager := model.GetTaggedArticleList(tag, page, getArticleListSize())
//	// fix dotted tag
//	if len(articles) < 1 && strings.Contains(tag, "-") {
//		articles, pager = model.GetTaggedArticleList(strings.Replace(tag, "-", ".", -1), page, size)
//	}
//	data := map[string]interface{}{
//		"Articles": articles,
//		"Pager":    pager,
//		"Tag":      tag,
//		"Title":    tag,
//	}
//	updateSidebarData(data)
//	ctx.Loader("theme").Render("tag.html", data)
//}
//
//func SiteMapHandler(ctx *Golf.Context) {
//	baseUrl := model.GetSetting("site_url")
//	println(baseUrl)
//	article, _ := model.GetPublishArticleList(1, 50)
//	navigators := model.GetNavigators()
//	now := time.Unix(utils.Now(), 0).Format(time.RFC3339)
//
//	articleMap := make([]map[string]string, len(article))
//	for i, a := range article {
//		m := make(map[string]string)
//		m["Link"] = strings.Replace(baseUrl+a.Link(), baseUrl+"/", baseUrl, -1)
//		m["Created"] = time.Unix(a.CreateTime, 0).Format(time.RFC3339)
//		articleMap[i] = m
//	}
//
//	navMap := make([]map[string]string, 0)
//	for _, n := range navigators {
//		m := make(map[string]string)
//		if n.Link == "/" {
//			continue
//		}
//		if strings.HasPrefix(n.Link, "/") {
//			m["Link"] = strings.Replace(baseUrl+n.Link, baseUrl+"/", baseUrl, -1)
//		} else {
//			m["Link"] = n.Link
//		}
//		m["Created"] = now
//		navMap = append(navMap, m)
//	}
//
//	ctx.Header["Content-Type"] = "application/rss+xml;charset=UTF-8"
//	ctx.Loader("base").Render("sitemap.xml", map[string]interface{}{
//		"Title":      model.GetSetting("site_title"),
//		"Link":       baseUrl,
//		"Created":    now,
//		"Articles":   articleMap,
//		"Navigators": navMap,
//	})
//}
//
//func RssHandler(ctx *Golf.Context) {
//	baseUrl := model.GetSetting("site_url")
//	article, _ := model.GetPublishArticleList(1, 20)
//	author := model.GetUsersByRole("ADMIN")[0]
//
//	articleMap := make([]map[string]string, len(article))
//	for i, a := range article {
//		m := make(map[string]string)
//		m["Title"] = a.Title
//		m["Link"] = strings.Replace(baseUrl+a.Link(), baseUrl+"/", baseUrl, -1)
//		m["Author"] = author.Nick
//		str := utils.Markdown2Html(a.Content())
//		str = strings.Replace(str, `src="/`, `src="`+strings.TrimSuffix(baseUrl, "/")+"/", -1)
//		str = strings.Replace(str, `href="/`, `href="`+strings.TrimSuffix(baseUrl, "/")+"/", -1)
//		m["Desc"] = str
//		m["Created"] = time.Unix(a.CreateTime, 0).Format(time.RFC822)
//		articleMap[i] = m
//	}
//
//	ctx.Header["Content-Type"] = "application/rss+xml;charset=UTF-8"
//
//	ctx.Loader("base").Render("rss.xml", map[string]interface{}{
//		"Title":    model.GetSetting("site_title"),
//		"Link":     baseUrl,
//		"Desc":     model.GetSetting("site_description"),
//		"Created":  time.Unix(utils.Now(), 0).Format(time.RFC822),
//		"Articles": articleMap,
//	})
//}