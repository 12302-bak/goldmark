package goldmark_test

import (
	"bytes"
	"fmt"
	mathjax "github.com/litao91/goldmark-mathjax"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
	"go.abhg.dev/goldmark/mermaid"
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	md := goldmark.New(
		goldmark.WithExtensions(
			//extension.GFM,
			// emoji.Emoji,
			extension.NewTable(
				extension.WithTableCellAlignMethod(extension.TableCellAlignAttribute),
			),
			//extension.NewCommentBlockExtension(),
			mathjax.MathJax,
			&mermaid.Extender{
				RenderMode:   mermaid.RenderModeClient,
				ContainerTag: "div",
				NoScript:     true,
			},
		),
		goldmark.WithParserOptions(
			//parser.WithAutoHeadingID(),
			parser.WithAttribute(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			//html.WithXHTML(),
			html.WithUnsafe(),
		),
	)

	source, err := os.ReadFile("/Users/stevenobelia/Desktop/archive/archive-docsify/README.md")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}
	h := buf.String()
	fmt.Printf("\n%s\n", h)
	writeErr := os.WriteFile("/Users/stevenobelia/Desktop/archive/archive-docsify/README.html", buf.Bytes(), 0644)
	if writeErr != nil {
		panic(writeErr)
	}
	fmt.Println()
}

func TestMathjax(t *testing.T) {
	md := goldmark.New(
		goldmark.WithExtensions(mathjax.MathJax),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	// todo more control on the parsing process
	var html bytes.Buffer
	mdContent := []byte(`
$$
\mathbb{E}(X) = \int x d F(x) = \left\{ \begin{aligned} \sum_x x f(x) \; & \text{ if } X \text{ is discrete} 
\\ \int x f(x) dx \; & \text{ if } X \text{ is continuous }
\end{aligned} \right.
$$


Inline math $\frac{1}{2}$
`)
	if err := md.Convert(mdContent, &html); err != nil {
		fmt.Println(err)
	}
	fmt.Println(html.String())
}

func TestString(t *testing.T) {
	data := "NI"
	length := util.TrimLeftSpaceLength([]byte(data[2:]))
	if length > 0 {
	}
}
