Small project for Boot.dev 

This project is still a WIP as its something I use to help me learn Bosnian!

What does it do?
- App makes simple requests to MyMemory API to get translations based on the applications supported languanges (currently only Bosnia, Portuguese and English)
- Caches the requests to avoid spamming the API, basically creating a local dictionary with time!

How to use it:
- After cloning the repo:
  - CD into the folder and run "make install" and then run the app from anywhere on your terminal with the following commands:
    - bos <phrase> <language pair> - If only one word, quotations ("") are not necessary | Language pair is used for the API to know what to translate, current available ones are:
      - ba (Bosnian to English), en (English to Bosnian), pt (Portuguese to Bosnian)
    - bos --help - Displays how to use all the commands

To-do:
- Create a new table for "Words", filter each word from the phrase and save their specific translations
- Figure out a way to get more in-depth information on specific words (like verb variations)
- Add a --history command to browse most recent searches (also add a searched_at entry to the "phrases" table)
