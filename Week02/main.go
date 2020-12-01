

type Article struct {
	id      int
	tagID   int
	title   string
	content string
}

// dao

func GetArticleByTitle(title string) (*model.Article, error) {
	article := &model.Article{}
	err := sql.ErrNoRows
	if err != nil {
		return nil, err
	}
	return article, nil
}

//serice
func GetArticleByTitle(title string) (*model.Article, error) {
	article, err := dao.GetArticleByTitle(title)
	// 检查是否是 “查询结果为空” 类型的错误，如果是，则吞掉该错误
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	// 其他错误
	// return nil, err
	return article, nil
}

func main() {

	http.HandleFunc("/", GetArticle)
	http.ListenAndServe(":8080", nil)
}

func GetArticle(w http.ResponseWriter, req *http.Request) {
	title := "golang"
	// 略去参数检验
	article, err := service.GetArticleByTitle(title)
	if err != nil {
		log.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if article == nil {
		msg := fmt.Sprintf("There is not article which title is %s", title)
		w.Write([]byte(msg))
		return
	}

	bytes, err := json.Marshal(article)
	if err != nil {
		log.Printf("handler Marshal err: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(bytes)
}