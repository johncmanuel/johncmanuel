<h1 align="center">johncmanuel 🇵🇭 🇺🇸</h1>

- Software engineer with experience in full-stack web development 
- Currently completing my BS in Computer Science at Cal State Fullerton
- {{.PublicReposCount}} public repositories, {{.StarGazersCount}} total stars
- Most used languages:  {{- range $i, $lang := .Languages -}}{{.Language | html}} ({{.Percentage}}%){{if notLastElement $i $.Languages}}, {{end}}{{end}} 
- Interests: Full-Stack Development, Game Development, Distributed Systems
- Hobbies: Video Games, Anime, Code, Writing, Manga, Reading

## Important Links

- Website: [johncarlomanuel.com](https://johncarlomanuel.com/)
- Resume: [johncmanuel.com/resume](https://johncarlomanuel.com/resume)
- Blog: [registers.johncarlomanuel.com](https://registers.johncarlomanuel.com/)

For quick responses, [email me](mailto:johncnmanuel@gmail.com) or [send me a DM on X](https://x.com/messages/compose?recipient_id=1727183654676500480).

<details>
<summary><b>bio in the programmer way</b></summary>
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
</details>

<hr />

<a href="https://johncarlomanuel.com/" target="_blank"><img src="media/banner.png" alt="banner" /></a>
