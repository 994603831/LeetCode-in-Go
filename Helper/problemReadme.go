package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func creatREADME(p problem) {
	fileFormat := `# [%d. %s](%s)

%s
`

	questionDescription := strings.TrimSpace(getDescription(p.link()))

	content := fmt.Sprintf(fileFormat, p.ID, p.Title, p.link(), questionDescription)

	content = replaceCharacters(content)

	filename := fmt.Sprintf("%s/README.md", p.Dir())

	write(filename, content)

	vscodeOpen(filename)

}

func replaceCharacters(s string) string {
	changeMap := map[string]string{
		"&amp;lt;":   "<",
		"&amp;quot;": "\"",
		"&amp;nbsp;": " ",
		"&amp;#39;":  "`",
		"&quot;":     "\"",
		"&lt;":       "<",
		"&gt;":       ">",
		"&ge;":       ">=",
		"&nbsp;":     "`",
		"&amp;":      "&",
		"&#39;":      "'",
		"   \n":      "\n",
		"  \n":       "\n",
		" \n":        "\n",
		"\n\n\n\n\n": "\n\n",
		"\n\n\n\n":   "\n\n",
		"\n\n\n":     "\n\n",
	}

	olds := make([]string, 0, len(changeMap))
	for old := range changeMap {
		olds = append(olds, old)
	}

	sort.Strings(olds)

	news := make([]string, 0, len(olds))
	for _, old := range olds {
		news = append(news, changeMap[old])
	}

	for i := len(olds) - 1; 0 <= i; i-- {
		// 先替换长的，再替换短的
		old, new := olds[i], news[i]
		s = strings.Replace(s, old, new, -1)
	}

	return s

}

func getDescription(url string) string {
	log.Printf("准备访问 %s", url)

	raw := string(getRaw(url))

	sub := "<meta name=\"description\" content=\""
	index := strings.Index(raw, sub)
	raw = raw[index+len(sub):]

	sub = "\""
	index = strings.Index(raw, sub)
	raw = raw[:index]

	fmt.Println(raw)

	return raw
}

// 	var err error

// 	// create context
// 	ctxt, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	// create chrome instance
// 	// c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
// 	c, err := chromedp.New(ctxt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// run task list
// 	var res string
// 	err = c.Run(ctxt, text(url, &res))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// shutdown chrome
// 	err = c.Shutdown(ctxt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// wait for chrome to finish
// 	err = c.Wait()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Desc:", res)

// 	return res
// }

// func text(url string, res *string) chromedp.Tasks {
// 	sel := `#question-detail-main-tabs > div.tab-pane__280T.css-c363ty-TabContent.e5i1odf5 > div > div.content__eAC7`
// 	return chromedp.Tasks{
// 		chromedp.Navigate(url),
// 		chromedp.Text(sel, res, chromedp.NodeVisible, chromedp.BySearch),
// 	}
// }
