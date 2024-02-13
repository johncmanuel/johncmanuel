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

For quick responses, email me.

<a href="mailto:mail@johncarlomanuel.com">
    <img alt="Email" src="https://img.shields.io/badge/Email Me-D14836?style=for-the-badge&logoColor=white" />
</a>

## Socials

<p>
  <a href="https://www.linkedin.com/in/johncarlomanuel/" target="_blank">
    <img alt="LinkedIn" src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" />
  </a>
 <a href="https://www.johncarlomanuel.com/" target="_blank">
    <img alt="Website" src="https://img.shields.io/badge/website-000000?style=for-the-badge&logo=About.me&logoColor=white" />
  </a>
  <a rel="me" href="https://mastodon.social/@johncmanuel" target="_blank">
  <img alt="Mastodon" src="https://img.shields.io/badge/-MASTODON-%232B90D9?style=for-the-badge&logo=mastodon&logoColor=white">
  </a>
  <a href="https://www.youtube.com/@johncarlomanuel" target="_blank">
  <img alt="YouTube" src="https://img.shields.io/badge/YouTube-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white"></img></a>
  <a href="https://twitter.com/johncmanuel" target="_blank">
  <img alt="Twitter" src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white"></img>
  </a>
</p>

<hr />

<a href="https://www.johncarlomanuel.com/" target="_blank"><img src="media/banner.png" alt="banner" /></a>
