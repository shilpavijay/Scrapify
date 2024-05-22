# Scrapify
A Web Scraper built using Golang

Sample website used for scraping: https://www.chrismytton.com/plain-text-websites/

The function takes two arguments:
* The Target Element: "div"
* The target Class: "post"

For the above example, the function parses through the page to find the class "post" that contains the "div" tag and returns the contents inside the tag:

"
In a world of web apps, videos, social media and multi-media distraction it’s nice to know there are still some websites out there which use simple plain text.

I came across Sijmen J. Mulder’s list of Hyperlinked Text websites the other day, which is a list of websites that mostly use plain old HTML, and skip the JavaScript bloat that has become so common to the web these days.

The sites on this page don’t have loads of adverts, or many images, they’re just simple plain text, and honestly it’s a pleasure reading them. They load fast and you can get straight to the content you’re after without scrolling past adverts and dismissing popups inviting you to provide your email address for dubious purposes.
"