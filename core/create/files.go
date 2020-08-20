// Generated by cmd/generate/main.go. Do not edit
// manually, run go generate in the root instead.

package create

var (
	files = map[string]string{

		"assets/css/style.css": `* {
    font-family: Arial, Tahoma, sans-serif;
}

body {
    padding: 4rem;
}

main {
    margin-right: auto;
    margin-left: auto;
}`,

		"content/about.md": `---
Title: About me
---

Hi! My name is Clara Crema and I love coffee. In this blog, I'll guide
you through coffee making techniques using a portafilter Espresso machine
for making Espresso-based coffees.

Feel free to contact me if you have an idea for a new blog post. Have fun!`,

		"content/blog/making-barista-quality-espresso.md": `---
Title: Making Barista-Quality Espresso
Date: 2020-08-14
Description: This is a guide for making italian Espresso.
Img: espresso.jpg
Credit: Burst on Pexels
---

Do you enjoy a high-quality italian Espresso as much as I do? Quite frankly,
making Espresso at this level isn't easy - but with the right tools, patience
and practice, you'll be able to make delicious coffee at home.

...`,

		"content/blog/steaming-milk-for-cappuccino.md": `---
Title: Steaming Milk for Cappuccino
Date: 2020-08-15
Description: This is a guide for steaming milk.
Img: /assets/img/blog/cappuccino.jpg
Credit: Chevanon Photography on Pexels
---

Did you ever steam milk using a portafilter machine? Aside from the more
technical aspects like milk temperature and jug size, steaming milk is a
matter of feeling and experience.

...`,

		"templates/index-page.html": `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>{{.Meta.Title}}</title>
        <meta name="author" content="{{.Meta.Author}}" />
        <meta name="description" content="{{.Meta.Description}}" />
        <link rel="stylesheet" type="text/css" href="/assets/css/style.css" />
    </head>
    <body>
        <main>
            {{range $page := .Pages}}
                <div>
                    <h3>{{$page.Title}}</h3>
                    <p><small>Posted on {{$page.Date.Format "Jan 2 2006"}}</small></p>
                    <p>{{$page.Description}}</p>
                    <p><a href="{{$page.ID}}">read post</a></p>
                </div>
            {{end}}
        </main>
    </body>
</html>`,

		"templates/page.html": `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>{{.Page.Title}}</title>
        <meta name="author" content="{{.Meta.Author}}" />
        <meta name="description" content="{{.Page.Description}}" />
        <link rel="stylesheet" type="text/css" href="/assets/css/style.css" />
    </head>
    <body>
        <main>
            <h1>{{.Page.Title}}</h1>
            <h4>{{.Page.Description}}</h4>

            {{if .Page.Img}}
                <img src="{{.Page.Img}}" alt="{{.Page.Title}}" />
                <p><small>Image: {{.Page.Credit}}</small></p>
            {{end}}

            {{if .Page.Date}}
                <p>Posted on {{.Page.Date.Format "Jan 2 2006"}}</p>
            {{end}}

            {{.Page.Content}}
        </main>
    </body>
</html>`,

		"verless.yml": `# verless.yml is your project configuration. This example
# file contains all configuration keys available.
version: 1
site:
  # General information about your project.
  meta:
    title: My Coffee Blog
    subtitle: About Espresso & Cappuccino
    description: I'm Clara and write alot about coffee. Welcome to my blog!
    author: Clara Crema
    base: http://localhost
  # Settings for your website's navigation.
  nav:
    items:
      - label: Instagram
        target: https://instagram.com
    overwrite: false
  # Settings for your website's footer.
  footer:
    items:
      - label: Home
        target: http://localhost
    overwrite: false
# Enable plugins for your project.
plugins:
  - atom
build:
  # Uncomment the following line to automatically overwrite
  # the output directory without having to use --overwrite.
  # Use with caution!
  # overwrite: true
`,
	}
)