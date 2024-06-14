<h1 align="center">johncmanuel 🇵🇭 🇺🇸</h1>

```python
>>> from goated_programmers import johncmanuel
>>> import json
>>> john = johncmanuel()
>>> print(json.dumps(john.bio, indent=2))
{
  "name": "John Carlo Manuel",
  "occupation": "Software Engineer",
  "pronouns": "he/him",
  "schools": [
    "Skyline College",
    "California State University, Fullerton"
  ],
  "interests": [
    "Full-Stack Development",
    "Game Development",
    "Distributed Systems"
  ],
  "github_stats": {
    "public_repos": {{.PublicReposCount}},
    "stargazers": {{.StarGazersCount}}
  },
  "languages": {
    {{- range $i, $lang := .Languages}}
    "{{.Language | html}}": {
      "usage_percent": {{.Percentage}}
    }{{if notLastElement $i $.Languages}},{{end}}{{end}}
  },
  "hobbies": [
    "Video Games",
    "Anime",
    "Code",
    "Writing",
    "Manga",
    "Reading"
  ]
}
```

## Contact

For quick responses, [email me](mailto:mail@johncarlomanuel.com).

<hr />

<a href="https://www.johncarlomanuel.com/" target="_blank"><img src="media/banner.png" alt="banner" /></a>
